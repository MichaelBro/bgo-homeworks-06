package card

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		transactions []Transaction
	}
	tests := []struct {
		name string
		args args
		want []Transaction
	}{
		{
			name: "#1",
			args: args{
				transactions: []Transaction{
					{
						Id:     1,
						Amount: 100,
					},
					{
						Id:     2,
						Amount: 3000,
					},
					{
						Id:     3,
						Amount: 200,
					},
					{
						Id:     4,
						Amount: 1000,
					},
				},
			},
			want: []Transaction{
				{
					Id:     2,
					Amount: 3000,
				},
				{
					Id:     4,
					Amount: 1000,
				},
				{
					Id:     3,
					Amount: 200,
				},
				{
					Id:     1,
					Amount: 100,
				},
			},
		},
		{
			name: "#2",
			args: args{
				transactions: []Transaction{
					{
						Id:     5,
						Amount: 1000,
					},
					{
						Id:     2,
						Amount: 3000,
					},
					{
						Id:     4,
						Amount: 1000,
					},
				},
			},
			want: []Transaction{
				{
					Id:     2,
					Amount: 3000,
				},
				{
					Id:     5,
					Amount: 1000,
				},
				{
					Id:     4,
					Amount: 1000,
				},
			},
		},
		{
			name: "#3",
			args: args{
				transactions: []Transaction{
					{
						Id:     1,
						Amount: 100,
					},
				},
			},
			want: []Transaction{
				{
					Id:     1,
					Amount: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		if got := Sort(tt.args.transactions); !reflect.DeepEqual(tt.args.transactions, tt.want) {
			t.Errorf("Sort() = %v, want %v", got, tt.want)
		}
	}
}
