package dtos

type UserClaim struct{
	ID 			string `json:"id" binding:"required"`
	Email 		string `json:"email" binding:"required"`
	Role 		string `json:"role" binding:"required"`
}