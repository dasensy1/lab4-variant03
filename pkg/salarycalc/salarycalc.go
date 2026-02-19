// Package salarycalc предоставляет функции для расчета зарплаты сотрудника.
package salarycalc

import (
	"fmt"

	"github.com/dasensy1/lab4-variant03/pkg/waterbill" // твой GitHub-пакет
)


// GrossSalary вычисляет общую зарплату до вычета налогов.
func GrossSalary(hours, rate float64) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("hours cannot be negative")
	}
	if rate < 0 {
		return 0, fmt.Errorf("rate cannot be negative")
	}
	return hours * rate, nil
}

// NetSalary вычисляет чистую зарплату после вычета налога.
func NetSalary(gross, taxPercent float64) (float64, error) {
	if gross < 0 {
		return 0, fmt.Errorf("gross salary cannot be negative")
	}
	if taxPercent < 0 || taxPercent > 100 {
		return 0, fmt.Errorf("tax percent must be between 0 and 100")
	}

	net := gross * (1 - taxPercent/100)
	return net, nil
}

// ApplyBonus добавляет бонус к чистой зарплате через указатель.
func ApplyBonus(net *float64, bonusPercent float64) error {
	if net == nil {
		return fmt.Errorf("net salary pointer is nil")
	}
	if bonusPercent < 0 {
		return fmt.Errorf("bonus percent cannot be negative")
	}

	*net += *net * bonusPercent / 100
	return nil
}

// FormatSalaryReport формирует строку отчета по зарплате.
func FormatSalaryReport(employee string, gross, net float64) (string, error) {
	if employee == "" {
		return "", fmt.Errorf("employee name cannot be empty")
	}

	report := fmt.Sprintf(
		"Employee: %s\nGross Salary: %.2f\nNet Salary: %.2f\n",
		employee,
		gross,
		net,
	)

	// Пример использования функции из твоего GitHub-пакета
	waterCost := waterbill.CalculateWaterBill(10, 150) // пример вызова функции
	fmt.Printf("Example water bill from my GitHub: %.2f\n", waterCost)

	return report, nil
}
