package storage

import (
	"encoding/json"
	"go-core/05-fifth-lesson/pkg/crawler"
	"io"
	"io/ioutil"
	"os"
)

const path = "./store/scan_result.txt"

// Используется для записи данных в файл
func Rec(docs []crawler.Document) (bool, error) {
	f, err := os.Create(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	data, err := json.Marshal(docs)
	if err != nil {
		return false, err
	}
	err = ioutil.WriteFile(f.Name(), data, 0666)
	if err != nil {
		return false, err
	}
	return true, err
}

// для извлечение данных из имеющегося файла
func Exrtact() ([]crawler.Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	var data []byte
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if n > 0 {
			data = append(data, buf[:n]...)
		}
	}
	var res []crawler.Document
	json.Unmarshal([]byte(data), &res)
	return res, nil
}

// для проверки существования файла с данными
func Empty() bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}
