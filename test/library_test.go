package test

import (
	"context"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"library/api/model"
	"library/client"
	"testing"
)

func TestService(t *testing.T) {
	ctx := context.Background()

	tester := RunTester(t, model.ServerOptions{DBOptions: model.DBOptions{
		User:     "test_service",
		Password: "test_service",
		Connect:  "127.0.0.1:3306",
		DBName:   "test_service",
	},
		Port: flag.Int("port", 50051, "The server port")})
	defer tester.Close()

	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	cl, err := client.NewLibraryClient(addr)
	require.NoError(t, err)
	defer cl.Close()

	resp, err := cl.GetBooksByAuthor(ctx, "Author1")
	require.NoError(t, err)
	require.Equal(t, len(resp.Books), 2)
	assert.Equal(t, resp.Books[0].Title, "Book1")
	assert.Equal(t, resp.Books[1].Title, "Book2")

	resp, err = cl.GetBooksByTitle(ctx, "Book3")
	require.NoError(t, err)
	require.Equal(t, len(resp.Books), 1)
	assert.Equal(t, resp.Books[0].Author, "Author2")
}
