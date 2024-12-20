package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var UndefinedDateFormat = errors.New("undefined date format")

func main() {
	dates := []string{
		"12.09.1978",  // dd.mm.YYYY
		"1990/06/10",  // YYYY/mm/dd
		"08.03.2021",  // dd.mm.YYYY
		"12.04.1986",  // dd.mm.YYYY
		"25 dec 1988", // Не поддерживаемый формат
		"2001/05/25",  // YYYY/mm/dd
	}

	for _, d := range dates {
		year, month, day, err := parseDate(d)
		if err != nil {
			fmt.Println("ERROR!", err, "-", d)
			continue
		}

		fmt.Printf("%04d.%02d.%02d\n", year, month, day)
	}
}

func parseDate(date string) (year, month, day int64, err error) {
	// Регулярное выражения для формата dd.mm.YYYY
	reDDMMYYYY := regexp.MustCompile(`^(\d{2})\.(\d{2})\.(\d{4})$`)
	// Регулярное выражения для формата YYYY/mm/dd
	reYYYYMMDD := regexp.MustCompile(`^(\d{4})/(\d{2})/(\d{2})$`)

	// Проверяем формат dd.mm.YYYY
	if matches := reDDMMYYYY.FindStringSubmatch(date); matches != nil {
		day = strToInt(matches[1])
		month = strToInt(matches[2])
		year = strToInt(matches[3])
		return
	}

	// Проверяем формат YYYY/mm/dd
	if matches := reYYYYMMDD.FindStringSubmatch(date); matches != nil {
		day = strToInt(matches[3])
		month = strToInt(matches[2])
		year = strToInt(matches[1])
		return
	}

	err = UndefinedDateFormat
	return
}

func strToInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
