package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"restful-api-server/config"
	"restful-api-server/model"
	"restful-api-server/router"

	v "restful-api-server/pkg/version"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

// main 入口函数
// @title Restful API
// @description This is restful api server with golang.
// @termsOfService https://github.com/jweboy/restfult-api-server

// @contract.name jweboy
// @contract.url https://jweboy.github.io/
// @contract.email jweboy0630@gmail.com

// @license.name MIT

// @host localhost:4000
// @BasePath /v1
func main() {
	pflag.Parse()
	// Get version
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")

		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// Init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// Init db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes
	router.Load(
		// Cores
		g,

		// Middlewares
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	tlsAddr := viper.GetString("tls.addr")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", tlsAddr)
			log.Info(http.ListenAndServeTLS(tlsAddr, cert, key, g).Error())
		}()
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router")
}
