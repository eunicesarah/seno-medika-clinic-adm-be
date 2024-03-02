package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	require.Equal(t, 3, Add(1, 2))
}
