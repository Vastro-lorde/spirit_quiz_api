package data

import "github.com/google/uuid"

type CreateQuestion struct {
	ID         uuid.UUID `json:"id" binding:"required"`
	CategoryID uuid.UUID `json:"category_id" binding:"required"`
	Text       string    `json:"text" binding:"required"`
}

var Questions = []CreateQuestion{
	{ID: uuid.MustParse("12e7b982-1c19-4699-94d2-8c21ed52327e"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What is JavaScript?"},
	{ID: uuid.MustParse("78f7c8a6-f7e5-40d4-8ebf-ebf7c68016d9"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the data types in JavaScript?"},
	{ID: uuid.MustParse("eedca8a3-2361-4284-a1e7-634234f1f0d5"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the differences between null and undefined in JavaScript?"},
	{ID: uuid.MustParse("82d5633b-b69e-4c8c-ba37-2929f0ad7473"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the different ways to declare variables in JavaScript?"},
	{ID: uuid.MustParse("0ed8704b-d3d6-45a3-92d3-118ce3f1381c"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What is a closure in JavaScript?"},
	{ID: uuid.MustParse("e6c7e2c6-c604-4dd5-a0dc-1010c930a6e7"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What is the difference between == and === operators in JavaScript?"},
	{ID: uuid.MustParse("96f47d7f-3443-4f68-af16-05b754d07e88"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What is the difference between let and var in JavaScript?"},
	{ID: uuid.MustParse("6f82f33b-35b3-4d61-bfe6-73addd7d05a9"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the different types of events in JavaScript?"},
	{ID: uuid.MustParse("6c203f4b-5402-47e5-8c4d-7f4ab41b35b5"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the differences between synchronous and asynchronous programming in JavaScript?"},
	{ID: uuid.MustParse("2713c1fc-fa4b-437a-83d3-72a8a3f068d3"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the different ways to create an object in JavaScript?"},
	{ID: uuid.MustParse("8a1c3f4a-6c27-47a7-80e6-c3b7484b0a9f"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What is hoisting in JavaScript?"},
	{ID: uuid.MustParse("47f7aa34-d39b-4d2d-a7a7-7711f0b163fe"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the different types of error in JavaScript?"},
	{ID: uuid.MustParse("d4d0fdca-6ee0-4b4f-b1aa-69b8c5eb9eb3"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What is the purpose of the 'use strict' statement in JavaScript?"},
	{ID: uuid.MustParse("a6d6b45d-d1b6-468e-8e5b-f2c693104a7a"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What is the purpose of the 'this' keyword in JavaScript?"},
	{ID: uuid.MustParse("94a80e6b-b58c-49a9-b994-d32c13b45a71"), CategoryID: uuid.MustParse("0fd60066-3ace-4856-8541-27894de8af2d"), Text: "What are the different types of loops in JavaScript?"},
}
