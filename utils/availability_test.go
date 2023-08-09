package utils_test

import (
	"testing"

	"github.com/brianr01/go-blockus-serverless/constants"
	"github.com/brianr01/go-blockus-serverless/types"
	"github.com/brianr01/go-blockus-serverless/utils"
)

func TestAvailabilityNumberInAvailabilityNumbers(t *testing.T) {
	test := []struct {
		name                string
		availabilityNumber  types.AvailabilityNumber
		availabilityNumbers []types.AvailabilityNumber
		want                bool
	}{
		{
			name:               "Returns true when first matches",
			availabilityNumber: constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberPlayable,
			},
			want: true,
		},
		{
			name:               "Returns true when last matches",
			availabilityNumber: constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberEmpty,
				constants.AvailabilityNumberOccupied,
				constants.AvailabilityNumberValid,
				constants.AvailabilityNumberPlayable,
			},
			want: true,
		},
		{
			name:               "Returns false with no matches",
			availabilityNumber: constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberEmpty,
				constants.AvailabilityNumberOccupied,
				constants.AvailabilityNumberEmpty,
				constants.AvailabilityNumberValid,
			},
			want: false,
		}, {
			name:                "Returns false with empty",
			availabilityNumber:  constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{},
			want:                false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.AvailabilityNumberInAvailabilityNumbers(tt.availabilityNumber, tt.availabilityNumbers)
			if result != tt.want {
				t.Errorf("The output for test %s did not match.\nExpected: %v\nActual: %v\n", tt.name, tt.want, result)
			}
		})
	}
}

func TestAvailabilityNumbersWithMinium(t *testing.T) {
	test := []struct {
		name                string
		availabilityNumber  types.AvailabilityNumber
		availabilityNumbers []types.AvailabilityNumber
		want                bool
	}{
		{
			name:               "Returns true when first matches",
			availabilityNumber: constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberPlayable,
			},
			want: true,
		},
		{
			name:               "Returns false even when last matches",
			availabilityNumber: constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberEmpty,
				constants.AvailabilityNumberOccupied,
				constants.AvailabilityNumberValid,
				constants.AvailabilityNumberPlayable,
			},
			want: false,
		},
		{
			name:               "Returns false with no matches",
			availabilityNumber: constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberEmpty,
				constants.AvailabilityNumberOccupied,
				constants.AvailabilityNumberEmpty,
				constants.AvailabilityNumberValid,
			},
			want: false,
		}, {
			name:                "Returns false with not items",
			availabilityNumber:  constants.AvailabilityNumberPlayable,
			availabilityNumbers: []types.AvailabilityNumber{},
			want:                false,
		}, {
			name:               "Returns true when greater than empty",
			availabilityNumber: constants.AvailabilityNumberEmpty,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberValid,
			},
			want: true,
		}, {
			name:               "Returns false when items is less than empty",
			availabilityNumber: constants.AvailabilityNumberEmpty,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberOccupied,
			},
			want: false,
		}, {
			name:               "Returns false when last one lest than empty",
			availabilityNumber: constants.AvailabilityNumberEmpty,
			availabilityNumbers: []types.AvailabilityNumber{
				constants.AvailabilityNumberValid,
				constants.AvailabilityNumberPlayable,
				constants.AvailabilityNumberEmpty,
				constants.AvailabilityNumberOccupied,
			},
			want: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.AvailabilityNumbersWithMinium(tt.availabilityNumber, tt.availabilityNumbers)
			if result != tt.want {
				t.Errorf("The output for test '%s' did not match.\nExpected: %v\nActual: %v\n", tt.name, tt.want, result)
			}
		})
	}
}
