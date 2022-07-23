package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/ocr/idcard-driver", usebOcr)

	r.Run(":8082")
}
