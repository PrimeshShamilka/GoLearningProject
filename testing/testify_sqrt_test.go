package testing

// TESTFIRY provides methods like assert, require, and mock for easier testing

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestSqrtWithTestify(t *testing.T) {
	val, err := Sqrt(2)
	require.NoError(t, err)
	require.InDelta(t, 1.4142, val, 0.0001)
}
