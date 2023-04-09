package service

import (
	"math"
	"platform_science_code_exercise/internal/domain/entity"
)

func PlatformScienceCodeExercise(m entity.Message) []entity.MessageResponse {
	var result []entity.MessageResponse
	// sample data
	shipments := m.Address
	drivers := m.Driver

	// assign shipments to drivers
	assignments := assignShipments(shipments, drivers)

	// print results
	for i, a := range assignments {
		r := entity.MessageResponse{
			Id:      i,
			Driver:  a.driver,
			Address: a.shipment,
		}
		result = append(result, r)
	}

	return result
}

func assignShipments(shipments []string, drivers []string) []Assignment {
	assignments := make([]Assignment, len(shipments))
	usedDrivers := make(map[int]bool)

	for i, s := range shipments {
		bestScore := math.Inf(-1)
		var bestDriver string

		for j, d := range drivers {
			if usedDrivers[j] {
				continue
			}

			ss := calculateSuitabilityScore(s, d)
			if ss > bestScore {
				bestScore = ss
				bestDriver = d
			}
		}

		assignments[i] = Assignment{shipment: s, driver: bestDriver}
		usedDrivers[findDriverIndex(drivers, bestDriver)] = true
	}

	return assignments
}

func calculateSuitabilityScore(s string, d string) float64 {
	sn := len(s)
	dn := len(d)

	// calculate base SS
	var baseSS float64
	if sn%2 == 0 {
		baseSS = float64(countVowels(d)) * 1.5
	} else {
		baseSS = float64(countConsonants(d))
	}

	// apply factor bonus
	if findCommonFactor(sn, dn) > 1 {
		baseSS *= 1.5
	}

	return baseSS
}

func countVowels(s string) int {
	count := 0
	for _, r := range s {
		if isVowel(r) {
			count++
		}
	}
	return count
}

func countConsonants(s string) int {
	count := 0
	for _, r := range s {
		if !isVowel(r) {
			count++
		}
	}
	return count
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func findCommonFactor(a, b int) int {
	for i := 2; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

func findDriverIndex(drivers []string, driver string) int {
	for i, d := range drivers {
		if d == driver {
			return i
		}
	}
	return -1
}

type Assignment struct {
	shipment string
	driver   string
}
