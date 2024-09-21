package commands

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"richard_adekponya_fasttrack_cli_quizapp.com/app/constants"
	"richard_adekponya_fasttrack_cli_quizapp.com/app/models"
	"richard_adekponya_fasttrack_cli_quizapp.com/app/utils"

	"github.com/spf13/cobra"
)


func getActiveUserId(userStore map[int]int)int{
	 for _,k:=range userStore{
		return k
	 }
	 return 0
}

var SubmitAnswersCmd = &cobra.Command{
	Use:   "submit-answers",
	Short: "Submit quiz answers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)==0{
			log.Println("Invalid Selection. RetakeQuiz")
			return
		}
		userAnswers := []models.UserAnswer{}
		userAnswer := models.UserAnswer{}
		i := 1
		for _, arg := range args {
			// answer, err := arg
			// if err != nil {
			// 	fmt.Println("Invalid answer:", err)
			// 	return
			// }
			userAnswer.QuestionId = i
			userAnswer.UserSelection = strings.ToUpper(arg)
			// answers = append(answers, arg)
			userAnswers = append(userAnswers, userAnswer)
			i++
		}
		var userId=getActiveUserIdFromQuestionsTaken(utils.BaseURL+"/user/get-questions")

		userSubmissions := models.UserSubmissions{UserId: userId, UserAnswers: userAnswers}
		log.Println("post request of user answer populated ", userSubmissions)
		jsonData, _ := json.Marshal(userSubmissions)

		resp, err := http.Post(utils.BaseURL+"/user/submit-answers", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error submitting answers:", err)
			return
		}
		defer resp.Body.Close()

		var finalQuizResponse models.FinalQuizResponse

		if err := json.NewDecoder(resp.Body).Decode(&finalQuizResponse); err != nil {
			log.Println("Error fetching final quiz response from submit api:", err)
			return
		}
		if finalQuizResponse.ResponseCode == constants.SUCCESS_CODE {
			quizResponseBuilder := "\n============================================\nBelow are the results from your quiz taken\n============================================"
			totalQuestions := finalQuizResponse.TotalQuestions
			totalCorrectAnswers:=finalQuizResponse.TotalCorrectAnswers
			percentageScoredInQuiz:=finalQuizResponse.PercentageScoreInQuiz
			percentageOfUsersBetterThan:=finalQuizResponse.PercentageOfUsersPerformedBetterThan
			totalUsersTakenQuiz:=len(finalQuizResponse.RecentUsersScoreBoard)
			userScoreBoard:=generateUserScoreBoard(finalQuizResponse.RecentUsersScoreBoard,userId,totalUsersTakenQuiz)
			quizResponseBuilder+="\nTotal Questions: "+strconv.Itoa(totalQuestions)+"\nTotal Correct Answers: "+strconv.Itoa(totalCorrectAnswers)+"\nPercentage Score In Quiz :"+strconv.FormatFloat(percentageScoredInQuiz,'f', -1, 64)+"%\n"+"Total Quiz Participants You Inclusive: "+strconv.Itoa(totalUsersTakenQuiz)+"\n"+"You were better than "+strconv.FormatFloat(percentageOfUsersBetterThan,'f', -1, 64)+"%"+" of all quizzers"+userScoreBoard
			log.Println(quizResponseBuilder)
		}else{
			failedQuizResponseBuilder := "\nThere Were Some Invalid Choices In Your Possible Answers Selections. Ensure You Choose Possible Answers Seperating Each Question With A Space Egs. A B C D For A List Of 4 Answers To 4 Questions In An Order. Any Other Is Void And Quiz Must Be Retaken\n"
			log.Println(failedQuizResponseBuilder)
		}

	},
}

func generateUserScoreBoard(recentUserScoreBoard []models.RecentUserScoreBoard, id int,totalQuizzers int) string{
	builder:="\n========================================================\n\t\tUser Score Board\n========================================================\nUserName\tUserId\t\tScore\t\tRank\n========================================================\n"
	userId:=""
	for _,users:=range recentUserScoreBoard{
		if users.UserId==id{
			userId=strconv.Itoa(users.UserId)+" (You)"
		}else{
			userId=strconv.Itoa(users.UserId)
		}
	builder+=users.UserName+"\t"+userId+"\t\t"+strconv.FormatFloat(users.PercentageScore,'f',-1,64)+"%\t\t"+generateRankingForQuizzers(recentUserScoreBoard,users,totalQuizzers)+"\n"
	}
	return builder
}


func getActiveUserIdFromQuestionsTaken(url string)int{
	resp, err := http.Get(url)
		if err != nil {
			log.Println("Error fetching questions:", err)
			return 1
		}
		defer resp.Body.Close()
		var quizQuestionResponse models.QuizQuestionResponse

		if err := json.NewDecoder(resp.Body).Decode(&quizQuestionResponse); err != nil {
			log.Println("Error fetching questions:", err)
			return 1
		}
		if quizQuestionResponse.ResponseCode == constants.SUCCESS_CODE {
			return quizQuestionResponse.QuizQuestion.UserId
		}
		return 1
}

func generateRankingForQuizzers(recentScoreBoarders []models.RecentUserScoreBoard, userBoard models.RecentUserScoreBoard,totalQuizzers int)string{
	usersPassed:=0
	totalUsers:=totalQuizzers
	for _,user:=range recentScoreBoarders{
		if user.PercentageScore<=userBoard.PercentageScore&&!(user.UserId==userBoard.UserId){
			usersPassed++
		}
	}
	position:=totalUsers-usersPassed
	if position==0{
		position++
	}
	rank:=strconv.Itoa(position)
	if strings.HasSuffix(rank,"1"){
		return rank+"st"
	} 
	if strings.HasSuffix(rank,"2"){
		return rank+"nd"
	} 
	if strings.HasSuffix(rank,"3"){
		return rank+"rd"
	} 
	return rank+"th"

}
