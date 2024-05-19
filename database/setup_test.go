package database

import (
	"path/filepath"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	testSampleDb, _ := filepath.Abs(filepath.Join("testdata", "test_db.db"))
	tests := []struct {
		name    string
		args    string //pathToDb
		wantErr bool
	}{
		// TODO: Test case for an invalid database file
		{
			"Empty Db path", "", true,
		},
		{
			"Db path does not exist", "does/not/exist.db", true,
		},
		{
			"Valid db", testSampleDb, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDatabase(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Doing this to check whether the output we got
				// is an initialized Database instance
				if got.dsn != tt.args || got.Get == nil || got.Set == nil || got.connection == nil {
					t.Fatalf("NewDatabase() returned uninitialized Database instance")
					return
				}
			}
		})
	}
}
