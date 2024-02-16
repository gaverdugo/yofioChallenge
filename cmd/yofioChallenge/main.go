package main

import (
	"github.com/gaverdugo/yofioChallenge/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)
	router.Run()
}
