package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/dasensy1/lab4-variant03/pkg/salarycalc"
	"github.com/dasebsy1/lab4-variant03/pkg/waterbill" // твой пакет
)

func main() {

	hours := 160.0
	rate := 15.0
	tax := 10.0
	bonus := 5.0

	// 1. Gross Salary
	gross, err := salarycalc.GrossSalary(hours, rate)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Net Salary
	net, err := salarycalc.NetSalary(gross, tax)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Apply Bonus
	err = salarycalc.ApplyBonus(&net, bonus)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Salary Report
	report, err := salarycalc.FormatSalaryReport("Alikhan", gross, net)
	if err != nil {
		log.Fatal(err)
	}

	color.Cyan("=== Salary Report ===")
	fmt.Println(report)

	// 5. Example usage of your GitHub package
	waterCost := waterbill.CalculateWaterBill(25, 150)
	color.Green("=== Water Bill Example ===")
	fmt.Printf("Water bill from my GitHub: %.2f\n", waterCost)
}
