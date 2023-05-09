package data

import "github.com/google/uuid"

type CreateOption struct {
	ID         uuid.UUID `json:"id" binding:"required"`
	QuestionID uuid.UUID `json:"question_id" binding:"required"`
	Text       string    `json:"text" binding:"required"`
	IsCorrect  bool      `json:"is_correct" binding:"required"`
}

var Options = []CreateOption{
	{
		ID:         uuid.MustParse("9d8d8461-33aa-4361-8e18-6c77c9cf4da9"),
		QuestionID: uuid.MustParse("94a80e6b-b58c-49a9-b994-d32c13b45a71"),
		Text:       "The only type of loop in JavaScript is the 'for' loop.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("a9f6c7ec-2c9a-4d51-8cfc-742d350e99c8"),
		QuestionID: uuid.MustParse("94a80e6b-b58c-49a9-b994-d32c13b45a71"),
		Text:       "There are two types of loops in JavaScript: 'for' loops and 'while' loops.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("62f57c98-3247-4b2a-846e-f6f09a54c202"),
		QuestionID: uuid.MustParse("94a80e6b-b58c-49a9-b994-d32c13b45a71"),
		Text:       "There are three types of loops in JavaScript: 'for' loops, 'while' loops, and 'do-while' loops.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("ecc93236-8d99-4d1f-950a-d9d35e4a56bf"),
		QuestionID: uuid.MustParse("a6d6b45d-d1b6-468e-8e5b-f2c693104a7a"),
		Text:       "'this' is a reference to the current object and can be used to access properties and methods of that object.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("bc59df9c-3836-4dce-9d3d-c9d25ba6f95e"),
		QuestionID: uuid.MustParse("a6d6b45d-d1b6-468e-8e5b-f2c693104a7a"),
		Text:       "'this' is a reference to the parent object and can be used to access properties and methods of that object.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("f788d47e-d0b3-4a67-9cf9-cc8be72d9d03"),
		QuestionID: uuid.MustParse("a6d6b45d-d1b6-468e-8e5b-f2c693104a7a"),
		Text:       "'this' is a reference to the global object and can be used to access global variables and functions.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("32a96c3a-4f67-41b8-9d52-b3e9f58e6167"),
		QuestionID: uuid.MustParse("d4d0fdca-6ee0-4b4f-b1aa-69b8c5eb9eb3"),
		Text:       "The 'use strict' statement enforces stricter parsing and error handling rules in JavaScript.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("dbb50b1f-84eb-43f3-99c3-75c16f3a8a29"),
		QuestionID: uuid.MustParse("d4d0fdca-6ee0-4b4f-b1aa-69b8c5eb9eb3"),
		Text:       "The 'use strict' statement enables the use of modern JavaScript features.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("c5b291e9-05b8-4665-a4e5-9ccf86b2a291"),
		QuestionID: uuid.MustParse("d4d0fdca-6ee0-4b4f-b1aa-69b8c5eb9eb3"),
		Text:       "The 'use strict' statement is used to indicate that a JavaScript file should be executed in strict mode.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("2f5ef4d6-6cfc-41cf-8d1d-9d0669c6c935"),
		QuestionID: uuid.MustParse("47f7aa34-d39b-4d2d-a7a7-7711f0b163fe"),
		Text:       "There are three types of error in JavaScript: syntax errors, runtime errors, and logical errors.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("4b9da9fc-7b9b-4d10-b69c-07c5df55587e"),
		QuestionID: uuid.MustParse("47f7aa34-d39b-4d2d-a7a7-7711f0b163fe"),
		Text:       "There are two types of error in JavaScript: syntax errors and runtime errors.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("b3f0b9e4-31c9-4965-9f9f-3df5a3b7c31b"),
		QuestionID: uuid.MustParse("47f7aa34-d39b-4d2d-a7a7-7711f0b163fe"),
		Text:       "There are four types of error in JavaScript: syntax errors, runtime errors, logical errors, and type errors.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("dbcd51aa-59af-4c94-a4f5-21a34a472cc8"),
		QuestionID: uuid.MustParse("8a1c3f4a-6c27-47a7-80e6-c3b7484b0a9f"),
		Text:       "Hoisting in JavaScript is a behavior where variable and function declarations are moved to the top of their respective scopes during the compilation phase, regardless of where they were declared in the code. This means that you can use a variable or function before it has been declared, without getting a ReferenceError.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("ab8d2df0-cf31-4149-a73e-ff7eecc53950"),
		QuestionID: uuid.MustParse("8a1c3f4a-6c27-47a7-80e6-c3b7484b0a9f"),
		Text:       "Hoisting in JavaScript refers to the mechanism in which variable declarations are moved to the top of their respective scopes, but not their assignments. Function declarations are also hoisted in their entirety. This can lead to unexpected behavior when using variables before they have been declared or assigned.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("181792ee-b119-4388-9e9d-7a2bbd301b13"),
		QuestionID: uuid.MustParse("8a1c3f4a-6c27-47a7-80e6-c3b7484b0a9f"),
		Text:       "Hoisting in JavaScript is a term used to describe the behavior of moving variable and function declarations to the top of their respective scopes during the execution phase. This means that you can use a function before it has been declared, but not a variable.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("4d60cc21-2f75-4466-a6d6-89f97b033a4c"),
		QuestionID: uuid.MustParse("2713c1fc-fa4b-437a-83d3-72a8a3f068d3"),
		Text:       "There are two ways to create an object in JavaScript: using the object literal notation and using the Object constructor function.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("44763f2f-4d44-4184-8971-1f6b35e1911f"),
		QuestionID: uuid.MustParse("2713c1fc-fa4b-437a-83d3-72a8a3f068d3"),
		Text:       "There is only one way to create an object in JavaScript: using the Object constructor function.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("f647b70a-9d84-4f7a-9035-65d16d14d317"),
		QuestionID: uuid.MustParse("2713c1fc-fa4b-437a-83d3-72a8a3f068d3"),
		Text:       "There are multiple ways to create an object in JavaScript, including using the object literal notation, using the Object constructor function, and using the create() method of Object.prototype.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("5e1d9943-0ed5-43a4-b5a4-f96879d5d7d5"),
		QuestionID: uuid.MustParse("6c203f4b-5402-47e5-8c4d-7f4ab41b35b5"),
		Text:       "Synchronous programming blocks the execution of code until a task is completed, while asynchronous programming allows code to continue executing while a task is being completed.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("19d854f5-4a94-4ca7-9444-79ceef5ab9f8"),
		QuestionID: uuid.MustParse("6c203f4b-5402-47e5-8c4d-7f4ab41b35b5"),
		Text:       "In synchronous programming, each task must be completed before the next task can begin, while in asynchronous programming, multiple tasks can be completed simultaneously.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("4d8e4bcf-48fc-4a28-bd7d-83c2377f8209"),
		QuestionID: uuid.MustParse("6c203f4b-5402-47e5-8c4d-7f4ab41b35b5"),
		Text:       "Asynchronous programming is always faster than synchronous programming because it allows multiple tasks to be completed simultaneously.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("12e7b982-1c19-4699-94d2-8c21ed52327e"),
		QuestionID: uuid.MustParse("12e7b982-1c19-4699-94d2-8c21ed52327e"),
		Text:       "A programming language used to create interactive effects within web browsers",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("78f7c8a6-f7e5-40d4-8ebf-ebf7c68016d9"),
		QuestionID: uuid.MustParse("12e7b982-1c19-4699-94d2-8c21ed52327e"),
		Text:       "A type of coffee bean grown in Java, Indonesia",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("eedca8a3-2361-4284-a1e7-634234f1f0d5"),
		QuestionID: uuid.MustParse("12e7b982-1c19-4699-94d2-8c21ed52327e"),
		Text:       "A musical instrument traditionally used in Japanese folk music",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("63403396-8f7e-4f90-a2d9-bc3c8498c4b5"),
		QuestionID: uuid.MustParse("78f7c8a6-f7e5-40d4-8ebf-ebf7c68016d9"),
		Text:       "Number, String, Boolean, Null, Undefined, Object, Symbol",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("89f8e8c1-97a3-4f21-a316-05d4a9ce4263"),
		QuestionID: uuid.MustParse("78f7c8a6-f7e5-40d4-8ebf-ebf7c68016d9"),
		Text:       "Number, String, Boolean, Object, Array, Null",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("f4c5349c-3961-4e72-a9ba-cdd61c55ba31"),
		QuestionID: uuid.MustParse("78f7c8a6-f7e5-40d4-8ebf-ebf7c68016d9"),
		Text:       "Integer, Float, Double, Long, Short, Byte",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("c29e48e1-fc6f-4d58-a3c6-2cb980d36952"),
		QuestionID: uuid.MustParse("eedca8a3-2361-4284-a1e7-634234f1f0d5"),
		Text:       "Null and undefined are two terms for the same thing in JavaScript.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("b5a87cd5-d44c-4c09-9f80-cd39205504c2"),
		QuestionID: uuid.MustParse("eedca8a3-2361-4284-a1e7-634234f1f0d5"),
		Text:       "Null is a value that represents the intentional absence of any object value, while undefined is a value used to represent a missing primitive value.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("f9c4f551-68d9-44d5-a8eb-31612e6515d3"),
		QuestionID: uuid.MustParse("eedca8a3-2361-4284-a1e7-634234f1f0d5"),
		Text:       "Undefined is a value that represents the intentional absence of any object value, while null is a value used to represent a missing primitive value.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("f3a6d47e-6c1c-44b4-b924-6797ab54d0ee"),
		QuestionID: uuid.MustParse("82d5633b-b69e-4c8c-ba37-2929f0ad7473"),
		Text:       "var, let, and const",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("b4dbdd1c-94c8-4222-9836-88b17df020ea"),
		QuestionID: uuid.MustParse("82d5633b-b69e-4c8c-ba37-2929f0ad7473"),
		Text:       "declare, const, and let",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("5f6d3da7-17cf-42a9-a838-3a4f71c80f64"),
		QuestionID: uuid.MustParse("82d5633b-b69e-4c8c-ba37-2929f0ad7473"),
		Text:       "create, var, and let",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("9cdd63c6-d8b5-465e-bde4-70f39c27c5a9"),
		QuestionID: uuid.MustParse("0ed8704b-d3d6-45a3-92d3-118ce3f1381c"),
		Text:       "A closure is a function that has access to variables in its outer scope, even after the outer function has returned.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("9f31d2d7-ec1c-4ef9-8f20-7fbf738e6658"),
		QuestionID: uuid.MustParse("0ed8704b-d3d6-45a3-92d3-118ce3f1381c"),
		Text:       "A closure is a function that has no access to variables in its outer scope, even after the outer function has returned.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("2b3f121d-8e67-46dc-bccb-365af4ef2b7e"),
		QuestionID: uuid.MustParse("0ed8704b-d3d6-45a3-92d3-118ce3f1381c"),
		Text:       "A closure is a function that is always called first, before any other function in the same scope.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("e37d28e9-2462-4479-8e99-1d684cf4f63b"),
		QuestionID: uuid.MustParse("e6c7e2c6-c604-4dd5-a0dc-1010c930a6e7"),
		Text:       "== and === operators both perform type coercion before comparison.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("3e3d64a8-00a8-4ad4-8b26-47198c6b1b94"),
		QuestionID: uuid.MustParse("e6c7e2c6-c604-4dd5-a0dc-1010c930a6e7"),
		Text:       "== operator checks for value equality after type coercion, while === operator checks for value equality without type coercion.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("f9707741-f4f5-4d6e-88a9-74f7716c4da2"),
		QuestionID: uuid.MustParse("e6c7e2c6-c604-4dd5-a0dc-1010c930a6e7"),
		Text:       "=== operator checks for value equality after type coercion, while == operator checks for value equality without type coercion.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("b61797d6-61dd-452c-921a-8a7e87f1737e"),
		QuestionID: uuid.MustParse("96f47d7f-3443-4f68-af16-05b754d07e88"),
		Text:       "let variables are block-scoped, while var variables are function-scoped.",
		IsCorrect:  true,
	},
	{
		ID:         uuid.MustParse("fc75a9c7-2c18-45ba-89a1-7ca155da5b1d"),
		QuestionID: uuid.MustParse("96f47d7f-3443-4f68-af16-05b754d07e88"),
		Text:       "let variables are function-scoped, while var variables are block-scoped.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("f1379b94-6706-47e8-9f34-1ad7d284ef0f"),
		QuestionID: uuid.MustParse("96f47d7f-3443-4f68-af16-05b754d07e88"),
		Text:       "Both let and var variables are block-scoped.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("e30eb8f7-3f8a-4046-9960-989dc8f0b1c2"),
		QuestionID: uuid.MustParse("6f82f33b-35b3-4d61-bfe6-73addd7d05a9"),
		Text:       "There are two types of events in JavaScript: user-generated events and browser-generated events.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("9a487522-bd6c-4ec6-9380-2a0c684f1d5f"),
		QuestionID: uuid.MustParse("6f82f33b-35b3-4d61-bfe6-73addd7d05a9"),
		Text:       "There are three types of events in JavaScript: user-generated events, browser-generated events, and system-generated events.",
		IsCorrect:  false,
	},
	{
		ID:         uuid.MustParse("0e0de602-8f9d-40c6-89b2-f4e8a760bd5d"),
		QuestionID: uuid.MustParse("6f82f33b-35b3-4d61-bfe6-73addd7d05a9"),
		Text:       "There are many types of events in JavaScript, including mouse events, keyboard events, form events, and document/window events.",
		IsCorrect:  true,
	},
}
