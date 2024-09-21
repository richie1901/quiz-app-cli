package models

type CorrectAnswer struct{
	QuestionId int `json:"questionId"`
	CorrectAnswer string `json:"correctAnswer"`
	AnswerDetails string `json:"answerDetails"`
}

type UserAnswer struct{
	QuestionId int `json:"questionId"`
	UserSelection string `json:"correctAnswer"`
}