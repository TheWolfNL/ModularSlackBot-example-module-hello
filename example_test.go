package example

import (
	"testing"

	"github.com/thewolfnl/ModularSlackBot"
	"github.com/thewolfnl/ModularSlackBot/testLib"
)

func setup() (*botTestLib.MockBot, *bot.Module) {
	module := New()
	mock := botTestLib.NewMockBot()
	mock.AddModule(module)
	return mock, module
}

func TestNew(t *testing.T) {
	_, module := setup()

	if module.Name() != "TestModule" {
		t.Error("Expected module name TestModule, got ", module.Name())
	}

	if module.Version() != "0.0.1" {
		t.Error("Expected module version 0.0.1, got ", module.Version())
	}
}

func ExampleHello() {
	mock, module := setup()
	message := mock.CreateMessage("Hello")

	module.HandleInput(message)
	// Output: Sending to #C2147483705
	// Response: Hey John, how you doing?
	// Message not sent to slack because slack api is not configured
}

func ExampleHelloIM() {
	mock, module := setup()
	message := mock.MockMessage(mock.CreateMessageEvent("Hello"), botTestLib.ReturnFalse, botTestLib.ReturnTrue, botTestLib.ReturnFalse)

	module.HandleInput(message)
	// Output: Sending to #C2147483705
	// Response: Hey buddy, long time no see
	// Message not sent to slack because slack api is not configured
}

func ExampleTest() {
	mock, module := setup()
	module.HandleInput(mock.CreateMessage("test"))
	// Output: MessageJSON:
	// {"type":"message","channel":"C2147483705","user":"U2147483697","text":"test","ts":"1355517523.000005","pinned_to":null}
	//
	// Sending to #C2147483705
	// Response: Read you loud and clear..
	// Message not sent to slack because slack api is not configured
}
