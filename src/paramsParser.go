package main

import (
	"fmt"
)

// ParseParameters Parse given array into uri and flags map
func ParseParameters(rawArgs []string) (uri string, flags map[string]string, err error) {
	uri = parseURI(rawArgs)
	flags, err = parseFlags(rawArgs)

	if err != nil {
		return "", nil, err
	}

	return uri, flags, nil
}

func parseURI(rawArgs []string) string {
	return rawArgs[lastPosition(rawArgs)]
}

func parseFlags(rawArgs []string) (map[string]string, error) {
	if len(rawArgs) <= 2 {
		return map[string]string{}, nil
	}

	flagsSequence := rawArgs[1:lastPosition(rawArgs)]

	flags := make(map[string]string)

	for i := 0; i < len(flagsSequence); {
		key, value, hopSize, err := parseKeyValueParam(flagsSequence, i)

		if err != nil {
			return nil, err
		}

		flags[key] = value

		i += hopSize
	}

	return flags, nil
}

const hopSizeValuedFlag = 2
const hopSizeSilentFlag = 1

func parseKeyValueParam(flags []string, keyPosition int) (key string, value string, hopSize int, err error) {
	valuePosition := keyPosition + 1

	key = flags[keyPosition]
	value = ""
	hopSize = hopSizeSilentFlag

	if key[0] != '-' {
		return "", "", -1, fmt.Errorf("Argument '%s' is invalid (must start with '-')", key)
	}

	key = key[1:]

	if len(flags) > valuePosition &&
		flags[valuePosition][0] != '-' {
		value = flags[valuePosition]
		hopSize = hopSizeValuedFlag
	}

	return key, value, hopSize, nil
}

func lastPosition(rawArgs []string) int {
	return len(rawArgs) - 1
}
