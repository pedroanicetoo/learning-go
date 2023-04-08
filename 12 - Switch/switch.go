package main

import "fmt"

func WeekDay(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid Number"
	}
}

// also

func WeekDay2(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Sunday"
		fallthrough
	case number == 2:
		weekDay = "Monday"
	case number == 3:
		weekDay = "Tuesday"
	case number == 4:
		weekDay = "Wednesday"
	case number == 5:
		weekDay = "Thursday"
	case number == 6:
		weekDay = "Friday"
	case number == 7:
		weekDay = "Saturday"
	default:
		weekDay = "Invalid Number"
	}

	return weekDay
}

func main() {
	fmt.Println("Switch")

	day := WeekDay(200)

	fmt.Println(day) // Invalid Number

	day2 := WeekDay2(7)

	fmt.Println(day2) // Saturday

	// using fallthrough clause
	day3 := WeekDay2(1)

	fmt.Println(day3) // Monday

}
