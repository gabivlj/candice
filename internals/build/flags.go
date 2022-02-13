package build

import (
	"errors"
	"os"
)

type Flags struct {
	Path    string
	Mode    string
	Release bool
}

func retrieveFlags() (Flags, error) {
	flagsToReturn := Flags{}
	flags := os.Args
	if len(flags) < 3 {
		return flagsToReturn, errors.New("not enough arguments")
	}
	mode := flags[1]
	path := flags[2]
	for _, fl := range flags[2:] {
		if fl == "--release" {
			flagsToReturn.Release = true
		}
	}

	flagsToReturn.Mode = mode
	flagsToReturn.Path = path
	return flagsToReturn, nil
}
