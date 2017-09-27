package example

import (
	"testing"

	"github.com/thewolfnl/ModularSlackBot"
)

func TestNew(t *testing.T) {
	module := New()

	if module.Name() != "TestModule" {
		t.Error("Expected module name TestModule, got ", module.Name())
	}

	if module.Version() != "0.0.1" {
		t.Error("Expected module version 0.0.1, got ", module.Version())
	}
}

func ExampleHello() {
	module := New()
	module.HandleInput(bot.CreateMessage("Hello"))
	// Output: Sending to #C2147483705
	// Response: Hey there, to you too
	// Message not sent to slack because slack api is not configured
}

func ExampleTest() {
	module := New()
	module.HandleInput(bot.CreateMessage("test"))
	// Output: MessageJSON:
	// {"channel":"C2147483705","user":"U2147483697","text":"test","pinned_to":null}
	//
	// Sending to #C2147483705
	// Response: Read you loud and clear..
	// Message not sent to slack because slack api is not configured
}
