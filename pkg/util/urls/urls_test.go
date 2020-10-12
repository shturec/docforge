package urls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ext(t *testing.T) {
	testCases := []struct {
		link string
		want string
	}{
		{
			"a", "",
		},
		{
			".", "",
		},
		{
			"a.b", "b",
		},
		{
			".b", "b",
		},
		{
			"/a/b.c", "c",
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got := Ext(tc.link)
			assert.Equal(t, tc.want, got)
		})
	}
}
