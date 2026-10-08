package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("Количество шагов должно быть больше нуля")
	}
	// проверяем что вес пользователя не меньше 0
	if weight <= 0 {
		return 0, fmt.Errorf("Вес не может быть отрицательным")
	}
	// проверяем что рост пользователяя не меньше 0
	if height <= 0 {
		return 0, fmt.Errorf("Рост не может быть отрицательным")
	}
	// проверяем что время не отрицательное
	if duration <= 0 {
		return 0, fmt.Errorf("Время не может быть отрицательным")
	}
	// mnSpeed - средняя скорость
	mnSpeed := MeanSpeed(steps, height, duration)
	// durationMinutes - время в минутах
	durationMinutes := duration.Minutes()
	result := ((weight * mnSpeed * durationMinutes) / 60) * walkingCaloriesCoefficient
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверяем что кол-во шагов не меньше 0
	if steps <= 0 {
		return 0, fmt.Errorf("Количество шагов должно быть больше нуля")
	}
	// проверяем что вес пользователя не меньше 0
	if weight <= 0 {
		return 0, fmt.Errorf("Вес не может быть отрицательным")
	}
	// проверяем что рост пользователяя не меньше 0
	if height <= 0 {
		return 0, fmt.Errorf("Рост не может быть отрицательным")
	}
	// проверяем что время не отрицательное
	if duration <= 0 {
		return 0, fmt.Errorf("Время не может быть отрицательным")
	}
	// mnSpeed - средняя скорость
	mnSpeed := MeanSpeed(steps, height, duration)
	// durationMinutes - время в минутах
	durationMinutes := duration.Minutes()
	// считаем результат
	result := (weight * mnSpeed * durationMinutes) / minInH
	return result, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверка что кол-во шагов не отрицательно
	if steps < 0 {
		return 0
	}
	// проверка что время не отрицательное
	if duration <= 0 {
		return 0
	}
	// Dist - пройденная дистанция
	Dist := Distance(steps, height)
	// возвращаем среднюю скорость
	return Dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps < 0 {
		return 0
	}
	if height < 0 {
		return 0
	}
	// lenSh - длина шага
	lenSh := height * stepLengthCoefficient
	return ((float64(steps) * lenSh) / mInKm)
}
