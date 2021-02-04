package processing

import (
	"bgo-homeworks-06/pkg/card"
	"errors"
	"fmt"
	"time"
)

type SumByMonth struct {
	MonthNumber  int
	MonthName    string
	SumByMonth   int
	Transactions []card.Transaction
}

type Months struct {
	SumByMonthsList []SumByMonth
}

func (list *SumByMonth) AddTransaction(item card.Transaction) []card.Transaction {
	list.Transactions = append(list.Transactions, item)
	return list.Transactions
}

func (list *Months) AddMonth(item SumByMonth) []SumByMonth {
	list.SumByMonthsList = append(list.SumByMonthsList, item)
	return list.SumByMonthsList
}

func GetSumByMonthList(transactions []card.Transaction, startDate, endDate string) ([]Months, error) {

	startDateTime, endDateTime, errorPrepare := getPrepareDate(startDate, endDate)

	if errorPrepare != nil {
		return nil, errorPrepare
	}

	months := make(map[int64]string)
	next := startDateTime
	for next.Before(endDateTime) {
		months[next.Unix()] = next.Month().String()
		//months = append(months, next.Unix())
		next = next.AddDate(0, 1, 0)
	}

	fmt.Println(months)

	//for _, transaction := range transactions {
	//	if transaction.Timestamp >= startDateTime.Unix() && transaction.Timestamp <= endDateTime.Unix() {
	//
	//	}
	//}

	return nil, nil
}

var (
	layout           = "2006-01-02"
	ErrIncorrectDate = errors.New("incorrect date")
)

func getPrepareDate(startDate, endDate string) (startDateTime, endDateTime time.Time, err error) {

	startDateTime, err = time.Parse(layout, startDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	endDateTime, err = time.Parse(layout, endDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if startDateTime.Unix() > endDateTime.Unix() {
		return time.Time{}, time.Time{}, ErrIncorrectDate
	}

	return startDateTime, endDateTime, nil
}
