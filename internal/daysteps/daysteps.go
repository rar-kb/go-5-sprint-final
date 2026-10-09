package daysteps

import (
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
	dtstring := strings.Split(datastring, ",")
	if len(dtstring) != 2 {
		return fmt.Errorf("получено некорретное количество данных")
	}
	steps, err := strconv.Atoi(dtstring[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов не может быть отрицательным")
	}
	duration, err := time.ParseDuration(dtstring[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("время не может быть отрицательным")
	}
	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("количество шагов должно быть положительным: %d", ds.Steps)
	}
	if ds.Duration <= 0 {
		return "", fmt.Errorf("продолжительность должна быть положительной: %s", ds.Duration)
	}
	if ds.Weight <= 0 {
		return "", fmt.Errorf("вес должен быть положительным: %f", ds.Weight)
	}
	if ds.Height <= 0 {
		return "", fmt.Errorf("рост должен быть положительным: %f", ds.Height)
	}

	distant := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distant, calories), nil
}
