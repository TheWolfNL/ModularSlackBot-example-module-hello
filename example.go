package example

import (
	"fmt"

	"github.com/thewolfnl/ModularSlackBot"
)

// New function to return a new instance of this bot
func New() *bot.Module {
	module := bot.NewModule("TestModule", "0.0.2")

	// Define triggers
	module.AddTrigger("(?i)(hello|hi|hey).*", hello)

	module.AddTrigger("test", func(message *bot.Message) {
		fmt.Printf("\nMessageJSON:\n%s\n", message.ToJson())
		message.Respond("Read you loud and clear..")
	})

	return module
}

func hello(message *bot.Message) {
	if message.IsIM() {
		message.Respond("Hey buddy, long time no see")
	} else {
		user := message.GetUser()
		name := "John"
		if user != nil {
			name = user.Name
		}
		message.Respond(fmt.Sprintf("Hey %s, how you doing?", name))
	}
}
