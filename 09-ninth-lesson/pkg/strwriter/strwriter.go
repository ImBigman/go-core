package strwriter

import (
	"io"
	"log"
)

func WriteStr(w io.Writer, args ...interface{}) {
	for _, v := range args {
		if str, ok := v.(string); ok {
			buf := []byte(str)
			_, err := w.Write(buf)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
