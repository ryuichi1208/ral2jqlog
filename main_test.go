package main

import (
	"bytes"
	"testing"
)

func Test_message2CSV(t *testing.T) {
	type args struct {
		csvString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CSV Test1",
			args: args{
				csvString: "20220920 05:38:07,test,common,10.1.1.1,54988141,330935144,QUERY,,'\"set names \\'ujis\\'\",0",
			},
			wantErr: false,
			want:    "{\"timeStamp\":\"20220920 05:38:07\",\"user\":\"common\",\"host\":\"10.1.1.1\",\"command\":\"10.1.1.1\",\"query\":\"\\\"set names \\\\'ujis\\\\'\"}\n",
		},
		{
			name: "CSV Test2 Error",
			args: args{
				csvString: "20220920 05:38:07,test-server,common,100.1.1.1,54988141,330935146,QUERY,,'select * from `test` where `id` is null limit 1',0",
			},
			wantErr: false,
			want:    "{\"timeStamp\":\"20220920 05:38:07\",\"user\":\"common\",\"host\":\"100.1.1.1\",\"command\":\"100.1.1.1\",\"query\":\"select * from `test` where `id` is null limit 1\"}\n",
		},
	}
	var buffer *bytes.Buffer
	buffer = &bytes.Buffer{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer.Reset()
			if err := message2CSV(tt.args.csvString, buffer); (err != nil) != tt.wantErr {
				t.Errorf("message2CSV() error = %v, wantErr %v", err, tt.wantErr)
			}
			if buffer.String() != tt.want {
				t.Errorf("message2CSV() error = %v, want %v", buffer.String(), tt.want)
			}
		})
	}
}

func Test_auditLog2Json(t *testing.T) {
	type args struct {
		jsonString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := auditLog2Json(tt.args.jsonString); (err != nil) != tt.wantErr {
				t.Errorf("auditLog2Json() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readGzip(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test003 nomal",
			args: args{
				filename: "./testdata/test002.gz",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readGzip(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("readGzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDetectFileContentType(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test001 empty files",
			args: args{
				filename: "./testdata/test001.gz",
			},
			want:    "application/x-gzip",
			wantErr: false,
		},
		{
			name: "Test002 no such file",
			args: args{
				filename: "./testdata/test000.gz",
			},
			wantErr: true,
		},
		{
			name: "Test003 nomal",
			args: args{
				filename: "./testdata/test002.gz",
			},
			want:    "application/x-gzip",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectFileContentType(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectFileContentType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DetectFileContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	type args struct {
		opts options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Do(tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
