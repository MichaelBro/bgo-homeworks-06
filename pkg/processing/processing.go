package processing

import (
	"bgo-homeworks-06/pkg/card"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Months struct {
	List map[string][]card.Transaction
}

func NewList() *Months {
	return &Months{
		List: make(map[string][]card.Transaction),
	}
}

func (m *Months) GetSumByMonthsList(transactions []card.Transaction, startDate, endDate string) error {
	startDateTime, endDateTime, errorPrepare := getPrepareDate(startDate, endDate)
	if errorPrepare != nil {
		return errorPrepare
	}

	m.createListMonths(startDateTime, endDateTime)
	m.distributionOfTransactionsByMonths(transactions, startDateTime.Unix(), endDateTime.Unix())
	m.showSumConcurrently()
	return nil
}

func (m *Months) distributionOfTransactionsByMonths(transactions []card.Transaction, start int64, end int64) *Months {
	for _, transaction := range transactions {
		timestamp := transaction.Timestamp
		if timestamp >= start && timestamp <= end {
			date := time.Unix(timestamp, 0).Format("2006-01")
			m.List[date] = append(m.List[date], transaction)
		}
	}
	return m
}

func (m *Months) createListMonths(startDateTime, endDateTime time.Time) map[string][]card.Transaction {
	next := startDateTime
	for next.Before(endDateTime) {
		m.List[next.Format("2006-01")] = nil
		next = next.AddDate(0, 1, 0)
	}
	return m.List
}

func (m *Months) showSumConcurrently() {
	wg := sync.WaitGroup{}
	wg.Add(len(m.List))

	for s, t := range m.List {
		str := s
		trn := t
		go func() {
			sum := sumTransactionsByMonth(trn)
			fmt.Println("Sum by:", str, "Amount:", sum)
			wg.Done()
		}()
	}
	wg.Wait()
}

func sumTransactionsByMonth(transactions []card.Transaction) int64 {
	sum := int64(0)
	for _, transaction := range transactions {
		sum += transaction.Amount
	}
	return sum
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
	endDateTime = endDateTime.AddDate(0, 0, 1).Add(-time.Second) // need time 2020-07-31 23:59:59

	if startDateTime.Unix() > endDateTime.Unix() {
		return time.Time{}, time.Time{}, ErrIncorrectDate
	}

	return startDateTime, endDateTime, nil
}
