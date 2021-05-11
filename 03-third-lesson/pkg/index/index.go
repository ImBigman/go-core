package index

import (
	"errors"
	"go-core/03-third-lesson/pkg/crawler"
	"log"
	"regexp"
	"strings"
)

// Используется для создания индекса типа -  "строка": [номера страниц, на которых она находится]
func Make(d []crawler.Document) Index {
	docIndex := make(Index)
	var docs []Document
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range d {
		// применяется для формирования совокупности слов из URL и Title - как ключи индекса
		s := reg.ReplaceAllString(strings.ToLower(v.Title+" "+v.URL), " ")
		docStrings := strings.Fields(s)
		// используется для формирования карты индекса из слов документа и его ID
		for _, v := range docStrings {
			docIndex[v] = append(docIndex[v], i)
		}
		// применяется  для формирования слайса из документов нового типа данных
		docs = append(docs, Document{i, v.URL, v.Title})
	}
	// используется для формирования слайса уникальных ID страниц - как значение индекса
	for k, v := range docIndex {
		docIndex[k] = uniqInt(v)
	}
	Documents = docs
	return docIndex
}

func Search(n []int, d []Document) ([]Document, error) {
	var result []Document
	if len(n) != 0 {
		for _, v := range n {
			res, err := binarySearch(d, v)
			if err != nil {
				return result, errors.New("Нет совпадений")
			}
			result = append(result, res)
		}
	} else {
		return result, errors.New("Совпадений не найдено")
	}
	return result, nil
}

var Documents []Document

type Index map[string][]int
type Document struct {
	ID    int
	URL   string
	Title string
}

// для создания слайсов с уникальными значениями для чисел
func uniqInt(slice []int) []int {
	keys := make(map[int]bool)
	res := []int{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			res = append(res, entry)
		}
	}
	return res
}

func binarySearch(d []Document, n int) (Document, error) {
	hf := len(d) / 2
	var result Document
	switch {
	case len(d) == 0:
		return Document{}, errors.New("пустой список")
	case d[hf].ID > n:
		result, _ = binarySearch(d[:hf], n)
	case d[hf].ID < n:
		result, _ = binarySearch(d[hf+1:], n)
	default:
		result = d[hf]
	}
	return result, nil
}
