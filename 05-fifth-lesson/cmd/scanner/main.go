package main

import (
	"flag"
	"fmt"
	"go-core/05-fifth-lesson/pkg/crawler"
	"go-core/05-fifth-lesson/pkg/crawler/spider"
	"go-core/05-fifth-lesson/pkg/index"
	"go-core/05-fifth-lesson/pkg/storage"
	"log"
)

func main() {
	sFlag := flag.String("s", "", "поиск слов в ссылках после флага -s")
	flag.Parse()
	fmt.Println("Ожидайте, идет сканирование целевых ресурсов.")
	res := source()
	ind := index.Make(res)
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

func source() []crawler.Document {
	var data []crawler.Document
	resources := map[int]string{5: "http://go.dev", 2: "http://golang.org"}
	if storage.Empty() {
		data = scan(resources)
	} else {
		data = storage.Exrtact()
	}
	return data
}

func scan(m map[int]string) []crawler.Document {
	service := spider.New()
	var result []crawler.Document
	for k, v := range m {
		res, err := service.Scan(v, k)
		if err != nil {
			log.Fatal(err)
			return result
		}
		result = append(result, res...)
	}
	if len(result) > 0 {
		storage.Rec(result)
	}
	return result
}
