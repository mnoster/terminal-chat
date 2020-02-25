// Code generated by goa v3.0.10, DO NOT EDIT.
//
// tmess HTTP client CLI support package
//
// Command:
// $ goa gen terminal-chat/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	tmessc "terminal-chat/gen/http/tmess/client"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `tmess (login|echoer|listener|summary|subscribe|history)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` tmess login --user "username" --password "password"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	tmessConfigurer *tmessc.ConnConfigurer,
) (goa.Endpoint, interface{}, error) {
	var (
		tmessFlags = flag.NewFlagSet("tmess", flag.ContinueOnError)

		tmessLoginFlags        = flag.NewFlagSet("login", flag.ExitOnError)
		tmessLoginUserFlag     = tmessLoginFlags.String("user", "REQUIRED", "")
		tmessLoginPasswordFlag = tmessLoginFlags.String("password", "REQUIRED", "")

		tmessEchoerFlags     = flag.NewFlagSet("echoer", flag.ExitOnError)
		tmessEchoerTokenFlag = tmessEchoerFlags.String("token", "REQUIRED", "")

		tmessListenerFlags     = flag.NewFlagSet("listener", flag.ExitOnError)
		tmessListenerTokenFlag = tmessListenerFlags.String("token", "REQUIRED", "")

		tmessSummaryFlags     = flag.NewFlagSet("summary", flag.ExitOnError)
		tmessSummaryTokenFlag = tmessSummaryFlags.String("token", "REQUIRED", "")

		tmessSubscribeFlags     = flag.NewFlagSet("subscribe", flag.ExitOnError)
		tmessSubscribeTokenFlag = tmessSubscribeFlags.String("token", "REQUIRED", "")

		tmessHistoryFlags     = flag.NewFlagSet("history", flag.ExitOnError)
		tmessHistoryViewFlag  = tmessHistoryFlags.String("view", "", "")
		tmessHistoryTokenFlag = tmessHistoryFlags.String("token", "REQUIRED", "")
	)
	tmessFlags.Usage = tmessUsage
	tmessLoginFlags.Usage = tmessLoginUsage
	tmessEchoerFlags.Usage = tmessEchoerUsage
	tmessListenerFlags.Usage = tmessListenerUsage
	tmessSummaryFlags.Usage = tmessSummaryUsage
	tmessSubscribeFlags.Usage = tmessSubscribeUsage
	tmessHistoryFlags.Usage = tmessHistoryUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "tmess":
			svcf = tmessFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "tmess":
			switch epn {
			case "login":
				epf = tmessLoginFlags

			case "echoer":
				epf = tmessEchoerFlags

			case "listener":
				epf = tmessListenerFlags

			case "summary":
				epf = tmessSummaryFlags

			case "subscribe":
				epf = tmessSubscribeFlags

			case "history":
				epf = tmessHistoryFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "tmess":
			c := tmessc.NewClient(scheme, host, doer, enc, dec, restore, dialer, tmessConfigurer)
			switch epn {
			case "login":
				endpoint = c.Login()
				data, err = tmessc.BuildLoginPayload(*tmessLoginUserFlag, *tmessLoginPasswordFlag)
			case "echoer":
				endpoint = c.Echoer()
				data, err = tmessc.BuildEchoerPayload(*tmessEchoerTokenFlag)
			case "listener":
				endpoint = c.Listener()
				data, err = tmessc.BuildListenerPayload(*tmessListenerTokenFlag)
			case "summary":
				endpoint = c.Summary()
				data, err = tmessc.BuildSummaryPayload(*tmessSummaryTokenFlag)
			case "subscribe":
				endpoint = c.Subscribe()
				data, err = tmessc.BuildSubscribePayload(*tmessSubscribeTokenFlag)
			case "history":
				endpoint = c.History()
				data, err = tmessc.BuildHistoryPayload(*tmessHistoryViewFlag, *tmessHistoryTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// tmessUsage displays the usage of the tmess command and its subcommands.
func tmessUsage() {
	fmt.Fprintf(os.Stderr, `The chatter service implements a simple client and server chat.
Usage:
    %s [globalflags] tmess COMMAND [flags]

COMMAND:
    login: Creates a valid JWT token for auth to chat.
    echoer: Echoes the message sent by the client.
    listener: Listens to the messages sent by the client.
    summary: Summarizes the chat messages sent by the client.
    subscribe: Subscribe to events sent when new chat messages are added.
    history: Returns the chat messages sent to the server.

Additional help:
    %s tmess COMMAND --help
`, os.Args[0], os.Args[0])
}
func tmessLoginUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tmess login -user STRING -password STRING

Creates a valid JWT token for auth to chat.
    -user STRING: 
    -password STRING: 

Example:
    `+os.Args[0]+` tmess login --user "username" --password "password"
`, os.Args[0])
}

func tmessEchoerUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tmess echoer -token STRING

Echoes the message sent by the client.
    -token STRING: 

Example:
    `+os.Args[0]+` tmess echoer --token "Fugiat sed sequi consequatur."
`, os.Args[0])
}

func tmessListenerUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tmess listener -token STRING

Listens to the messages sent by the client.
    -token STRING: 

Example:
    `+os.Args[0]+` tmess listener --token "Cum tempore incidunt vel nulla aut."
`, os.Args[0])
}

func tmessSummaryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tmess summary -token STRING

Summarizes the chat messages sent by the client.
    -token STRING: 

Example:
    `+os.Args[0]+` tmess summary --token "Ipsam non exercitationem reprehenderit itaque ea."
`, os.Args[0])
}

func tmessSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tmess subscribe -token STRING

Subscribe to events sent when new chat messages are added.
    -token STRING: 

Example:
    `+os.Args[0]+` tmess subscribe --token "Accusantium similique voluptatem voluptas facere beatae non."
`, os.Args[0])
}

func tmessHistoryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tmess history -view STRING -token STRING

Returns the chat messages sent to the server.
    -view STRING: 
    -token STRING: 

Example:
    `+os.Args[0]+` tmess history --view "Ut consequatur." --token "Voluptatem rem."
`, os.Args[0])
}
