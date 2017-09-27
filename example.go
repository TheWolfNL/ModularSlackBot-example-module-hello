package example

import (
	"fmt"

	"github.com/thewolfnl/ModularSlackBot"
)

type TestModule struct {
	*bot.Module
}

// New function to return a new instance of this bot
func New() *TestModule {
	module := TestModule{bot.NewModule("TestModule", "0.0.1")}

	// Define triggers
	module.AddTrigger("(?i)(hello|hi|hey).*", module.hello)

	module.AddTrigger("test", func(message *bot.Message) error {
		fmt.Printf("\nMessageJSON:\n%s\n", message.ToJson())
		module.Respond("Read you loud and clear..")
		return nil
	})

	return &module
}

func (module *TestModule) hello(message *bot.Message) error {
	module.Respond("Hey there, to you too")
	return nil
}
