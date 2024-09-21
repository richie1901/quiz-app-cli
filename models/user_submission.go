package models

type UserSubmissions struct{
	UserId int `json:"userId"`
	UserAnswers []UserAnswer `json:"userAnswers"`
}