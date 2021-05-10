package main

import (
	"flag"
	"fmt"
	"go-core/03-third-lesson/pkg/crawler"
	"go-core/03-third-lesson/pkg/crawler/spider"
	"go-core/03-third-lesson/pkg/index"
	"log"
)

func main() {
	sFlag := flag.String("s", "", "поиск слов в ссылках после флага -s")
	flag.Parse()
	fmt.Println("Ожидайте, идет сканирование целевых ресурсов.")
	resources := targets{{"http://go.dev", 5}, {"http://golang.org", 2}}
	res := resources.Scan()
	docIndex := index.Make(res)
	if *sFlag != "" {
		fmt.Printf("Искомая строка: %s \n", *sFlag)
		fmt.Printf("Искомая строка содержится на %d страницах: %v \n", len(docIndex[*sFlag]), docIndex[*sFlag])
		fmt.Printf("Результат поиска: \n")
		doc := index.Search(docIndex[*sFlag], index.Documents)
		for _, v := range doc {
			fmt.Printf("%d - номер из списка, адрес: %s , заголовок: %s\n", v.ID, v.URL, v.Title)
		}
	} else {
		fmt.Println("Результат сканирования: \n", res)
	}
}

type target struct {
	url   string
	depth int
}
type targets []target

func (t targets) Scan() []crawler.Document {
	service := spider.New()
	var result []crawler.Document
	for _, v := range t {
		res, err := service.Scan(v.url, v.depth)
		if err != nil {
			log.Fatal(err)
			return result
		}
		result = append(result, res...)
	}
	return result
}
