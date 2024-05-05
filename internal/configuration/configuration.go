package configuration

import (
	"bytes"
	"errors"
	"flag"
)

type Configuration struct {
	Address     string
	Destination string
}

func Load(args []string) (*Configuration, error) {

	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)

	errMsg := bytes.NewBufferString("")
	flagSet.SetOutput(errMsg)

	address := flagSet.String(
		"address",
		":8080",
		"Address that the server listens on",
	)

	destination := flagSet.String(
		"destination",
		".",
		"Directory in which to store files",
	)

	err := flagSet.Parse(args[1:])
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil, &HelpError{usage: errMsg.String()}
		}
		return nil, errors.New(errMsg.String())
	}

	return &Configuration{
		Address:     *address,
		Destination: *destination,
	}, nil
}
