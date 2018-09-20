package main

import (
	"errors"
	"net/http"
	"time"

	"restful-api-server/config"
	"restful-api-server/model"
	"restful-api-server/router"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

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
