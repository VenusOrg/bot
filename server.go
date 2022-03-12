package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunServer(batchSendMap map[string]gin.HandlerFunc) {
	router := gin.Default()

	// All kinds of transfer call
	for key, handler := range batchSendMap {
		router.GET(fmt.Sprintf("/%s", key), handler)
	}

	router.Run(":8080")
}
