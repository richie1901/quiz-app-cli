package models


type QuizAnswer struct{
	CorrectAnswer []CorrectAnswer `json:"correctAnswers"`
}