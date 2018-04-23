package services

import (
	"reflect"
	"testing"
)

func TestGetLabInfo(t *testing.T) {
	type args struct {
		lab string
	}
	tests := []struct {
		name    string
		args    args
		want    LabInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLabInfo(tt.args.lab)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLabInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLabInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
