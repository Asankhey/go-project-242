package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"code"
)

func TestGetPathSize_File(t *testing.T) {
	size, err := code.GetSize("../testdata/512b.txt", true, false)
	require.NoError(t, err)
	require.Equal(t, int64(512), size)
}
