package spentenergy

import (
	"errors"
	"time"
)

const (
	stepLengthCoefficient = 0.65
	mInKm                 = 1000.0
	minInH                = 60.0
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

	// Исправленная формула для бега
	speed := distance / hours
	calories := 0.035*weight + (speed*speed/height)*0.029*weight
	calories *= hours

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

	// Исправленная формула для ходьбы
	speed := distance / hours
	calories := (0.035 * weight) + (speed*speed/height)*0.029*weight
	calories *= hours

	return calories, nil
}
