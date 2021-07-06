package strwriter

import (
	"bytes"
	"go-core/09-ninth-lesson/pkg/older"
	"os"
	"testing"
)

func TestWriteStr(t *testing.T) {
	emp := older.Employee{Age: 52}
	str := "hello"
	file, _ := os.Create("./output.txt")
	defer file.Close()
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name:  "First",
			args:  args{[]interface{}{file, emp, str}},
			wantW: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			WriteStr(w, tt.args.args...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteStr() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
