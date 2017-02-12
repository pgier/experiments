package main

import (
	"fmt"
)

const costcat float32 = 10.00
const costfish float32 = 3.00
const costbird float32 = 0.50

const totalcost float32 = 100.00
const totalnum int = 100

const maxcats int = (int)(totalcost/costcat)

func main() {
	calc()
}

func calc() {
	var numcat int = 1
	var numfish int = 1
	var numbird int = 1

	var currentcost float32 = 0
	var catfishcost float32 = 0

	calcloop:
	for(numcat < maxcats) {
		for(currentcost < totalcost) {
		        catfishcost = costcat * float32(numcat) + costfish * float32(numfish)
			numbird = int((totalcost - catfishcost)/costbird)
			//printNums(numcat, numfish, numbird)
			currentcost = costcat * float32(numcat) + costfish * float32(numfish) + costbird * float32(numbird)
			if (currentcost == totalcost && (numcat + numfish + numbird) == totalnum) {
				fmt.Println("Found solution")
				break calcloop
			}
			numfish++
			numbird = 1
			currentcost = costcat * float32(numcat) + costfish * float32(numfish) + costbird * float32(numbird)
		}
		numcat++
		numbird = 1
		numfish = 1
		currentcost = costcat * float32(numcat) + costfish * float32(numfish) + costbird * float32(numbird)
	}
	if ((numcat + numfish + numbird) != totalnum) {
		fmt.Printf("No solution found")
	}

	printNums(numcat, numfish, numbird)
}

func printNums(numcat int, numfish int, numbird int) {
	fmt.Printf("Num of cats: %v\n", numcat)
	fmt.Printf("Num of fish: %v\n", numfish)
	fmt.Printf("Num of birds: %v\n", numbird)
}

