package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shreyans-sureja/chatgpt-api/constants"
	"github.com/shreyans-sureja/chatgpt-api/services"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to chatgpt project")
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error while loading the environment variables")
	}

	chatgptApiKey := os.Getenv("CHATGPT_API_KEY")
	if len(chatgptApiKey) == 0 {
		log.Fatalln("API key is missing")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Write (quit) to exit: ")
		scanner.Scan()
		text := scanner.Text()
		if text == "quit" {
			fmt.Println("Closing the chatgpt api")
			break
		}
		payload := createGPTPayload(text)
		answer, err := services.ChatgptAPICall(payload)
		if err != nil {
			log.Fatalln("Error while calling chat gpt api: ", err)
		}
		if len(answer.Choices) != 0 && len(answer.Choices[0].Text) != 0 {
			fmt.Println("answer: ", answer.Choices[0].Text)
		} else {
			log.Println("Chatgpt gave unusal response: ", answer)
		}

	}
}

func createGPTPayload(text string) services.ChatgptPayload {
	return services.ChatgptPayload{
		Model:       constants.CHATGPT_TEXT_MODEL,
		Prompt:      text,
		Temperature: 0,
		MaxTokens:   100,
	}
}
