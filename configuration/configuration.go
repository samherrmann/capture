package configuration

import (
	"flag"
)

type Configuration struct {
	Address     string
	Destination string
}

func Load(args []string) (*Configuration, error) {

	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)

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
		return nil, err
	}

	return &Configuration{
		Address:     *address,
		Destination: *destination,
	}, nil
}
