package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат данных: ожидается 3 части")
	}

	stepsStr := strings.TrimSpace(parts[0])
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
	t.Steps = steps

	trainingType := strings.TrimSpace(parts[1])
	if trainingType == "" {
		return errors.New("тип тренировки не может быть пустым")
	}

	if strings.TrimSpace(trainingType) != trainingType {
		return errors.New("неверный формат типа тренировки: пробелы в начале или конце не допускаются")
	}
	t.TrainingType = trainingType

	durationStr := strings.TrimSpace(parts[2])
	if durationStr == "" {
		return errors.New("продолжительность не может быть пустой")
	}

	if strings.TrimSpace(durationStr) != durationStr {
		return errors.New("неверный формат продолжительности: пробелы в начале или конце не допускаются")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга продолжительности: %v", err)
	}

	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) {
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType)
	info += fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours())
	info += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
	info += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	info += fmt.Sprintf("Сожгли калорий: %.2f\n", calories)

	return info, nil
}
