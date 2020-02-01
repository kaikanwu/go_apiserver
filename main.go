package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"main/router"
	"net/http"
	"time"
)

func main() {

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

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

}

// ping server to make sure the router is working
func pingServer() error {

	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// sleep one second, retry
		log.Print("waiting for the router, retry in one second")
		time.Sleep(time.Second)

	}
	return errors.New("CANNOT CONNECT TO THE ROUTER")

}
