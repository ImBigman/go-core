package main

import (
	"go-core/08-eighth-lesson/pkg/crawler"
	"go-core/08-eighth-lesson/pkg/crawler/spider"
	"go-core/08-eighth-lesson/pkg/index"
	"go-core/08-eighth-lesson/pkg/storage"
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	trace.Start(f)
	defer trace.Stop()

	resources := map[int]string{5: "http://go.dev", 2: "http://golang.org"}
	path := "./store/scan_result.txt"
	res, err := dataSource(resources, path)
	if err != nil {
		log.Fatal(err)
	}
	ind := index.Make(res)
	word := "dev"
	ind.Search(word)
}

func dataSource(m map[int]string, path string) ([]crawler.Document, error) {
	var data []crawler.Document
	var err error
	if !storage.CheckStorageExist(path) {
		return scan(m, path)
	}
	if storage.CheckStorageExist(path) {
		return storage.Extract(path)
	}
	return data, err
}

func scan(m map[int]string, path string) ([]crawler.Document, error) {
	service := spider.New()
	var result []crawler.Document
	for k, v := range m {
		res, err := service.Scan(v, k)
		if err != nil {
			return nil, err
		}
		result = append(result, res...)
	}
	if len(result) > 0 {
		_, err := storage.Store(result, path)
		return result, err
	}
	return result, nil
}
