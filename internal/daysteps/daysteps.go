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

	stepsStr := strings.TrimSpace(parts[0])
	if stepsStr == "" {
		return errors.New("количество шагов не может быть пустым")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %v", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	durationStr := strings.TrimSpace(parts[1])
	if durationStr == "" {
		return errors.New("продолжительность не может быть пустой")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {

		if strings.Count(durationStr, ":") == 2 {
			timeParts := strings.Split(durationStr, ":")
			if len(timeParts) == 3 {
				hours, _ := strconv.Atoi(timeParts[0])
				minutes, _ := strconv.Atoi(timeParts[1])
				seconds, _ := strconv.Atoi(timeParts[2])
				duration = time.Duration(hours)*time.Hour +
					time.Duration(minutes)*time.Minute +
					time.Duration(seconds)*time.Second
			}
		} else {
			return fmt.Errorf("ошибка парсинга продолжительности: %v", err)
		}
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
