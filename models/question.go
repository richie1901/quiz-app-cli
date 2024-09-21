package models

type Question struct{
	Id int `json:"id"`
	Question string `json:"question"`
	PossibleAnswers map[string]string `json:"possibleAnswers"`
}