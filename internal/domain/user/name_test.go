package user

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	t.Run("valid name cases", func(t *testing.T) {
		validNames := []string{
			"John Doe",
			"Mary-Jane",
			"O'Connor",
			"Jean-Luc Picard",
			"Anne Marie",
			"José María",
			"李小龙",
			"Иван Иванович",
			"محمد علي",
			"Renée O'Connor",
			"Anaïs Nin",
			"Chloë Sevigny",
			"Zoë Kravitz",
			strings.Repeat("a", maxNameLength),
			"abc",
		}

		for _, name := range validNames {
			t.Run(name, func(t *testing.T) {
				n, err := NewName(name)
				assert.NoError(t, err)
				assert.Equal(t, name, n.String())
			})
		}
	})

	t.Run("invalid name cases", func(t *testing.T) {
		invalidNames := []string{
			"",
			"   ",
			"Jo",
			"This is a very long name that exceeds the maximum allowed length for a name in this system",
			"John@Doe!",
			strings.Repeat("a", maxNameLength+1),
		}

		for _, name := range invalidNames {
			t.Run(name, func(t *testing.T) {
				n, err := NewName(name)
				assert.Error(t, err)
				assert.IsType(t, &InvalidNameError{}, err)
				assert.Equal(t, "", n.String())
			})
		}
	})

	t.Run("extreme cases", func(t *testing.T) {
		extremeNames := []string{
			strings.Repeat("a", 100000),
			strings.Repeat("a", 4000000),
			strings.Repeat("a", 10000000),
		}

		for _, name := range extremeNames {
			t.Run(name, func(t *testing.T) {
				n, err := NewName(name)
				assert.Error(t, err)
				assert.IsType(t, &InvalidNameError{}, err)
				assert.Equal(t, "", n.String())
			})
		}
	})
}
