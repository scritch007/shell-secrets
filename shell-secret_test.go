package shellsecret

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_shellSecret(t *testing.T) {
	filePath := path.Join(t.TempDir(), "test.secureshell")

	s := &shellSecret{
		secureFilePath: filePath,
		key:            []byte("112233445566778899001122"),
	}

	type test struct {
		Test string
	}
	a := test{}

	assert.Error(t, s.Get("a", &a))
	input := test{
		Test: "value",
	}
	assert.NoError(t, s.Add("a", &input))
	assert.NoError(t, s.Get("a", &a))
	assert.Equal(t, input, a)
	assert.NoError(t, s.Add("b", &input))

	list, err := s.List()
	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, list)
	assert.NoError(t, s.Delete("a"))
	list, err = s.List()
	require.NoError(t, err)
	assert.Equal(t, []string{"b"}, list)

}
