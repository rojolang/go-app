package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElementStoreMount(t *testing.T) {
	t.Run("mounting an element succeeds", func(t *testing.T) {
		var s elementStore
		err := s.Mount(&htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		})
		require.NoError(t, err)
		require.Len(t, s.elements, 1)
	})

	t.Run("mounting an element without id returns an error", func(t *testing.T) {
		var s elementStore
		err := s.Mount(&htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					Depth:     1,
					JSElement: value{},
				},
			},
		})
		require.Error(t, err)
	})

	t.Run("mounting an element without depth returns an error", func(t *testing.T) {
		var s elementStore
		err := s.Mount(&htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					JSElement: value{},
				},
			},
		})
		require.Error(t, err)
	})

	t.Run("mounting an non component element without a js value returns an error", func(t *testing.T) {
		var s elementStore
		err := s.Mount(&htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:    1,
					Depth: 1,
				},
			},
		})
		require.Error(t, err)

	})

	t.Run("mounting an element already mounted returns an error", func(t *testing.T) {
		var s elementStore

		div := &htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		}

		err := s.Mount(div)
		require.NoError(t, err)
		err = s.Mount(div)
		require.Error(t, err)
	})
}

func TestElementStoreMounted(t *testing.T) {
	t.Run("element is mounted", func(t *testing.T) {
		var s elementStore

		div := &htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		}

		err := s.Mount(div)
		require.NoError(t, err)
		require.True(t, s.Mounted(div))
	})

	t.Run("element is not mounted", func(t *testing.T) {
		var s elementStore

		div := &htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		}

		require.False(t, s.Mounted(div))
	})
}

func BenchmarkElementStoreMounted(b *testing.B) {
	var s elementStore

	div := &htmlDiv{
		htmlElement: htmlElement{
			elementDescriptor: elementDescriptor{
				ID:        1,
				Depth:     1,
				JSElement: value{},
			},
		},
	}

	s.Mount(div)

	for n := 0; n < b.N; n++ {
		s.Mounted(div)
	}
}

func TestElementStoreDismount(t *testing.T) {
	t.Run("element is dismounted", func(t *testing.T) {
		var s elementStore

		div := &htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		}

		err := s.Mount(div)
		require.NoError(t, err)

		s.Dismount(div)
		require.Empty(t, s.elements)
	})

}

func BenchmarkElementStoreMountDismount(b *testing.B) {
	var s elementStore

	div := &htmlDiv{
		htmlElement: htmlElement{
			elementDescriptor: elementDescriptor{
				ID:        1,
				Depth:     1,
				JSElement: value{},
			},
		},
	}

	for n := 0; n < b.N; n++ {
		s.Mount(div)
		s.Dismount(div)
	}
}

func TestElementStoreUpdate(t *testing.T) {
	t.Run("updating an element succeeds", func(t *testing.T) {
		var s elementStore

		a := htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		}
		b := a

		err := s.Mount(&a)
		require.NoError(t, err)

		err = s.Update(&b)
		require.NoError(t, err)
	})

	t.Run("updating a non mounted element returns an error", func(t *testing.T) {
		var s elementStore

		a := htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		}

		err := s.Update(&a)
		require.Error(t, err)
	})

	t.Run("updating an element with non matching descriptor returns an error", func(t *testing.T) {
		var s elementStore

		a := htmlDiv{
			htmlElement: htmlElement{
				elementDescriptor: elementDescriptor{
					ID:        1,
					Depth:     1,
					JSElement: value{},
				},
			},
		}
		b := a
		b.elementDescriptor.Depth = 2

		err := s.Mount(&a)
		require.NoError(t, err)

		err = s.Update(&b)
		require.Error(t, err)
	})
}

func BenchmarkElementStoreUpdate(b *testing.B) {
	var s elementStore

	a := htmlDiv{
		htmlElement: htmlElement{
			elementDescriptor: elementDescriptor{
				ID:        1,
				Depth:     1,
				JSElement: value{},
			},
		},
	}

	s.Mount(&a)

	for n := 0; n < b.N; n++ {
		s.Update(&a)
	}
}
