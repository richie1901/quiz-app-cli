package models

type FinalQuizResponse struct{
	ResponseCode string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	TotalQuestions int `json:"totalQuestions"`
	TotalCorrectAnswers int `json:"totalCorrectAnswers"`
	PercentageScoreInQuiz float64 `json:"percentageScoredInQuiz"`
	RecentUsersScoreBoard []RecentUserScoreBoard `json:"recentUserScoreBoard"`
	PercentageOfUsersPerformedBetterThan float64 `json:"percentageOfUsersPerformedBetterThan"`
}

type ApiErrorResponse struct{
	ResponseCode string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ErrorMessage string `json:"errorMessage"`
}

type QuizQuestionResponse struct{
	ResponseCode string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	QuizQuestion QuizQuestion `json:"quizQuestion"`
}