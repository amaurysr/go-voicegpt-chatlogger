package main

import (
	"bufio" 
	"fmt"
	"os"
	"context"
	openai "github.com/sashabaranov/go-openai"
)

func openAI_Auth(input string){
	client := openai.NewClient("your token")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
	if err != nil{
		fmt.Printf("Login error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Message.Content)
	
}

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your openAI secret key: ") 
	scanner.Scan()
	input := scanner.Text()
	openAI_Auth(input)

}
