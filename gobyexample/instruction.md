
## This Folder contains the instructions for the Go by Example website.

- It will contain problems about the topics covered in the Go by Example website.
- Each topic will contain 5 real world problems.
- After Each Topic the solution is given



## Topic 1: Values

### Problem 1: Shopping Cart Calculation

A customer adds multiple items to their shopping cart in an online store. Each item has a price. Write a program that calculates the total cost of the items in the cart and displays it.

#### Solution:

```go
package main

import "fmt"

func main() {
	prices := []float64{15.99, 23.50, 9.99, 12.75}
	var total float64
	for _, price := range prices {
		total += price
	}
	fmt.Printf("Total cost: $%.2f\n", total)
}
```

### Problem 2: Employee Salary Adjustment

You need to adjust the salaries of employees based on their performance ratings. The program should calculate the updated salaries based on a percentage increase and display the results.

#### Solution:

```go
package main

import "fmt"

func main() {
	salaries := []float64{50000, 60000, 55000}
	percentageIncrease := []float64{5, 10, 7}

	for i, salary := range salaries {
		newSalary := salary + (salary * percentageIncrease[i] / 100)
		fmt.Printf("Updated salary for employee %d: $%.2f\n", i+1, newSalary)
	}
}
```

### Problem 3: Converting Temperatures

Write a program that converts a list of temperatures in Fahrenheit to Celsius and displays the results.

#### Solution:

```go
package main

import "fmt"

func main() {
	temperaturesF := []float64{32, 68, 104, 50}

	for _, tempF := range temperaturesF {
		tempC := (tempF - 32) * 5 / 9
		fmt.Printf("%.2f F is %.2f C\n", tempF, tempC)
	}
}
```

### Problem 4: Daily Water Consumption

A person records the amount of water they drink daily for a week. Write a program to calculate the total and average water consumption.

#### Solution:

```go
package main

import "fmt"

func main() {
	waterConsumption := []float64{1.5, 2.0, 1.8, 2.2, 1.9, 2.5, 2.3}
	var total float64
	for _, water := range waterConsumption {
		total += water
	}
	average := total / float64(len(waterConsumption))
	fmt.Printf("Total water consumption: %.2f L\n", total)
	fmt.Printf("Average water consumption: %.2f L/day\n", average)
}
```

### Problem 5: Bookstore Inventory

A bookstore keeps track of its inventory with the number of books in different genres. Write a program to calculate the total number of books available.

#### Solution:

```go
package main

import "fmt"

func main() {
	inventory := map[string]int{
		"Fiction":     120,
		"Non-Fiction": 80,
		"Mystery":     50,
		"Science":     40,
	}
	totalBooks := 0
	for _, count := range inventory {
		totalBooks += count
	}
	fmt.Printf("Total books in inventory: %d\n", totalBooks)
}
```

---

## Topic 2: Variables

### Problem 1: Car Rental Pricing

Write a program that calculates the total cost of renting a car for a given number of days. Use variables to store the daily rental rate and the number of days rented.

#### Solution:

```go
package main

import "fmt"

func main() {
	var dailyRate float64 = 30.0
	var daysRented int = 5
	totalCost := dailyRate * float64(daysRented)
	fmt.Printf("Total rental cost: $%.2f\n", totalCost)
}
```

### Problem 2: Electricity Bill Calculation

A household records its electricity usage in kWh. Write a program to calculate the monthly bill based on a rate per kWh.

#### Solution:

```go
package main

import "fmt"

func main() {
	var usage float64 = 350.5
	var ratePerKWh float64 = 0.12
	bill := usage * ratePerKWh
	fmt.Printf("Electricity bill: $%.2f\n", bill)
}
```

### Problem 3: Student Grade Calculation

Write a program that calculates the average grade of a student based on their scores in three subjects.

#### Solution:

```go
package main

import "fmt"

func main() {
	var math, science, english float64 = 85, 90, 88
	average := (math + science + english) / 3
	fmt.Printf("Average grade: %.2f\n", average)
}
```

### Problem 4: Bank Savings

A person deposits a certain amount in their bank account every month. Write a program to calculate the total savings after one year.

#### Solution:

```go
package main

import "fmt"

func main() {
	var monthlyDeposit float64 = 500
	totalSavings := monthlyDeposit * 12
	fmt.Printf("Total savings after one year: $%.2f\n", totalSavings)
}
```

### Problem 5: Distance Conversion

Write a program that converts a distance in kilometers to miles using a variable for the conversion factor.

#### Solution:

```go
package main

import "fmt"

func main() {
	var kilometers float64 = 10.0
	var conversionFactor float64 = 0.621371
	miles := kilometers * conversionFactor
	fmt.Printf("%.2f kilometers is %.2f miles\n", kilometers, miles)
}
```

---

## Topic 3: Constants

### Problem 1: Area of a Circle

Write a program to calculate the area of a circle using a constant for pi.

#### Solution:

```go
package main

import "fmt"

const Pi = 3.14159

func main() {
	radius := 5.0
	area := Pi * radius * radius
	fmt.Printf("Area of the circle: %.2f\n", area)
}
```

### Problem 2: Speed of Light

Using the constant speed of light (299,792,458 m/s), calculate how far light travels in one minute.

#### Solution:

```go
package main

import "fmt"

const SpeedOfLight = 299792458 // in meters per second

func main() {
	time := 60 // seconds
	distance := SpeedOfLight * time
	fmt.Printf("Light travels %d meters in one minute\n", distance)
}
```

### Problem 3: Tax Calculation

Define a constant tax rate and write a program to calculate the tax for a given income.

#### Solution:

```go
package main

import "fmt"

const TaxRate = 0.15

func main() {
	income := 45000.0
	tax := income * TaxRate
	fmt.Printf("Tax on an income of $%.2f is $%.2f\n", income, tax)
}
```

### Problem 4: Gravitational Force

Write a program to calculate the gravitational force between two masses using the gravitational constant (G = 6.67430e-11).

#### Solution:

```go
package main

import "fmt"

const G = 6.67430e-11

func main() {
	mass1 := 5.97e24 // Earth’s mass in kg
	mass2 := 7.35e22 // Moon’s mass in kg
	distance := 3.844e8 // Distance between Earth and Moon in meters

	force := G * (mass1 * mass2) / (distance * distance)
	fmt.Printf("Gravitational force between Earth and Moon: %.2e N\n", force)
}
```

### Problem 5: Currency Conversion

Use a constant conversion rate to convert an amount from USD to EUR.

#### Solution:

```go
package main

import "fmt"

const ConversionRate = 0.85

func main() {
	usd := 100.0
	eur := usd * ConversionRate
	fmt.Printf("$%.2f USD is equivalent to €%.2f EUR\n", usd, eur)
}
```

---

## Topic 4: Arrays

### Problem 1: Finding the Maximum Element

Write a program that finds the maximum value in an array of integers.

#### Solution:

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var max int = numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	println("Max number is: ", max)
}
```

### Problem 2: Reversing an Array

Write a program to reverse the elements of an array.

#### Solution:

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", numbers)

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	fmt.Println("Reversed array:", numbers)
}
```

### Problem 3: Summing Array Elements

Write a program to calculate the sum of all elements in an array.

#### Solution:

```go
package main

import "fmt"

func main() {
	numbers := []int{3, 7, 9, 12, 15}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("The sum of the array elements is: %d\n", sum)
}
```

### Problem 4: Checking for Duplicates

Write a program to check if an array contains any duplicate elements.

#### Solution:

```go
package main

import "fmt"

func main() {
	numbers := []int{4, 5, 9, 4, 7}
	seen := make(map[int]bool)
	hasDuplicates := false

	for _, num := range numbers {
		if seen[num] {
			hasDuplicates = true
			break
		}
		seen[num] = true
	}

	if hasDuplicates {
		fmt.Println("The array contains duplicates.")
	} else {
		fmt.Println("The array does not contain duplicates.")
	}
}
```

### Problem 5: Rotating an Array

Write a program to rotate an array to the right by a given number of positions.

#### Solution:

```go
package main

import "fmt"

func rotateRight(arr []int, positions int) []int {
	n := len(arr)
	positions = positions % n
	return append(arr[n-positions:], arr[:n-positions]...)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	positions := 2
	result := rotateRight(numbers, positions)
	fmt.Println("Original array:", numbers)
	fmt.Println("Rotated array:", result)
}
```

