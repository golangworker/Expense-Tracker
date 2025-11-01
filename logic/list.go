package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const FILE_NAME = "file.json"

// тип для мапы, хранящая структуры расходов по ID
type exp map[int]expence

// структура для хранения списка расходов
type list struct {
	Expenses exp `json:"expenses"`
}

// функция для инициализации списка расходов
func InitList() list {
	return list{
		Expenses: make(exp),
	}
}

// метод для нового расхода в мапе
func (l *list) Add(e expence) {
	l.Expenses[e.ID] = e
}

// метод для обновления расхода
func (l *list) Update(id int, amount float64) error {
	if _, ok := l.Expenses[id]; ok {
		e := l.Expenses[id]
		e.Update(amount)
		l.Expenses[id] = e
		return nil
	}
	return fmt.Errorf("expense with id %d not found", id)
}

// метод для удаления расхода
func (l *list) Delete(id int) error {
	if _, ok := l.Expenses[id]; ok {
		delete(l.Expenses, id)
		return nil
	}
	return fmt.Errorf("expense with id %d not found", id)
}

// метод для просмотра всех расходов
func (l *list) ShowAllExpenses() {
	fmt.Printf("%-5s %-11s %-20s %-s\n", "ID", "Date", "Description", "Amount")
	for _, v := range l.Expenses {
		fmt.Println(v)
	}
}

// метод для просмотра расходов определенного месяца
func (l *list) ShowForMonthExpenses(month string) error {
	// проверка, что пользователь ввел месяц, а не что-то другое
	data, err := time.Parse(time.January.String(), month)
	if err != nil {
		return err
	}
	curYear := time.Now().Year()
	fmt.Printf("%-5s %-11s %-20s %-s\n", "ID", "Date", "Description", "Amount")
	for _, v := range l.Expenses {
		if v.Date.Year() == curYear && v.Date.Month() == data.Month() {
			fmt.Println(v)
		}
	}
	return nil
}

// метод для загрузки данных из файла
func (l *list) LoadFromFile() error {
	_, err := os.Stat(FILE_NAME)
	switch err {
	case nil:
		data, _ := os.ReadFile(FILE_NAME)
		json.Unmarshal(data, &l.Expenses)
		return nil
	case os.ErrNotExist:
		f, err := os.Create(FILE_NAME)
		if err != nil {
			return err
		}
		defer f.Close()
		return nil
	default:
		return fmt.Errorf("error loading data from file: %w", err)
	}
}

// метод для сохранения данных
func (l *list) SaveToFile() error {
	data, err := json.MarshalIndent(l.Expenses, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}
	os.WriteFile(FILE_NAME, data, 0644)
	return nil
}
