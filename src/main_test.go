package main

import "testing"

func TestSoma(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Teste 1",
			args: args{10, 10},
			want: 20,
		}, {
			name: "Teste 2",
			args: args{5, 10},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Soma(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Soma() = %v, want %v", got, tt.want)
			}
		})
	}
}
