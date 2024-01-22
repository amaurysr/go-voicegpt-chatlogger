package main

import (
	//"fmt"
	"log"
	"os"
	"github.com/amaurysr/go-voicegpt-chatlogger/pkg/open_ai"
)

func main(){
	audiorecord()
	open_ai.OpenAI_Whisper()
	e:= os.Remove("file.wav")
	if e != nil{
		log.Fatal(e)
	}
}
