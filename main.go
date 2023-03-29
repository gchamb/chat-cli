package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

/*
	ChatGPT but for my computer as a cli tool.

	Steps:
	1. Take in a prompt from CLI
	2. Use OpenAI API
	3. Display Response
*/

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("OPEN_AI_KEY")
	client := openai.NewClient(API_KEY)
	ctx := context.Background()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("Enter in your prompt below or q - to quit ⬇️")
		fmt.Println("--------------------------------------------")
		

		scanner.Scan()
		prompt := scanner.Text()

		if strings.ToLower(prompt) == "q" {
			break
		}

		if prompt == ""{
			log.Fatal("Prompt shouldn't be empty.")
		}

		// create request structure
		req := openai.CompletionRequest{
			Model:     openai.GPT3TextDavinci003,
			MaxTokens: 1000,
			Prompt:    prompt,
		}

		// make request
		res, err := client.CreateCompletion(ctx, req)
		if err != nil {
			log.Fatal(err)
		}

		// display model result
		fmt.Println(res.Choices[0].Text)
	}

}
