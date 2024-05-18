package database

import (
	"reflect"
	"testing"
)

func TestGet_AllTableNames(t *testing.T) {
	type fields struct {
		DB *Database
	}
	tests := []struct {
		name    string
		fields  fields
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
			got, err := g.AllTableNames()
			if (err != nil) != tt.wantErr {
				t.Errorf("AllTableNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllTableNames() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet_DatabaseName(t *testing.T) {
	type fields struct {
		DB *Database
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Get{
				DB: tt.fields.DB,
			}
			got, err := g.DatabaseName()
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DatabaseName() got = %v, want %v", got, tt.want)
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
