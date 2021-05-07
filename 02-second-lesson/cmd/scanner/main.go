package main

import (
	"flag"
	"fmt"
	"go-core/02-second-lesson/pkg/crawler"
	"go-core/02-second-lesson/pkg/crawler/spider"
	"log"
	"strings"
)

func main() {
	sFlag := flag.String("s", "", "поиск слов в ссылках после флага -s")
	flag.Parse()
	fmt.Println("Ожидайте, идет сканирование целевых ресурсов.")
	resources := targets{{"http://go.dev", 5}, {"http://golang.org", 2}}
	res := resources.Scan()
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
