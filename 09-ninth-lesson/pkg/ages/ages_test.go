package ages

import "testing"

func TestOldest(t *testing.T) {
	emp, emp_1 := NewEmployee(52), NewEmployee(13)
	customer, customer_1 := NewCustomer(10), NewCustomer(66)
	type args struct {
		arg []Ager
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First",
			args: args{[]Ager{emp, emp_1, customer, customer_1}},
			want: 66,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Oldest(tt.args.arg...); got != tt.want {
				t.Errorf("Oldest() = %v, want %v", got, tt.want)
			}
		})
	}
}
