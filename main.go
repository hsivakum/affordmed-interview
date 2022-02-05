package main

import (
	. "affordmed/init"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".envrc"); err != nil {
		logrus.Info("error loading env config")
	}

	router := gin.Default()

	Init(router)

	err := router.Run(":8080")
	if err != nil {
		logrus.Fatalf("unable to span the server %v", err)
	}
	//go oddAndEven()
}

var data = make(chan int)

func oddAndEven() {
	go evenValue()
	go oddValue()

	for _, value := range []int{1, 2, 3, 4, 5, 6} {
		go func(i int) {
			data <- i
		}(value)
	}

}

func consume() {
	for {
		value := <-data
		if value%2 != 0 {

			fmt.Printf("this is odd %v", value)
		} else {
			fmt.Printf("this is even %v", value)
		}
	}
}

func oddValue() {
	value := <-data
	if value%2 != 0 {
		fmt.Printf("this is odd %v", value)
	}
}

func evenValue() {
	value := <-data
	if value%2 == 0 {
		fmt.Printf("this is even %v", value)
	}
}
