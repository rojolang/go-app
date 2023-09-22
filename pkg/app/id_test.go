package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDGenerator(t *testing.T) {
	t.Run("a new non zero id is generated", func(t *testing.T) {
		var gen idGenerator

		id := gen.NewID()
		require.NotZero(t, id)
		require.Equal(t, ID(1), id)
	})

	t.Run("next created id is an incremented value", func(t *testing.T) {
		var gen idGenerator

		id1 := gen.NewID()
		id2 := gen.NewID()
		require.Equal(t, ID(1), id1)
		require.Equal(t, ID(2), id2)
	})

	t.Run("next created id is a reused value", func(t *testing.T) {
		var gen idGenerator

		id1 := gen.NewID()
		id2 := gen.NewID()
		id3 := gen.NewID()
		gen.ReuseID(id2)
		id4 := gen.NewID()
		require.Equal(t, ID(1), id1)
		require.Equal(t, ID(2), id2)
		require.Equal(t, ID(3), id3)
		require.Equal(t, ID(2), id4)
	})

	t.Run("zero id is not reused", func(t *testing.T) {
		var gen idGenerator

		gen.ReuseID(0)
		require.Empty(t, gen.reusableIDs)
	})
}

func BenchmarkIDGeneratorNewID(b *testing.B) {
	var gen idGenerator

	for n := 0; n < b.N; n++ {
		gen.NewID()
	}
}

func BenchmarkIDGeneratorReuseID(b *testing.B) {
	var gen idGenerator

	id := gen.NewID()

	for n := 0; n < b.N; n++ {
		gen.ReuseID(id)
		gen.NewID()
	}
}
