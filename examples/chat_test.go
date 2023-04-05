package examples

import (
	"fmt"
	"log"
)

func ExampleClient_GetChatText() {
	cli := getClient()

	message := "say hello to me"

	log.Printf("start get chat text")

	// chat in independent conversation
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("\nq: %s\na: %s\n", message, text.Content)

	fmt.Println("xxx")
	// Output: xxx
}

func ExampleClient_GetContinuousChatText() {
	cli := getClient()

	message := "say hello to me"

	log.Printf("start get chat text")

	// chat in independent conversation
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("\nq: %s\na: %s\n", message, text.Content)

	log.Printf("start get chat text again")

	// continue conversation with new message
	conversationID := text.ConversationID
	parentMessage := text.MessageID
	newMessage := "again"

	newText, err := cli.GetChatText(newMessage, conversationID, parentMessage)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("\nq: %s\na: %s\n", newMessage, newText.Content)

	fmt.Println("xxx")
	// Output: xxx
}

func ExampleClient_GetChatStream() {
	cli := getClient()

	message := "say hello to me"

	log.Printf("start get chat stream")

	stream, err := cli.WithModel("gpt-4").GetChatStream(message)
	if err != nil {
		log.Fatalf("get chat stream failed: %v\n", err)
	}

	var answer string
	for text := range stream.Stream {
		log.Printf("stream text: %s\n", text)
		answer = text.Content
	}

	if stream.Err != nil {
		log.Fatalf("stream closed with error: %v\n", stream.Err)
	}

	log.Printf("\nq: %s\na: %s\n", message, answer)

	fmt.Println("xxx")
	// Output: xxx
}
