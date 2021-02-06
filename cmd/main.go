package main

import (
	"github.com/MichaelBro/bgo-homeworks-06/pkg/card"
	"github.com/MichaelBro/bgo-homeworks-06/pkg/processing"
)

func main() {

	transactions := generateTransactions(10000)

	list := processing.NewList()
	_ = list.GetSumByMonthsList(transactions, "2020-01-01", "2020-07-31")
}

func generateTransactions(count int) []card.Transaction {
	var transactions []card.Transaction

	timeStamp := 1577836800 // 01.01.2020

	for i := 0; i < count; i++ {
		timeStamp = timeStamp + 60*60*12
		transactions = append(transactions, card.Transaction{
			Id:        int64(i),
			Amount:    1,
			Timestamp: int64(timeStamp),
		})
	}
	return transactions
}
