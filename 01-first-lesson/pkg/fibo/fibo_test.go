// Package fibo возвращает значения по порядковому номеру.
package fibo

import "testing"

func TestNum(t *testing.T) {
	type args struct {
		f int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{f: 8.},
			want: 13,
		},
		{
			name: "2",
			args: args{f: 19.},
			want: 2584,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Num(tt.args.f); got != tt.want {
				t.Errorf("Num() = %v, want %v", got, tt.want)
			}
		})
	}
}
