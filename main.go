package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

type Input struct {
	Temp   float64 `json:"temp"`
	Humid  float64 `json:"humid"`
	Gambar string  `json:"gambar"`
}

type Form struct {
	Temp   string `form:"temp" binding:"required"`
	Humid  string `form:"humid" binding:"required"`
	Gambar string `form:"gambar" binding:"required"`
}

func main() {
	topic := [3]string{"fd/temp", "fd/humid", "fd/send"}

	broker := "mqtt://103.127.98.21:1883"
	clientID := "testing2"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	r := gin.Default()

	r.POST("/img", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)

		request := string(body)

		// fmt.Println(request)

		token := client.Publish("fd/send", 0, false, request)
		token.Wait()
	})

	r.POST("/humid", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)

		request := string(body)

		fmt.Println(request)

		token := client.Publish("fd/humid", 0, false, request)
		token.Wait()
	})

	r.POST("/temp", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)

		request := string(body)

		fmt.Println(request)

		token := client.Publish("fd/temp", 0, false, request)
		token.Wait()
	})

	r.POST("all", func(c *gin.Context) {
		// var input Input
		// c.ShouldBindJSON(&input)

		// fmt.Println(input.Temp)
		// fmt.Println(input.Humid)
		// fmt.Println(input.Gambar)

		// token := client.Publish("fd/temp", 0, false, fmt.Sprintf("%f", input.Temp))
		// token.Wait()
		// token = client.Publish("fd/humid", 0, false, fmt.Sprintf("%f", input.Humid))
		// token.Wait()
		body, _ := ioutil.ReadAll(c.Request.Body)

		request := string(body)
		array := strings.Split(request, "&data=")
		// array := strings.Split(request, "mepmepmep,")c

		// fmt.Println(array)
		// for i, element := range array {
		// 	token := client.Publish(topic[i], 0, false, element)
		// 	token.Wait()
		// }

		fmt.Println(array[0])
		token := client.Publish(topic[0], 0, false, array[1])
		token.Wait()
		token = client.Publish(topic[1], 0, false, array[2])
		token.Wait()
		token = client.Publish(topic[2], 0, false, array[3])
		token.Wait()
	})

	r.Run()
}
