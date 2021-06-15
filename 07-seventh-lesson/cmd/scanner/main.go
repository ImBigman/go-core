package main

import (
	"flag"
	"fmt"
	"go-core/07-seventh-lesson/pkg/crawler"
	"go-core/07-seventh-lesson/pkg/crawler/spider"
	"go-core/07-seventh-lesson/pkg/index"
	"go-core/07-seventh-lesson/pkg/storage"
	"log"
)

func main() {
	sFlag := flag.String("s", "", "поиск слов в ссылках после флага -s")
	flag.Parse()
	fmt.Println("Ожидайте, идет сканирование целевых ресурсов.")
	resources := map[int]string{5: "http://go.dev", 2: "http://golang.org"}
	path := "./store/scan_result.txt"
	res, err := dataSource(resources, path)
	if err != nil {
		log.Fatal(err)
	}
	ind := index.Make(res)
	if *sFlag == "" {
		fmt.Println("Вы не указали параметров поиска.")
	}
	if *sFlag != "" {
		fmt.Printf("Искомая строка: %s \n", *sFlag)
		fmt.Printf("Искомая строка содержится на %d страницах: %v \n", len(ind.IndexMap[*sFlag]), ind.IndexMap[*sFlag])
		docs := ind.Search(*sFlag)
		fmt.Printf("Результат поиска: \n")
		for _, v := range docs {
			fmt.Printf("%d - номер из списка, адрес: %s , заголовок: %s\n", v.ID, v.URL, v.Title)
		}
	}
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
