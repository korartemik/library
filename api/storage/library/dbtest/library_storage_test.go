package dbtest

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"library/api/model"
	"library/api/storage/library"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	cmd := exec.Command("docker", "build", "-f", "docker/dockerfile", "-t", "mysql", ".")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	require.NoError(t, err)

	finishAndDelete()

	cmd = exec.Command("docker", "run", "--name", "mysql_container", "-d", "-p", "3306:3306", "mysql:latest")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	require.NoError(t, err)

	time.Sleep(30 * time.Second)
	t.Run("get by author", func(t *testing.T) {
		db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/test")
		defer db.Close()
		require.NoError(t, err)
		st := library.NewLibraryStorage()

		books, err := st.GetByAuthor(context.Background(), "Author1", db)

		require.NoError(t, err)
		require.Equal(t, len(books), 2)
		assert.Equal(t, books[0].Title, "Book1")
		assert.Equal(t, books[1].Title, "Book2")
	})

	t.Run("get by title", func(t *testing.T) {
		db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/test")
		defer db.Close()
		require.NoError(t, err)
		st := library.NewLibraryStorage()

		books, err := st.GetByTitle(context.Background(), "Book3", db)

		require.NoError(t, err)
		require.Equal(t, len(books), 1)
		assert.Equal(t, books[0].Author, "Author2")
	})

	t.Run("not found", func(t *testing.T) {
		db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/test")
		defer db.Close()
		require.NoError(t, err)
		st := library.NewLibraryStorage()

		books, err := st.GetByTitle(context.Background(), "Book6", db)

		require.Error(t, err)
		require.Nil(t, books)
		assert.True(t, model.IsNotFound(err))
	})

	finishAndDelete()
}

func finishAndDelete() {
	cmd := exec.Command("docker", "stop", "mysql_container")
	cmd.Run()

	cmd = exec.Command("docker", "rm", "mysql_container")
	cmd.Run()
}
