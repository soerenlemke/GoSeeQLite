package database

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestGet_AllTableNames(t *testing.T) {
	testSampleValidDb, _ := filepath.Abs(filepath.Join("testdata", "chinook.db"))
	type fields struct {
		dsn string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			"Valid database", fields{dsn: testSampleValidDb}, []string{"albums", "sqlite_sequence", "artists", "customers", "employees", "genres", "invoices", "invoice_items", "media_types", "playlists", "playlist_track", "tracks", "sqlite_stat1"}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _ := NewDatabase(tt.fields.dsn)
			got, err := db.Get.AllTableNames()
			if (err != nil) != tt.wantErr {
				t.Errorf("AllTableNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("AllTableNames() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet_TableColumns(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantFk  []bool // Since each column also has a value that shows if it is Fk or not
		wantErr bool
	}{
		{
			"Valid table with no foreign key", "artists", []bool{false, false}, false,
		},
		{
			"Valid table with foreign key", "albums", []bool{false, false, true}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSampleValidDb, _ := filepath.Abs(filepath.Join("testdata", "chinook.db"))
			db, _ := NewDatabase(testSampleValidDb)
			got, err := db.Get.TableColumns(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("TableColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, col := range got {
				if col.ForeignKey != tt.wantFk[i] {
					t.Errorf("TableColumns() error : Table : %s Column : %s WantFk : %t GotFk : %t", tt.args, col.Name, tt.wantFk[i], col.ForeignKey)
					return
				}
			}
		})
	}
}
