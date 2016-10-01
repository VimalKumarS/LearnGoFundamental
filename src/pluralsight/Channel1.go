package main

import (
	"fmt"
	"strings"
)

func main() {
	/*ch := make(chan string)
	ch <- "Hello"
	fmt.Println(<-ch)*/

	phrase := "These are the times that try men's souls. \n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch)
	/*for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}*/

	for msg := range ch {
		/*if msg, ok := <-ch; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}*/
		fmt.Print(msg + " ")
	}

	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)
	msg := Message{
		To:      []string{"fr@gg.com"},
		From:    "bb@hh.com",
		Content: "content",
	}
	failedMsg := FailedMessage{
		ErrorMessage:    "Error Message",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	//errCh <- failedMsg
	// to handle multipl eblock of code
	select {
	case revdMsg := <-msgCh:
		errCh <- failedMsg
		fmt.Println(revdMsg)
	case rcedError := <-errCh:
		fmt.Println(rcedError)
	default:
		fmt.Println("NO message found")
	}
}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
