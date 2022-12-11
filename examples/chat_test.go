package examples

import (
	"fmt"
	"log"
)

func ExampleClient_GetChatText() {
	message := "hello"

	// chat in independent conversation
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}
	fmt.Printf("q: %s, a: %s\n", message, text.Content)

	// Output: xxx
}

func ExampleClient_GetChatText_WithConversation() {
	message := "what's golang"

	// chat in independent conversation
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}
	fmt.Printf("q: %s, a: %s\n", message, text.Content)

	// continue conversation with new message
	conversationID := text.ConversationID
	parentMessage := text.MessageID
	newMessage := "use it to write hello world"

	newText, err := cli.GetChatText(newMessage, conversationID, parentMessage)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}
	fmt.Printf("q: %s, a: %s\n", newMessage, newText.Content)

	// Output: xxx
}
