package seed

import (
	"spirit_quiz/internal/data/models"
	"spirit_quiz/internal/data/seed/data"
	"sync"

	"gorm.io/gorm"
)
var seederLock sync.Mutex

func Seeder(db *gorm.DB) error {
	seederLock.Lock()
	defer seederLock.Unlock()
	var mu sync.Mutex


	var categoriesCheck []models.Category
	if err := db.Find(&categoriesCheck).Error; err != nil {
		return err
	}

	// if the categories table is not empty, it means that the data has already been seeded
	if len(categoriesCheck) > 0 {
		return nil
	}

	categories := data.Categories
	for i := range categories {
		mu.Lock()
		category := &categories[i]

		var newCategory models.Category
		newCategory.ID = category.ID
		newCategory.Name = category.Name
		newCategory.ImageUrl = category.ImageUrl

		err := db.Create(&newCategory).Error
		mu.Unlock()
		if err != nil {
			return err
		}
	}

	var questionsCheck []models.Question
	if err := db.Find(&questionsCheck).Error; err != nil {
		return err
	}

	// if the categories table is not empty, it means that the data has already been seeded
	if len(questionsCheck) > 0 {
		return nil
	}

	questions := data.Questions

	for i := range questions {
		mu.Lock()
		question := &questions[i]

		var newQuestion models.Question
		newQuestion.ID = question.ID
		newQuestion.CategoryID = question.CategoryID
		newQuestion.Text = question.Text

		err := db.Create(&newQuestion).Error
		mu.Unlock()
		if err != nil {
			return err
		}
	}

	var optionsCheck []models.Option
	if err := db.Find(&optionsCheck).Error; err != nil {
		return err
	}

	// if the categories table is not empty, it means that the data has already been seeded
	if len(optionsCheck) > 0 {
		return nil
	}

	options := data.Options

	for i := range options {
		mu.Lock()
		option := &options[i]

		var newOption models.Option
		newOption.ID = option.ID
		newOption.QuestionID = option.QuestionID
		newOption.Text = option.Text
		newOption.IsCorrect = option.IsCorrect

		err := db.Create(&newOption).Error
		mu.Unlock()
		if err != nil {
			return err
		}
	}

	return nil
}
