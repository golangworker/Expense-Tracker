package logic

import (
	"fmt"
	"math/rand"
	"time"
)

type expence struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

// функция для инициализации структуры расхода
func Init(description string, amount float64) expence {
	return expence{
		ID:          rand.Intn(10_000),
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
}

// setter для обновления структуры
func (e *expence) Update(amount float64) {
	e.Amount += amount
}

// getter для вывода всей структуры
func (e expence) String() string {
	return fmt.Sprintf("%-4d %-s %-s $%-.2f", e.ID, e.Date.Format("2006-01-02"), e.Description, e.Amount)
}
