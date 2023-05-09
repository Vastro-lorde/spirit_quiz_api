package data

import (
	"github.com/google/uuid"
)

type CreateCategory struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	ImageUrl string    `json:"image_url"`
}

var Categories = []CreateCategory{
	{ID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Name: "JavaScript", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594865/spirit_quiz/ku42ccoh7r3rzsznm0bj.png"},
	{ID: uuid.MustParse("d3cc00e8-7325-44e7-a51d-87d3f48ee3c3"), Name: "Python", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594763/spirit_quiz/osv9afullssjco3bnztg.png"},
	{ID: uuid.MustParse("2e1ba1b7-3b50-4f2d-8c26-cc62ca89a6b7"), Name: "Java", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594630/spirit_quiz/kkpnqt6iywiyl8joppfy.png"},
	{ID: uuid.MustParse("26c38434-6eeb-4a7b-b0f7-f3c2d0e3b62c"), Name: "Ruby", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594661/spirit_quiz/bpfhiq0imu9uwewz6jx7.png"},
	{ID: uuid.MustParse("4b0fac10-38f0-47a5-98c2-a7f1d2cf98b4"), Name: "C++", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594832/spirit_quiz/eigdr3lf3fm7asoe9s1l.png"},
	{ID: uuid.MustParse("c5b5f32c-0c96-45d5-9f34-6d198ae42e70"), Name: "Go", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594505/spirit_quiz/ycmqgkbdmcqj9vn7dli9.png"},
	{ID: uuid.MustParse("090c28d4-91e4-4b6d-b8c4-e6f11a4a2f23"), Name: "Swift", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594693/spirit_quiz/svo7be8kenni8vsqjouh.png"},
	{ID: uuid.MustParse("1a2e99ee-47eb-466a-ba7b-13e4243f3f14"), Name: "PHP", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594566/spirit_quiz/pkwxr51yideacybfiz0u.png"},
	{ID: uuid.MustParse("4c8d93b7-4aa3-476f-b12d-cc19aee0cbb9"), Name: "TypeScript", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594727/spirit_quiz/ermrbc9wa7wbdqfuhtiy.png"},
	{ID: uuid.MustParse("b6e0f13a-f88b-4f23-ae33-8f7022df738c"), Name: "Kotlin", ImageUrl: "http://res.cloudinary.com/spiritvd/image/upload/v1683594599/spirit_quiz/exglokhesx4gpgwp2qfu.png"},
}
