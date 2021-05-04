package main

import (
	"flag"
	"fmt"
	"go-core/02-second-lesson/pkg/crawler"
	"go-core/02-second-lesson/pkg/crawler/spider"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Ожидайте, идет сканирование целевых ресурсов.")
	a := search{"http://go.dev", 5}
	b := search{"http://golang.org", 2}
	d, _ := a.Scan()
	e, _ := b.Scan()
	res := append(d, e...)
	if len(res) == 0 {
		fmt.Println("Результат сканирования отстуствует.")
		os.Exit(1)
	}

	sFlag := flag.Bool("s", false, "поиск слов в ссылках после флага -s")
	flag.Parse()
	word := flag.Arg(0)

	if *sFlag && word != "" {
		resultSlice := pagesContains(res, word)
		fmt.Printf("Искомая строка: %s, совпадений: %d \n", word, len(resultSlice))
		fmt.Println("Результат поиска:", returnLinks(resultSlice))
	} else {
		fmt.Println("Результат сканирования: \n", res)
	}
}

type doc = crawler.Document
type pages = []doc
type search struct {
	url   string
	depth int
}

// pagesContains возвращает массив страниц, содержащих переданное значение
func pagesContains(slice pages, value string) pages {
	var result pages
	for _, v := range slice {
		if strings.Contains(concateFields(v), value) {
			result = append(result, v)
		}
	}
	return result
}

func returnLinks(slice pages) []string {
	var str []string
	for i, v := range slice {
		str = append(str, "\n"+strconv.Itoa(i+1)+": "+v.URL)
	}
	return str
}

func concateFields(d doc) string {
	return strings.Join([]string{strings.ToLower(d.Title), strings.ToLower(d.URL)}[:], " ")
}

func (s search) Scan() (pages, error) {
	service := spider.New()
	res, err := service.Scan(s.url, s.depth)
	return res, err
}
