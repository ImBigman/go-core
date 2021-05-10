package index

import (
	"errors"
	"go-core/03-third-lesson/pkg/crawler"
	"log"
	"regexp"
	"strings"
)

func Make(d []crawler.Document) map[string][]int {
	stringMap := make(map[string][]int)
	var stringSlice []string
	var docs []Document
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range d {
		// у всех страниц удаляю символы и возвращаю только слова из URL и Title для создания индекса
		s := reg.ReplaceAllString(strings.ToLower(v.Title+" "+v.URL), " ")
		docStrings := strings.Fields(s)
		stringSlice = append(stringSlice, docStrings...)
		stringSlice = uniqString(stringSlice)
		// создаю карту индекса
		for _, v := range docStrings {
			stringMap[v] = append(stringMap[v], i)
		}
		// формирую слайс из документов нового типа данных
		docs = append(docs, Document{i, v.URL, v.Title})
	}
	for k, v := range stringMap {
		stringMap[k] = uniqInt(v)
	}
	Documents = docs
	return stringMap
}
func Search(n []int, d []Document) []Document {
	var result []Document
	for _, v := range n {
		res, err := binarySearch(d, v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, res)
	}
	return result
}

var Documents []Document

type Document struct {
	ID    int
	URL   string
	Title string
}

// два метода для создания слайсов с уникальными значениями для строк и чисел
func uniqString(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func uniqInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func binarySearch(d []Document, search int) (Document, error) {
	hf := len(d) / 2
	var result Document
	switch {
	case len(d) == 0:
		return Document{}, errors.New("пустой список")
	case d[hf].ID > search:
		result, _ = binarySearch(d[:hf], search)
	case d[hf].ID < search:
		result, _ = binarySearch(d[hf+1:], search)
	default:
		result = d[hf]
	}
	return result, nil
}
