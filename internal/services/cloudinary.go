package services

import (
	"log"
	"spirit_quiz/config"
	"github.com/cloudinary/cloudinary-go/v2"
)

var configs, err = config.GetConfigs()
func CloudinaryService() (*cloudinary.Cloudinary,error){
	
	if err != nil {
		log.Fatal("Error getting env variable: ", err)
		return nil,err
	}
	cloudinaryConfig, err := cloudinary.NewFromURL(configs.CloudinaryUrl)
	if err != nil {
		log.Fatal("Cloudinary configuration error: ", err)
		return nil,err
	}
	cloudinaryConfig.Config.URL.Secure = true

	return cloudinaryConfig,nil
}
