package dtos

type RegisterUser struct{
	FullName 	string `json:"full_name"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}