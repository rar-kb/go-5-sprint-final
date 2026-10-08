package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rar-kb/go-5-sprint-final/internal/personaldata"
	"github.com/rar-kb/go-5-sprint-final/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dtstring := strings.Split(datastring, ",")
	if len(dtstring) != 2 {
		return fmt.Errorf("Получено некорретное количество данных")
	}
	steps, err := strconv.Atoi(dtstring[0])
	if err != nil {
		return err
	}
	if steps < 0 {
		return fmt.Errorf("Количество шагов не может быть отрицательным")
	}
	ds.Steps = steps
	time, err := time.ParseDuration(dtstring[2])
	if err != nil {
		return err
	}
	if time < 0 {
		return fmt.Errorf("Время не может быть отрицательным")
	}
	ds.Duration = time
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distant := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("Количество шагов: %d.\n Дистанция составила %.2f км. \n Вы сожгли %.2f ккал.",
		ds.Steps, distant, calories), nil
}
