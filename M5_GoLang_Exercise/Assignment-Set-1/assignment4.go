package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Question struct to represent a quiz question
type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

// Updated Question Bank
var questions = []Question{
	{"Which data type is used to store true or false values in Go?", [4]string{"boolean", "bool", "bit", "flag"}, 2},
	{"What is the default value of an uninitialized string in Go?", [4]string{"nil", "undefined", "empty string", "null"}, 3},
	{"Which Go keyword is used to create a constant?", [4]string{"const", "define", "static", "immutable"}, 1},
	{"What is the default visibility of variables declared without any access specifier in Go?", [4]string{"Public", "Private", "Protected", "Package-private"}, 2},
	{"Which function is used to print output to the console in Go?", [4]string{"print()", "log()", "fmt.Println()", "output()"}, 3},
	{"How do you declare an array in Go?", [4]string{"array<int> a", "int[] a", "[5]int a", "array[5]int a"}, 3},
	{"Which package is used for string manipulation in Go?", [4]string{"text", "fmt", "strings", "strconv"}, 3},
	{"How can you convert a string to an integer in Go?", [4]string{"int.Parse()", "strconv.Atoi()", "strings.Atoi()", "fmt.Atoi()"}, 2},
	{"What does the 'defer' keyword do in Go?", [4]string{"Stops execution", "Executes after a function returns", "Immediately calls a function", "Skips a function"}, 2},
	{"What is the purpose of the 'recover' function in Go?", [4]string{"To restart a goroutine", "To handle panic gracefully", "To allocate memory", "To stop an error"}, 2},
}

// Timer duration for each question
const questionTimeLimit = 15 * time.Second

func main() {
	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("Instructions:")
	fmt.Println("- Enter the option number to select an answer.")
	fmt.Println("- Type 'exit' to quit the quiz anytime.")
	fmt.Println("- You have a limited time of 15 seconds per question.")

	takeQuiz()
}

// takeQuiz handles the quiz process
func takeQuiz() {
	reader := bufio.NewReader(os.Stdin)
	score := 0

	for i, question := range questions {
		fmt.Printf("Question %d: %s\n", i+1, question.Question)
		for idx, option := range question.Options {
			fmt.Printf("%d. %s\n", idx+1, option)
		}

		answerChan := make(chan string)
		go func() {
			fmt.Print("Your answer: ")
			input, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(input)
		}()

		select {
		case answer := <-answerChan:
			if strings.ToLower(answer) == "exit" {
				fmt.Println("Exiting the quiz...")
				return
			}

			choice, err := strconv.Atoi(answer)
			if err != nil || choice < 1 || choice > 4 {
				fmt.Println("Invalid input! Please enter a valid option (1-4).")
				continue
			}

			if choice == question.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Incorrect!")
			}

		case <-time.After(questionTimeLimit):
			fmt.Println("\nTime's up for this question!")
			continue
		}

		fmt.Println()
	}

	calculateScore(score, len(questions))
}

// calculateScore calculates and displays the user's score and performance
func calculateScore(score, total int) {
	fmt.Printf("\nQuiz Completed! Your score: %d/%d\n", score, total)
	percentage := (float64(score) / float64(total)) * 100

	if percentage >= 80 {
		fmt.Println("Performance: Excellent!")
	} else if percentage >= 50 {
		fmt.Println("Performance: Good!")
	} else {
		fmt.Println("Performance: Needs Improvement.")
	}
}