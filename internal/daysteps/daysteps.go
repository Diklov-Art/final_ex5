package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат данных: ожидается 2 части")
	}

	stepsStr := parts[0]
	if stepsStr == "" {
		return errors.New("количество шагов не может быть пустым")
	}

	if strings.TrimSpace(stepsStr) != stepsStr {
		return errors.New("неверный формат шагов: пробелы в начале или конце не допускаются")
	}

	hasNonDigit := false
	for i, char := range stepsStr {
		if char == '+' || char == '-' {

			if i != 0 {
				hasNonDigit = true
				break
			}
		} else if char < '0' || char > '9' {
			hasNonDigit = true
			break
		}
	}

	if hasNonDigit {
		return errors.New("неверный формат шагов: должны быть только цифры, возможен знак в начале")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %v", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	durationStr := parts[1]
	if durationStr == "" {
		return errors.New("продолжительность не может быть пустой")
	}

	// Проверка на пробелы в начале или конце продолжительности
	if strings.TrimSpace(durationStr) != durationStr {
		return errors.New("неверный формат продолжительности: пробелы в начале или конце не допускаются")
	}

	// Парсинг продолжительности в формате time.Duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга продолжительности: %v", err)
	}

	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	info += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	info += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return info, nil
}
