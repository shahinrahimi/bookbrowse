package stores

import (
	"log"
	"os"
)

func setupTestLogger() *log.Logger {
	return log.New(os.Stdout, "[BOOKBROWSE-TEST] ", log.LstdFlags)
}

func SetupTestStore() *SqliteStore {
	logger := setupTestLogger()
	return NewTestSqliteStore(logger)
}
