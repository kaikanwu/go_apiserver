package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"main/config"
	"main/router"
	"net/http"
	"time"
)

var(
	cfg = pflag.StringP("config","c","","go api server config file path")
)

func main() {
	// 启动 config
	pflag.Parse()

	if err :=config.Init(*cfg); err != nil {
		panic(err)
	}

	// 启动 gin
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	//ping the server
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router is no response, or it may took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// ping server to make sure the router is working
func pingServer() error {

	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// sleep one second, retry
		log.Print("waiting for the router, retry in one second")
		time.Sleep(time.Second)

	}
	return errors.New("CANNOT CONNECT TO THE ROUTER")

}
