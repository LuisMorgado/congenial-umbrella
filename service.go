package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(myItems []Item, count int) ([]Item, int) {
	count++
	if count%3 == 0 && count%5 == 0 {
		myItems = append(myItems, Item{Id: count, Name: "fizzbuzz"})
		fmt.Println(count, "fizz"+"buzz")
		return myItems, count
	}
	if count%3 == 0 {
		myItems = append(myItems, Item{Id: count, Name: "fizz"})
		fmt.Println(count, "fizz")
		return myItems, count
	}
	if count%5 == 0 {
		myItems = append(myItems, Item{Id: count, Name: "buzz"})
		fmt.Println(count, "buzz")
		return myItems, count
	}
	myItems = append(myItems, Item{Id: count, Name: "item"})
	fmt.Println("item")
	return myItems, count
}

func getPagination(bounderies int, currentPage int, totalPages int, around int) (result string) {
	earlyElipse := false
	lateElipse := false

	for i := 1; i <= totalPages; i++ {
		if i == currentPage ||
			((i < currentPage && i >= currentPage-around) || (i > currentPage && i <= currentPage+around)) ||
			(i <= bounderies || i > totalPages-bounderies) {
			result += strconv.Itoa(i) + " "
			continue
		}

		if !earlyElipse && i > bounderies && i < currentPage {
			result += "... "
			earlyElipse = true
			continue
		}

		if !lateElipse && i > currentPage && i <= totalPages-bounderies {
			result += "... "
			lateElipse = true
			continue
		}

	}

	return result
}
