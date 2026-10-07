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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	slcData := strings.Split(datastring, ",")
	if len(slcData) != 3 {
		return errors.New("некорректное содержимое datastring")
	}
	steps, err := strconv.Atoi(slcData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("некорректное количество шагов")
	}
	spendTime, err := time.ParseDuration(slcData[2])
	if err != nil {
		return err
	}
	if spendTime <= 0 {
		return errors.New("некорректное время")
	}
	t.Steps = steps
	t.TrainingType = slcData[1]
	t.Duration = spendTime
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var numberOfCalories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		numberOfCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		numberOfCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", nil
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	distance := spentenergy.Distance(t.Steps, t.Height)
	meenSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	resultString := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n", t.TrainingType, t.Duration.Hours(), distance)
	resultString += fmt.Sprintf("Скорость: %.2f км/ч\nСожгли калорий: %.2f\n", meenSpeed, numberOfCalories)
	return resultString, nil
}
