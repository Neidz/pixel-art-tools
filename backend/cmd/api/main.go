package main

import (
	"net/http"
	"pixel-art-tools/backend/pkg/imageinfo"
	"pixel-art-tools/backend/pkg/imageloader"

	"github.com/gin-gonic/gin"
)

type MetadataDTO struct {
	ImgPath string `json:"imgPath"`
}

func main() {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Ok")
	})

	router.POST("/imageinfo/metadata", func(c *gin.Context) {
		var jsonData MetadataDTO
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		image, err := imageloader.LoadImage(jsonData.ImgPath)
		if err != nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		metadata := imageinfo.GenerateMetadata(image)
		c.JSON(http.StatusOK, metadata)
	})

	router.Run("localhost:8080")
}
