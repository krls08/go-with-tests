package main

import (
	"fmt"
	"regexp"
)

func main() {

	dateIc := "2023-01-09"
	//timeIc := "12:02:34"
	timeIc := "23:32:59"
	dateOk := regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`).MatchString(dateIc)
	if !dateOk {
		fmt.Println("date doesnt match regex expresion")
	}
	timeOk := regexp.MustCompile(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`).MatchString(timeIc)
	//timeOk := regexp.MustCompile(`^\d+:\d{2}:\d{2}$`).MatchString(timeIc)
	if !timeOk {
		fmt.Println("time doesnt match regex expresion")
	}

	//dateFormat := "2017-01-01"
	//timeFormat := "00:00:00"
	//_, err := time.Parse(dateFormat, dateIc)
	//if err != nil {
	//	fmt.Println("dateFormat error:", err.Error())
	//}
	//_, err = time.Parse(timeFormat, timeIc)
	//if err != nil {
	//	fmt.Println("timeFormat error:", err.Error())
	//}
}
