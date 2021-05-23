package main

import (
	"flag"
	"fmt"
	"go-core/05-fifth-lesson/pkg/crawler"
	"go-core/05-fifth-lesson/pkg/crawler/spider"
	"go-core/05-fifth-lesson/pkg/index"
	"go-core/05-fifth-lesson/pkg/storage"
)

func main() {
	sFlag := flag.String("s", "", "поиск слов в ссылках после флага -s")
	flag.Parse()
	fmt.Println("Ожидайте, идет сканирование целевых ресурсов.")
	res, err := dataSource()
	if err != nil {
		fmt.Println("Ошибка при работе с файлом данных.")
	}
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

func dataSource() ([]crawler.Document, error) {
	var data []crawler.Document
	var err error
	resources := map[int]string{5: "http://go.dev", 2: "http://golang.org"}
	if storage.Empty() {
		data, err = scan(resources)
	} else {
		data, err = storage.Exrtact()
	}
	return data, err
}

func scan(m map[int]string) ([]crawler.Document, error) {
	service := spider.New()
	var result []crawler.Document
	for k, v := range m {
		res, err := service.Scan(v, k)
		if err != nil {
			return result, err
		}
		result = append(result, res...)
	}
	if len(result) > 0 {
		_, err := storage.Rec(result)
		return result, err
	}
	return result, nil
}
