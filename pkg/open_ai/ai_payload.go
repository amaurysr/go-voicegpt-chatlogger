package open_ai 

import (
        "bufio"
        "fmt"
	"log"
        "os"
        "context"
        openai "github.com/sashabaranov/go-openai"
)
 
func Input()string{
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Printf("Enter your openAI secret key: ")
        scanner.Scan()
	fmt.Printf("\n")
        input := scanner.Text()
	return input
}

func OpenAI_chatGPT(){
        client := openai.NewClient(Input())
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

func OpenAI_Whisper(){
	client := openai.NewClient(Input())
	resp, err := client.CreateTranscription(
		context.Background(),
		openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: "file.wav",
		},
	)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	e := os.Remove("file.wav")
	if e != nil{
		log.Fatal(e)
	}
	fmt.Println(resp.Text)

}
