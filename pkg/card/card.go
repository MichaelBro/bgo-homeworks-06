package card

import "sort"

type Transaction struct {
	Id     int64
	Amount int64
}

func Sort(transactions []Transaction) []Transaction {
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Amount > transactions[j].Amount
	})
	return transactions
}
