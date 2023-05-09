package seed

import (
	"spirit_quiz/internal/data/models"
	"spirit_quiz/internal/data/seed/data"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {

	categories := data.Categories
	for i := range categories {
		category := &categories[i]
		var newCategory models.Category
		newCategory.ID = category.ID
		newCategory.Name = category.Name
		newCategory.ImageUrl = category.ImageUrl
		
		err := db.Create(&newCategory).Error
		if err != nil {
			return err
		}
	}

	questions := data.Questions

	for i := range questions {
		question := &questions[i]

		var newQuestion models.Question
		newQuestion.ID = question.ID
		newQuestion.CategoryID = question.CategoryID
		newQuestion.Text = question.Text

		err := db.Create(&newQuestion).Error
		if err != nil {
			return err
		}
	}

	options := data.Options

	for i := range questions {
		option := &options[i]

		var newOption models.Option
		newOption.ID = option.ID
		newOption.QuestionID = option.QuestionID
		newOption.Text = option.Text
		newOption.IsCorrect = option.IsCorrect

		err := db.Create(&newOption).Error
		if err != nil {
			return err
		}
	}

	return nil
}
