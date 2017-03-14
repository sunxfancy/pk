package pkget

import (
	"reflect"
	"testing"
)

func TestNewPkget(t *testing.T) {
	tests := []struct {
		name string
		want *Pkget
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPkget(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPkget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPkget_loadConfig(t *testing.T) {
	tests := []struct {
		name string
		p    Pkget
	}{
		// TODO: Add test cases.
		{"default_load", Pkget{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.loadConfig()
			if err != nil {
				t.Skip()
			}
			if tt.p.InstallPath != "/Pkget/" {
				t.Error()
			}
		})
	}
}

func TestPkget_getPkg(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		p    Pkget
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.getPkg(tt.args.url)
		})
	}
}

func TestPkget_checkVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		p    Pkget
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.checkVersion(tt.args.version)
		})
	}
}

func TestPkget_downloadFile(t *testing.T) {
	type args struct {
		filepath string
		url      string
	}
	tests := []struct {
		name    string
		p       Pkget
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"download_demo", Pkget{}, args{"/tmp/downloadtest.html", "http://www.baidu.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.downloadFile(tt.args.filepath, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Pkget.downloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
