# Spirit Quiz

Spirit Quiz is an online platform for taking tests in certain programming languages. This project was built as a part of the <b>Build With Vee 100 days of coding challenge</b>.
https://twitter.com/build_vee

## Demo
[Click here to watch the video](https://www.awesomescreenshot.com/video/17685268?key=a25d8e855a6e2111f7d8e7191691cf58)


<a href="https://documenter.getpostman.com/view/16535122/2s93m1b4x2">API Documentation</a>

## Endpoints

### Auth
- POST /auth/register - Register a new user
- POST /auth/verify-email - Verify user's email address
- POST /auth/login - Login with existing credentials
- POST /auth/request-change-password - Request a password reset link
- POST /auth/reset-password - Reset user's password

### Image
- POST /image/upload - Upload an image
- DELETE /image/delete - Delete an image

### User
- PATCH /user/:id - Update user's information
- GET /user/:id - Get user's information by ID
- GET /user/users - Get all users
- DELETE /user/:id - Delete user by ID

### Quiz
- POST /quiz/create/category - Create a new quiz category
- POST /quiz/create/question - Create a new quiz question
- POST /quiz/create/option - Create a new quiz option
- GET /quiz/categories - Get all quiz categories
- PATCH /quiz/categories/:id - Update quiz category by ID
- DELETE /quiz/categories/:id - Delete quiz category by ID
- GET /quiz/questions - Get all quiz questions
- GET /quiz/questions/:id - Get all quiz questions by category ID
- PATCH /quiz/questions/:id - Update quiz question by ID
- DELETE /quiz/questions/delete/:id - Delete quiz question by ID
- GET /quiz/options - Get all quiz options
- GET /quiz/options/:id - Get all quiz options by question ID
- PATCH /quiz/options/:id - Update quiz option by ID
- DELETE /quiz/options/:id - Delete quiz option by ID

### Result
- POST /create - Create Result of test with user's information and score
- GET /user/:id - Get user's Result by user_id
- GET /category/id - Get all results in a test category

## Technologies Used
- Golang
- Jwt
- Bcrypt
- Gin (Gin Gonic)
- PostgreSQl
- Render
- Cloudinary
- Gomailer

## Installation
1. Clone the repository
2. Install the dependencies
   ```bash
   go get ./...
   ```

## Start the server

```bash
go run main.go

```