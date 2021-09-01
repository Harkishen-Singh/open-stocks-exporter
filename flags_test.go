package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcessSymbols(t *testing.T) {
	tcs := []struct{
		name string
		symbolsStr string
		expectedSymbols []string
		shouldFail bool
	}{
		{
			name: "empty",
			symbolsStr: "",
			shouldFail: true,
		},
		{
			name: "empty with space",
			symbolsStr: " ",
			shouldFail: true,
		},
		{
			name: "valid one",
			symbolsStr: "ABCD",
			expectedSymbols: []string{"ABCD"},
		},
		{
			name: "valid two",
			symbolsStr: "ABCD,EFGH",
			expectedSymbols: []string{"ABCD", "EFGH"},
		},
	}

	for _, tc := range tcs {
		cfg := &config{symbolsStr: tc.symbolsStr}
		if tc.shouldFail {
			require.Error(t, processSymbols(cfg))
		} else {
			require.NoError(t, processSymbols(cfg))
		}
		require.Equal(t, tc.expectedSymbols, cfg.symbols)
	}
}