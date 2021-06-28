package index

import (
	"go-core/08-eighth-lesson/pkg/crawler"
	"log"
	"regexp"
	"strings"
)

type Document struct {
	ID    int
	URL   string
	Title string
}

type Data struct {
	IndexDocs []Document
	IndexMap  map[string][]int
}

// Используется для создания индекса типа -  "строка": [номера страниц, на которых она находится]
func Make(d []crawler.Document) Data {
	docIndex := make(map[string][]int)
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
	index := Data{docs, docIndex}
	return index
}

func (d Data) Search(s string) []Document {
	var result []Document
	if len(d.IndexDocs) != 0 {
		ran := d.IndexMap[s]
		for _, v := range ran {
			res := binSearch(d.IndexDocs, v)
			result = append(result, res)
		}
	}
	return result
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

func binSearch(d []Document, n int) Document {
	var result Document
	start := 0
	hf := len(d) / 2
	end := len(d) - 1
	for start <= end {
		res := d[hf]
		if res.ID == n {
			return d[hf]
		}
		if res.ID > n {
			end = hf - 1
			hf = (start + end) / 2
			continue
		}
		start = hf + 1
		hf = (start + end) / 2
	}
	return result
}
