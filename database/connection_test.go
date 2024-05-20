package database

import (
	"database/sql"
	"path/filepath"
	"testing"
)

func TestDatabase_Connect(t *testing.T) {
	type fields struct {
		dsn        string
		connection *sql.DB
		Get        *Get
		Set        *Set
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test case 1: Valid Connection",
			fields: fields{
				dsn:        "test.db",
				connection: &sql.DB{}, // assuming this is a valid connection
				Get:        &Get{},
				Set:        &Set{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				dsn:        tt.fields.dsn,
				connection: tt.fields.connection,
				Get:        tt.fields.Get,
				Set:        tt.fields.Set,
			}
			if err := db.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_ConnectionStatus(t *testing.T) {
	// This is pretty much a useless test as NewDatabase()
	// already calls ConnectionStatus.
	// TODO: Find a better test
	testSampleValidDb, _ := filepath.Abs(filepath.Join("testdata", "chinook.db"))
	tests := []struct {
		name    string
		wantErr bool
		wantOk  bool
		dsn     string
	}{
		{
			"Valid Database", false, true, testSampleValidDb,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := NewDatabase(tt.dsn)
			gotOk, gotErr := db.ConnectionStatus()
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("ConnectionStaus() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}

			if gotOk != tt.wantOk {
				t.Errorf("ConnectionStatus() ok = %v, wantOk %v", gotOk, tt.wantOk)
				return
			}
		})
	}
}
