package main

import (
	"context"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func postLeads(c *gin.Context) {

	lead := map[string]string{}

	c.Request.ParseForm()

	for i, v := range c.Request.PostForm {

		s := strings.Join(v, " ")
		lead[i] = s

	}

	go guardarRedis(lead)

}

func guardarRedis(mapa map[string]string) {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	r := strconv.Itoa(rand.Intn(999) + 1000)

	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {

		for i, v := range mapa {

			rdb.HSet(ctx, "lead-"+r, i, v)

		}

		return nil
	}); err != nil {
		panic(err)
	}

}

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/LandingPage", postLeads)
	router.Run(":8080")

}
