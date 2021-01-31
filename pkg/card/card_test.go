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
					{Id: 1, Amount: 100, Timestamp: 1612127900},
					{Id: 2, Amount: 3000, Timestamp: 1612127950},
					{Id: 3, Amount: 200, Timestamp: 1612128222},
					{Id: 4, Amount: 1000, Timestamp: 1612128422},
				},
			},
			want: []Transaction{
				{Id: 2, Amount: 3000, Timestamp: 1612127950},
				{Id: 4, Amount: 1000, Timestamp: 1612128422},
				{Id: 3, Amount: 200, Timestamp: 1612128222},
				{Id: 1, Amount: 100, Timestamp: 1612127900},
			},
		},
		{
			name: "#2",
			args: args{
				transactions: []Transaction{
					{Id: 1, Amount: 1000, Timestamp: 1612127900},
					{Id: 2, Amount: 3000, Timestamp: 1612127950},
					{Id: 4, Amount: 1000, Timestamp: 1612128222},
				},
			},
			want: []Transaction{
				{Id: 2, Amount: 3000, Timestamp: 1612127950},
				{Id: 1, Amount: 1000, Timestamp: 1612127900},
				{Id: 4, Amount: 1000, Timestamp: 1612128222},
			},
		},
		{
			name: "#3",
			args: args{
				transactions: []Transaction{
					{Id: 1, Amount: 1000, Timestamp: 1612127900},
				},
			},
			want: []Transaction{
				{Id: 1, Amount: 1000, Timestamp: 1612127900},
			},
		},
		{
			name: "#4",
			args: args{
				transactions: []Transaction{},
			},
			want: []Transaction{},
		},
	}
	for _, tt := range tests {
		if got := Sort(tt.args.transactions); !reflect.DeepEqual(tt.args.transactions, tt.want) {
			t.Errorf("Sort() = %v, want %v", got, tt.want)
		}
	}
}
