package main

import (
	"bgo-homeworks-06/pkg/card"
	"bgo-homeworks-06/pkg/processing"
	"fmt"
)

func main() {

	transactions := generateTransactions(100)

	months, _ := processing.GetSumByMonthList(transactions, "2020-01-01", "2020-07-31")

	fmt.Println(months)
}

func generateTransactions(count int) []card.Transaction {
	var transactions []card.Transaction

	timeStamp := 1577836800 // 01.01.2020

	for i := 0; i < count; i++ {
		timeStamp = timeStamp + 60*60*24*7
		transactions = append(transactions, card.Transaction{
			Id:        int64(i),
			Amount:    1,
			Timestamp: int64(timeStamp),
		})
	}
	return transactions
}
