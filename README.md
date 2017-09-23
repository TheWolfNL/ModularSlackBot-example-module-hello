# ModularSlackBot-example-module-hello
Fork this repo to get started making your own [ModularSlackBot](https://thewolfnl.github.io/ModularSlackBot/) module

# Where to start
1. Install dependencies using [Glide](https://github.com/Masterminds/glide#install) by running `glide install`
1. First change the package name in `example.go` and `example_test.go`
1. Change the file names to match the package name
1. Add [triggers](https://thewolfnl.github.io/ModularSlackBot/triggers.html)
1. You can add tests by copying and altering `func ExampleHello()` in the `example_test.go` file.
1. To run tests call `go test -v` in commandline

## Getting your module listed
If you've created a public module you can get it listed on [this page](https://thewolfnl.github.io/ModularSlackBot/modules.html).
Just create an [issue](https://github.com/TheWolfNL/ModularSlackBot/issues/new) with the label `module-listing` and we'll add it as soon as possible. 