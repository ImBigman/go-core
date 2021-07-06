package older

import (
	"reflect"
	"testing"
)

func TestEldest(t *testing.T) {
	emp, emp_1 := Employee{Age: 52}, Employee{Age: 13}
	customer, customer_1 := Customer{Age: 30}, Customer{Age: 53}
	type args struct {
		arg []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "First",
			args: args{[]interface{}{emp, emp_1, customer, customer_1}},
			want: customer_1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eldest(tt.args.arg...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eldest() = %v, want %v", got, tt.want)
			}
		})
	}
}
