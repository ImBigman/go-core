package main

import (
	"go-core/07-seventh-lesson/pkg/crawler"
	"reflect"
	"testing"
)

func Test_dataSource(t *testing.T) {
	res := map[int]string{5: "http://go.dev", 2: "http://golang.org"}
	path := "./store/test_result.txt"
	docs := []crawler.Document{
		{ID: 3, URL: "go.dev", Title: "docs", Body: ""},
		{ID: 1, URL: "dev.com", Title: "go", Body: ""},
		{ID: 4, URL: "go.dev", Title: "about", Body: ""},
		{ID: 2, URL: "go.dev", Title: "dev", Body: ""},
		{ID: 5, URL: "go.dev", Title: "Title", Body: ""},
	}
	type args struct {
		m    map[int]string
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []crawler.Document
		wantErr bool
	}{
		{
			name:    "Первый тест",
			args:    args{res, path},
			want:    docs,
			wantErr: false,
		},
		{
			name:    "Второй тест",
			args:    args{nil, ""},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataSource(tt.args.m, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("dataSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scan(t *testing.T) {
	type args struct {
		m    map[int]string
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []crawler.Document
		wantErr bool
	}{
		{
			name:    "Первый тест",
			args:    args{nil, "./store/empty.txt"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scan(tt.args.m, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
