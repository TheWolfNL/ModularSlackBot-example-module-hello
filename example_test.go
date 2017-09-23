package example

import (
	"encoding/json"
	"testing"

	"github.com/thewolfnl/ModularSlackBot/bot"
)

func TestNew(t *testing.T) {
	module := New()

	if module.Name() != "TestModule" {
		t.Error("Expected version 0.0.1, got ", module.Name())
	}

	if module.Version() != "0.0.1" {
		t.Error("Expected version 0.0.1, got ", module.Version())
	}
}

func ExampleHello() {
	module := New()
	message, err := createMessage("Hello")
	if err != nil {
		// exit
	}
	module.HandleInput(message)
	// Output: Sending to #C2147483705
	// Response: Hey there, to you too
	// Message not sent to slack because slack api is not configured
}

func ExampleTest() {
	module := New()
	message, err := createMessage("test")
	if err != nil {
		// exit
	}
	module.HandleInput(message)
	// Output: MessageJSON:
	// {"type":"message","channel":"C2147483705","user":"U2147483697","text":"test","ts":"1355517523.000005","pinned_to":null}
	//
	// Sending to #C2147483705
	// Response: Read you loud and clear..
	// Message not sent to slack because slack api is not configured
}

func createMessage(messageString string) (bot.Message, error) {
	message := bot.Message{}
	messageJson := `{
		"type": "message",
		"channel": "C2147483705",
		"user": "U2147483697",
		"text": "Hello world",
		"ts": "1355517523.000005"
		}`
	if err := json.Unmarshal([]byte(messageJson), &message); err != nil {
		return bot.Message{}, err
	}
	message.Text = messageString
	return message, nil
}

func TestCreateMessage(t *testing.T) {
	_, err := createMessage("Hello")
	if err != nil {
		t.Error("Could not create a test message")
	}
}
