package models

type RecentUserScoreBoard struct{
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	PercentageScore float64 `json:"percentageScore"`
}


// func generateRecentUsersScoreBoard() []RecentUserScoreBoard {
// 	return []RecentUserScoreBoard{
// 		{UserId: 1, UserName: "kfolly@32", PercentageScore: 32.65},
// 		{UserId: 2, UserName: "fgdgsy@32", PercentageScore: 12.76},
// 		{UserId: 3, UserName: "tury@32", PercentageScore: 86.54},
// 		{UserId: 4, UserName: "tyrh4y@32", PercentageScore: 62.31},
// 	}

// }