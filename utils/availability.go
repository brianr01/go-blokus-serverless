package utils

import "github.com/brianr01/go-blockus-serverless/types"

func AvailabilityNumberInAvailabilityNumbers(an types.AvailabilityNumber, ans []types.AvailabilityNumber) bool {
	for _, anToCheck := range ans {
		if anToCheck == an {
			return true
		}
	}

	return false
}

func AvailabilityNumbersWithMinium(anMin types.AvailabilityNumber, ans []types.AvailabilityNumber) bool {
	for _, anToCheck := range ans {
		if anToCheck >= anMin {
			return true
		}
	}

	return false
}
