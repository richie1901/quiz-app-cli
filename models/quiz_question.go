package models

type QuizQuestion struct{
	UserId int `json:"userId"`
	Questions []Question `json:"questions"`
}

