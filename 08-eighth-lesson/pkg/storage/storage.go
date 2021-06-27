package storage

import (
	"encoding/json"
	"go-core/08-eighth-lesson/pkg/crawler"
	"io"
	"os"
)

// Используется для записи данных в файл
func Store(docs []crawler.Document, path string) (bool, error) {
	f, err := os.Create(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	data, err := json.Marshal(docs)
	if err != nil {
		return false, err
	}
	err = os.WriteFile(f.Name(), data, 0666)
	if err != nil {
		return false, err
	}
	return true, err
}

// для извлечение данных из имеющегося файла
func Extract(path string) ([]crawler.Document, error) {
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
func CheckStorageExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	fileInfo, err := os.Lstat(path)
	if fileInfo.Size() == 0 {
		return false
	}
	if err != nil {
		return false
	}
	return true
}
