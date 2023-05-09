package handlers

import (
	"fmt"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context) {

	// Retrieve file from request body
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}

	// Open file
	fileReader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileReader.Close()

	params := uploader.UploadParams{
		Folder:       "spirit_quiz",
		ResourceType: "image",
		Type:         "upload",
	}
	// Upload file to Cloudinary
	uploadResult, err := cloudinaryClient.Upload.Upload(c.Request.Context(), fileReader, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to Cloudinary"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "url": uploadResult.URL})
}

func ImageDelete(context *gin.Context) {

	type ImageUrl struct {
		PublicUrl string `json:"public_url" binding:"required"`
	}
	var imageUrl ImageUrl
	if err := context.ShouldBindJSON(&imageUrl); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	params := uploader.DestroyParams{
		ResourceType: "image",
		PublicID:     imageUrl.PublicUrl,
	}

	fmt.Println(imageUrl.PublicUrl)
	// Delete image from Cloudinary
	deleteResult, err := cloudinaryClient.Upload.Destroy(context.Request.Context(), params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if image was successfully deleted
	if deleteResult.Result != "ok" {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":    deleteResult.Error.Message,
			"response": deleteResult.Result,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
