package spentenergy

import (
	"errors"
	"time"
)

const (
	stepLengthCoefficient = 0.725 // Исправленный коэффициент длины шага
	mInKm                 = 1000.0
)

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm
	return distance
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}

	return distance / hours
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}

	distance := Distance(steps, height)
	hours := duration.Hours()

	// Исправленная формула для бега (MET * вес * время)
	met := 9.8 // MET для бега со скоростью ~8 км/ч
	calories := met * weight * hours

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}

	distance := Distance(steps, height)
	hours := duration.Hours()

	// Исправленная формула для ходьбы (MET * вес * время)
	met := 4.0 // MET для быстрой ходьбы
	calories := met * weight * hours

	return calories, nil
}
