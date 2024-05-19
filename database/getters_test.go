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
	type fields struct {
		DB *Database
	}
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Get{
				DB: tt.fields.DB,
			}
			got, err := g.TableColumns(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("TableColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TableColumns() got = %v, want %v", got, tt.want)
			}
		})
	}
}
