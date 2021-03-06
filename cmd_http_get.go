package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/skx/subcommands"
)

// Structure for our options and state.
type httpGetCommand struct {

	// We embed the NoFlags option, because we accept no command-line flags.
	subcommands.NoFlags
}

// Info returns the name of this subcommand.
func (hg *httpGetCommand) Info() (string, string) {
	return "http-get", `Fetch a remote URL

Details:

This command is very much curl-lite, allowing you to fetch the contents of
a remote URL, with no configuration options of any kind.

While it is unusual to find hosts without curl or wget installed it does
happen, this command will bridge the gap a little.

Examples:

$ sysbox http-get https://steve.fi/`
}

// Execute is invoked if the user specifies `http-get` as the subcommand.
func (hg *httpGetCommand) Execute(args []string) int {

	// Ensure we have only a single URL
	if len(args) != 1 {
		fmt.Printf("Usage: http-get URL\n")
		return 1
	}

	// Make the request
	response, err := http.Get(args[0])
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return 1
	}

	// Get the body.
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return 1
	}

	// All OK
	fmt.Printf("%s\n", string(contents))
	return 0
}
