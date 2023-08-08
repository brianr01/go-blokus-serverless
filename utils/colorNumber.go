package utils

import "github.com/brianr01/go-blockus-serverless/types"

func ColorNumberInColorNumbers(c types.ColorNumber, cs []types.ColorNumber) bool {
	for _, colorToCheck := range cs {
		if c == colorToCheck {
			return true
		}
	}

	return false
}
