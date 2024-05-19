package database

import (
	"path/filepath"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	testSampleInvalidDb, _ := filepath.Abs(filepath.Join("testdata", "invaliddb.go"))
	testSampleValidDb, _ := filepath.Abs(filepath.Join("testdata", "chinook.db"))
	tests := []struct {
		name    string
		args    string //pathToDb
		wantErr bool
	}{
		{
			"Empty Db path", "", true,
		},
		{
			"Db path does not exist", "does/not/exist.db", true,
		},
		{
			"Invalid Db file", testSampleInvalidDb, true,
		},
		{
			"Valid db", testSampleValidDb, false,
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
