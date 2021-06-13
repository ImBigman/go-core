package index

import (
	"go-core/07-seventh-lesson/pkg/crawler"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestData_Search(t *testing.T) {
	docs := []Document{{1, "Search", "Go"}, {2, "Dev", "Go"}, {3, "Url", "Go"}}
	index := map[string][]int{"search": {1}, "url": {3}, "go": {1, 2, 3}}
	type fields struct {
		IndexDocs []Document
		IndexMap  map[string][]int
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Document
	}{
		{
			name:   "Первый тест",
			fields: fields{IndexDocs: docs, IndexMap: index},
			args:   args{"search"},
			want:   []Document{{1, "Search", "Go"}},
		},
		{
			name:   "Второй тест",
			fields: fields{IndexDocs: docs, IndexMap: index},
			args:   args{"url"},
			want:   []Document{{3, "Url", "Go"}},
		},
		{
			name:   "Третий тест",
			fields: fields{IndexDocs: docs, IndexMap: index},
			args:   args{"go"},
			want:   []Document{{1, "Search", "Go"}, {2, "Dev", "Go"}, {3, "Url", "Go"}},
		},
		{
			name:   "Четвертый тест",
			fields: fields{IndexDocs: docs, IndexMap: index},
			args:   args{"help"},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				IndexDocs: tt.fields.IndexDocs,
				IndexMap:  tt.fields.IndexMap,
			}
			if got := d.Search(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniqInt(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Первый тест",
			args: args{[]int{1, 2, 3, 4, 4, 6, 1, 3}},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "Второй тест",
			args: args{[]int{1, 1, 1, 4, 4, 4, 1, 2}},
			want: []int{1, 4, 2},
		},
		{
			name: "Третий тест",
			args: args{[]int{1, 1, 1, 1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqInt(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binSearch(t *testing.T) {
	docs := []Document{{1, "Url", "Title"}, {2, "Url", "Title"}, {3, "Url", "Title"}, {4, "Url", "Title"}, {5, "Url", "Title"}}
	type args struct {
		d []Document
		n int
	}
	tests := []struct {
		name string
		args args
		want Document
	}{
		{
			name: "Первый тест",
			args: args{d: docs, n: 3},
			want: docs[2],
		},
		{
			name: "Второй тест",
			args: args{d: docs, n: 0},
			want: Document{},
		},
		{
			name: "Третий тест",
			args: args{d: docs, n: 4},
			want: docs[3],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binSearch(tt.args.d, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_binSearch(b *testing.B) {
	data := sampleData(1_000_000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		res := binSearch(data, n)
		_ = res
	}
}

func sampleData(len int) []Document {
	rand.Seed(time.Now().UnixNano())
	var data []Document
	for i := 0; i < len; i++ {
		data = append(data, Document{ID: rand.Intn(1000), URL: "url ", Title: "title"})
	}
	sort.Slice(data, func(i, j int) bool { return data[i].ID < data[j].ID })
	return data
}

func TestMake(t *testing.T) {
	docs := []crawler.Document{
		{ID: 3, URL: "go.dev", Title: "docs", Body: ""},
		{ID: 1, URL: "dev.com", Title: "go", Body: ""},
		{ID: 4, URL: "go.dev", Title: "about", Body: ""},
		{ID: 2, URL: "go.dev", Title: "dev", Body: ""},
		{ID: 5, URL: "go.dev", Title: "Title", Body: ""},
	}
	type args struct {
		d []crawler.Document
	}
	tests := []struct {
		name string
		args args
		want Data
	}{
		{
			name: "Первый тест",
			args: args{docs},
			want: Data{
				[]Document{{0, "go.dev", "docs"}, {1, "dev.com", "go"}, {2, "go.dev", "about"}, {3, "go.dev", "dev"}, {4, "go.dev", "Title"}},
				map[string][]int{"about": {2}, "com": {1}, "dev": {0, 1, 2, 3, 4}, "docs": {0}, "go": {0, 1, 2, 3, 4}, "title": {4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Make(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Make() = %v, want %v", got, tt.want)
			}
		})
	}
}
