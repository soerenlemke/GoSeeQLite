package database

import (
	"reflect"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	type args struct {
		pathToDb string
	}
	tests := []struct {
		name    string
		args    args
		want    Database
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDatabase(tt.args.pathToDb)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabase() got = %v, want %v", got, tt.want)
			}
		})
	}
}
