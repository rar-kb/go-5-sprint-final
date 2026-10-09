package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rar-kb/go-5-sprint-final/internal/personaldata"
	"github.com/rar-kb/go-5-sprint-final/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dtstring := strings.Split(datastring, ",")
	// проверяем что получено нужное кол-во данных: "шаги, тип тренировки, время"
	if len(dtstring) != 3 {
		return fmt.Errorf("Введено некорретное количество данных")
	}
	steps, err := strconv.Atoi(dtstring[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов не может быть отрицательным")
	}
	duration, err := time.ParseDuration(dtstring[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("время не может быть отрицательным")
	}
	t.TrainingType = dtstring[1]
	t.Steps = steps
	t.Duration = duration
	return nil
}
func (t Training) ActionInfo() (string, error) {
	distant := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %q", t.TrainingType)
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distant, meanSpeed, calories), nil
}
