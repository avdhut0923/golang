package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	// fmt.Println("current time is ", currentTime)
	// fmt.Printf("type of  time is %T\n", currentTime)

	// formatted := currentTime.Format("2006/01/02,Monday") // not even monday is working it requir exact formate as Monday
	// fmt.Println("Formated time is ", formatted)

	layout_str := "02.01.2006"
	datestr := "25.11.2024"
	formatted_time, _ := time.Parse(layout_str, datestr)
	fmt.Println("Formated time is ", formatted_time)

	new_date := currentTime.Add(48 * time.Hour)
	fmt.Println("new_date time:", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday ")
	fmt.Println("formatted_new_date time :", formatted_new_date)

}
