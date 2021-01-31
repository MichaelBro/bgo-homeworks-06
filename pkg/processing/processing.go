package processing

import (
	"bgo-homeworks-06/pkg/card"
	"errors"
	"time"
)

type SumList struct {
	Month        int
	SumByMonth   int
	Transactions []card.Transaction
}

func showSumByMonthList(transactions []card.Transaction, startDate, endDate string) ([]SumList, error) {

	startDateTime, errStartDate, errorPrepare := prepareDate(startDate, endDate)

	if errorPrepare != nil {
		return nil, errorPrepare
	}





}



var (
	layout           = "2006-01-02"
	ErrIncorrectDate = errors.New("incorrect date")
)

func prepareDate(startDate, endDate string) (startDateTime, endDateTime time.Time, err error) {

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
