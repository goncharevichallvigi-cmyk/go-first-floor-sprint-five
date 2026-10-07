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
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	slcData := strings.Split(datastring, ",")
	if len(slcData) != 2 {
		return errors.New("некорректное содержимое datastring")
	}
	steps, err := strconv.Atoi(slcData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("некорректное количество шагов")
	}

	spendTime, err := time.ParseDuration(slcData[1])
	if err != nil {
		return err
	}
	if spendTime <= 0 {
		return errors.New("некорректное время")
	}
	ds.Steps = steps
	ds.Duration = spendTime
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if ds.Steps == 0 {
		return "", errors.New("некорректное количество шагов")
	}
	numberOfCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	resultString := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, numberOfCalories)
	return resultString, nil
}
