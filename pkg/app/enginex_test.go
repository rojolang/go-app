package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEngineX(t *testing.T) {
	e := newEngineX(Body())
	require.NotNil(t, e)
	require.NotNil(t, e.root)
}
