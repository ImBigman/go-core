package storage

import (
	"encoding/json"
	"go-core/05-fifth-lesson/pkg/crawler"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const path = "./store/scan_result.txt"

func Rec(docs []crawler.Document) (bool, error) {
	f, err := os.Create(path)
	catch(err)
	defer f.Close()

	data, err := json.Marshal(docs)
	catch(err)
	err = ioutil.WriteFile(f.Name(), data, 0666)
	catch(err)
	return true, err
}

func Exrtact() []crawler.Document {
	f, err := os.Open(path)
	catch(err)
	defer f.Close()

	buf := make([]byte, 1024)
	var data []byte
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			continue
		}
		if n > 0 {
			data = append(data, buf[:n]...)
		}
	}
	var res []crawler.Document
	json.Unmarshal([]byte(data), &res)
	return res
}

func Empty() bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func catch(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
