package adapters_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jerensl/api.jerenslensun.com/internal/adapters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newSqlLiteRepository(t *testing.T, dbPath string) *adapters.SQLiteTokenRepository {
	db, err := adapters.NewSQLiteConnection(dbPath)
	require.NoError(t, err)

	return adapters.NewSQLiteTokenRepository(db)
}

func TestRepository(t *testing.T) {
	dbPath := "../../database/db_test.sqlite"
	r := newSqlLiteRepository(t, dbPath)

	t.Run("Test Update token", func(t *testing.T) {
		testUpdatedToken(t, r)
	})

	t.Run("Test Get token", func(t *testing.T) {
		testGetToken(t, r)
	})

	t.Run("Test Get All token", func(t *testing.T) {
		testGetAllToken(t, r)
	})

	t.Run("Test Delete token", func(t *testing.T) {
		testDeleteToken(t, r)
	})

	t.Run("Test Get All token", func(t *testing.T) {
		testGetAll2Token(t, r)
	})

	err := os.Remove(dbPath)
	if err != nil {
		fmt.Println("cannot remove database")
	}
}

func testUpdatedToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	err := repository.UpdatedToken("abc123")
	require.NoError(t, err)
	err = repository.UpdatedToken("abc321")
	require.NoError(t, err)
}

func testGetToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	hasValue, err := repository.GetToken("abc123")
	require.NoError(t, err)

	assert.True(t, hasValue)
}

func testGetAllToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	expected := []string{"abc123", "abc321"}

	subscriber, err := repository.GetAllToken()
	require.NoError(t, err)

	assert.Equal(t, expected,subscriber)
}

func testGetAll2Token(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	expected := []string{"abc321"}

	subscriber, err := repository.GetAllToken()
	require.NoError(t, err)

	assert.Equal(t, expected,subscriber)
}

func testDeleteToken(t *testing.T, repository *adapters.SQLiteTokenRepository) {
	err := repository.DeleteToken("abc123")
	require.NoError(t, err)
}