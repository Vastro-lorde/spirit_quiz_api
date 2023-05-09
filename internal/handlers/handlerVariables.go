package handlers

import (
	"spirit_quiz/internal/data/database"
	"spirit_quiz/internal/services"
)

var db, err = database.Connect()

var cloudinaryClient, _ = services.CloudinaryService()