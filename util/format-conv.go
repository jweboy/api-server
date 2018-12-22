package util

import (
	"fmt"
	"net/url"
	"strconv"
)

func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("String compile error", err)
	}
	return i
}

func DecodeStr(str string) string {
	s, err := url.QueryUnescape(str)
	if err != nil {
		fmt.Println("String decode error", err)
	}
	return s
}
