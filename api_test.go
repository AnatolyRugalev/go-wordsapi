package wordsapi

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var apiKey string

func TestMain(m *testing.M) {
	apiKey = os.Getenv("WORDSAPI_KEY")
	os.Exit(m.Run())
}

func TestWordsAPI_Word(t *testing.T) {
	api := New(apiKey)
	word, err := api.Word(context.Background(), "hello")
	require.NoError(t, err)
	require.Equal(t, "hello", word.Word)
}
