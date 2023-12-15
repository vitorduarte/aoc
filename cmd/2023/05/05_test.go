package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapIdRange(t *testing.T) {
	sheet := Sheet{
		"50-97":   "52-99",
		"98-99":   "50-51",
		"120-130": "100-110",
	}

	testCases := []struct {
		name             string
		idRange          string
		expectedResponse []string
	}{
		{
			name:             "id inside of one key",
			idRange:          "53-80",
			expectedResponse: []string{"55-82"},
		},
		{
			name:             "id in both keys",
			idRange:          "95-99",
			expectedResponse: []string{"97-99", "50-51"},
		},
		{
			name:             "id is around one key",
			idRange:          "115-135",
			expectedResponse: []string{"100-110", "115-119", "131-135"},
		},
		{
			name:             "id upper is inside one key",
			idRange:          "115-129",
			expectedResponse: []string{"100-109", "115-119"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			valueMapped := sheet.mapIdRange(tc.idRange)
			assert.Equal(tt, tc.expectedResponse, valueMapped)
		})
	}

}
