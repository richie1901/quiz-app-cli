package commands

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"richard_adekponya_fasttrack_cli_quizapp.com/app/constants"
	"richard_adekponya_fasttrack_cli_quizapp.com/app/models"
	"richard_adekponya_fasttrack_cli_quizapp.com/app/utils"
	"strconv"
	"sort"
)

var GetQuestionsCmd = &cobra.Command{
	Use:   "get-questions",
	Short: "Get quiz questions",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(utils.BaseURL + "/user/get-questions")
		if err != nil {
			log.Println("Error fetching questions:", err)
			return
		}
		defer resp.Body.Close()
		var quizQuestionResponse models.QuizQuestionResponse

		if err := json.NewDecoder(resp.Body).Decode(&quizQuestionResponse); err != nil {
			log.Println("Error fetching questions:", err)
			return
		}
		if quizQuestionResponse.ResponseCode == constants.SUCCESS_CODE {
			lengthOfQuestions := len(quizQuestionResponse.QuizQuestion.Questions)
			questions := quizQuestionResponse.QuizQuestion.Questions
			questionBuilder := ""
			if lengthOfQuestions == 0 {
				log.Println("There are no active questions for the quiz ")
			} else {
				for _, question := range questions {
					possibleAnswerBuilder := ""
					// for key, value := range question.PossibleAnswers {

					// 	possibleAnswerBuilder += key + ". " + value + " \n"
					// }
					keys := make([]string, 0)
					for k:= range question.PossibleAnswers {
						keys = append(keys, k)
					}
					sort.Strings(keys)
					for _, k := range keys {
						possibleAnswerBuilder += k + ". " + question.PossibleAnswers[k] + " \n"
					}
					questionBuilder += "\n" + strconv.Itoa(question.Id) + ") " + question.Question + "\n" + possibleAnswerBuilder
				}
				log.Println(questionBuilder)
			}

		}

		// log.Println("found questions to show to client ",quizQuestionResponse)
	},
}



