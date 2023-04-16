package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	chooseAction       = "Write action (buy, fill, take, remaining, exit): "
	buyChoice          = "What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - back:"
	fillWaterQuestion  = "Write how many ml of water you want to add:"
	fillMilkQuestion   = "Write how many ml of milk you want to add: "
	fillCoffeeQuestion = "Write how many grams of coffee beans you want to add: "
	fillCupsQuestion   = "Write how many disposable cups you want to add: "
	espressoCost       = 4
	latteCost          = 7
	cappuccinoCost     = 6
	espressoWater      = 250
	espressoCoffee     = 16
	latteWater         = 350
	latteMilk          = 75
	latteCoffee        = 20
	cappuccinoWater    = 200
	cappuccinoMilk     = 100
	cappuccinoCoffee   = 12
)

func showInfo(water, milk, coffee, cup *int, money *float64) string {
	return "The coffee machine has: \n" +
		strconv.Itoa(*water) + " ml of water\n" +
		strconv.Itoa(*milk) + " ml of milk\n" +
		strconv.Itoa(*coffee) + " g of coffee beans\n" +
		strconv.Itoa(*cup) + " disposable cups\n" +
		"$" + strconv.FormatFloat(*money, 'g', 5, 64) + " of money\n"
}

func buy(cofType string, water, milk, coffee, cup *int, money *float64) string {
	var runOutProd string
	var isOk bool
	successMessage := "I have enough resources, making you a coffee!"
	switch cofType {
	case "1":
		if *water >= espressoWater && *coffee >= espressoCoffee {
			*water -= espressoWater
			*coffee -= espressoCoffee
			*money += espressoCost
			isOk = true
		} else if *water < espressoWater {
			runOutProd = "water"
			isOk = false
		} else if *coffee < espressoCoffee {
			runOutProd = "coffee"
			isOk = false
		}

	case "2":
		if *water >= latteWater && *coffee >= latteCoffee && *milk >= latteMilk {
			*water -= latteWater
			*milk -= latteMilk
			*coffee -= latteCoffee
			*money += latteCost
			isOk = true
		} else if *water < latteWater {
			runOutProd = "water"
			isOk = false
		} else if *coffee < latteCoffee {
			runOutProd = "coffee"
			isOk = false
		} else if *milk < latteMilk {
			runOutProd = "milk"
			isOk = false
		}
	case "3":
		if *water >= cappuccinoWater && *coffee >= cappuccinoCoffee && *milk >= cappuccinoMilk {
			*water -= cappuccinoWater
			*milk -= cappuccinoMilk
			*coffee -= cappuccinoCoffee
			*money += cappuccinoCost
			isOk = true
		} else if *water < cappuccinoWater {
			runOutProd = "water"
			isOk = false
		} else if *coffee < cappuccinoCoffee {
			runOutProd = "coffee"
			isOk = false
		} else if *milk < cappuccinoMilk {
			runOutProd = "milk"
			isOk = false
		}
	case "4":
		mainMenu(water, milk, coffee, cup, money)
	}
	if isOk {
		*cup -= 1
		return successMessage
	} else {
		return "Sorry, not enough " + runOutProd + "!"
	}

}

func fill(water, milk, coffee, cup *int) {
	addWater, _ := strconv.Atoi(scanChoice(fillWaterQuestion))
	addMilk, _ := strconv.Atoi(scanChoice(fillMilkQuestion))
	addCoffee, _ := strconv.Atoi(scanChoice(fillCoffeeQuestion))
	addCup, _ := strconv.Atoi(scanChoice(fillCupsQuestion))
	*water += addWater
	*milk += addMilk
	*coffee += addCoffee
	*cup += addCup
}

func take(money *float64) string {
	var resp = "I gave you " + "$" + strconv.FormatFloat(*money, 'g', 5, 64)
	*money -= *money
	return resp
}

func scanChoice(question string) string {
	var choice string
	fmt.Println(question)
	fmt.Scanln(&choice)
	return choice
}

func mainMenu(water, milk, coffee, cup *int, money *float64) {
	for {
		action := scanChoice(chooseAction)
		switch action {
		case "buy":
			fmt.Println(buy(scanChoice(buyChoice), water, milk, coffee, cup, money))
		case "fill":
			fill(water, milk, coffee, cup)
		case "take":
			fmt.Println(take(money))
		case "remaining":
			fmt.Println(showInfo(water, milk, coffee, cup, money))
		case "exit":
			os.Exit(0)
		}
	}
}

func main() {
	var waterLevel int = 400
	var milkLevel int = 540
	var coffeeLevel int = 120
	var cupAvail int = 9
	var money float64 = 550
	mainMenu(&waterLevel, &milkLevel, &coffeeLevel, &cupAvail, &money)
}
