package storage

import (
	"fmt"
	"go-core/07-seventh-lesson/pkg/crawler"
	"path/filepath"
	"reflect"
	"testing"
)

func TestStore(t *testing.T) {
	docs := []crawler.Document{
		{ID: 3, URL: "go.dev", Title: "docs", Body: ""},
		{ID: 1, URL: "dev.com", Title: "go", Body: ""},
		{ID: 4, URL: "go.dev", Title: "about", Body: ""},
		{ID: 2, URL: "go.dev", Title: "dev", Body: ""},
		{ID: 5, URL: "go.dev", Title: "Title", Body: ""},
	}

	dir, _ := filepath.Abs("../../cmd/scanner/store/test_result.txt")
	type args struct {
		docs []crawler.Document
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Первый тест",
			args:    args{docs, dir},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Второй тест",
			args:    args{docs, ""},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.path)
			got, err := Store(tt.args.docs, tt.args.path)
			fmt.Println(dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtract(t *testing.T) {
	dir, _ := filepath.Abs("../../cmd/scanner/store/test_result.txt")
	type args struct {
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
			args:    args{""},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Второй тест",
			args: args{dir},
			want: []crawler.Document{
				{ID: 3, URL: "go.dev", Title: "docs", Body: ""},
				{ID: 1, URL: "dev.com", Title: "go", Body: ""},
				{ID: 4, URL: "go.dev", Title: "about", Body: ""},
				{ID: 2, URL: "go.dev", Title: "dev", Body: ""},
				{ID: 5, URL: "go.dev", Title: "Title", Body: ""},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Extract(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckStorageExist(t *testing.T) {
	dir, _ := filepath.Abs("../../cmd/scanner/store/test_result.txt")
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Первый тест",
			args: args{dir},
			want: true,
		},
		{
			name: "Второй тест",
			args: args{""},
			want: false,
		},
		{
			name: "Третий тест",
			args: args{"../../cmd/scanner/store/empty.txt"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckStorageExist(tt.args.path); got != tt.want {
				t.Errorf("CheckStorageExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
