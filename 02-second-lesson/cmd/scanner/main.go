package main

import (
	"flag"
	"fmt"
	"go-core/02-second-lesson/pkg/crawler"
	"go-core/02-second-lesson/pkg/crawler/spider"
	"strings"
)

func main() {
	sFlag := flag.String("s", "", "поиск слов в ссылках после флага -s")
	flag.Parse()
	fmt.Println("Ожидайте, идет сканирование целевых ресурсов.")
	resources := targets{{"http://go.dev", 5}, {"http://golang.org", 2}}
	res, _ := resources.Scan()
	if *sFlag != "" {
		fmt.Printf("Искомая строка: %s \n", *sFlag)
		fmt.Println("Результат поиска:")
		for _, v := range res {
			if strings.Contains(strings.ToLower(v.Title), *sFlag) || strings.Contains(strings.ToLower(v.URL), *sFlag) {
				fmt.Println(v.URL)
			}
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

func (s targets) Scan() ([]crawler.Document, error) {
	service := spider.New()
	var result []crawler.Document
	for _, v := range s {
		res, _ := service.Scan(v.url, v.depth)
		result = append(result, res...)
	}
	return result, nil
}
