package term

import (
	"app/logic"
	"fmt"
	"strconv"
)

var ErrUseHelp error = fmt.Errorf("use \033[1m--help or -h\033[0m to show commands")

func RunningLoop(args []string) error {
	if len(args) < 2 {
		return ErrUseHelp
	}
	switch args[1] {
	case "--help", "-h":
		if len(args) != 2 {
			return ErrUseHelp
		}
		help()
	case "--add", "-a":
		if len(args) != 4 {
			return ErrUseHelp
		}
		description := args[2]
		amount, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			return fmt.Errorf("Invalid amount")
		}
		if err := add(description, amount); err != nil {
			return fmt.Errorf("Error adding expense: %v", err)
		}

	case "--delete", "-d":
		if len(args) != 3 {
			return ErrUseHelp
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("Invalid id")
		}
		if err := delete(id); err != nil {
			return fmt.Errorf("Error deleting expense: %v", err)
		}

	case "--edit", "-e":
	case "--summary", "-s":
		if len(args) != 2 {
			return ErrUseHelp
		}
		if err := summary(); err != nil {
			return fmt.Errorf("error showing summary: %w", err)
		}
	case "--month_summary", "-ms":
		if len(args) != 3 {
			return ErrUseHelp
		}
		if err := monthSummary(args[2]); err != nil {
			return fmt.Errorf("error showing month summary: %w", err)
		}
	default:
		return ErrUseHelp
	}
	return nil
}

func help() {
	fmt.Printf("Available commands:\n\n")
	fmt.Printf("• \033[1m--help, -h\033[0m info about program\n\n")

	fmt.Printf("• \033[1m--add, -a\033[0m (description string, amount float) create new expense\n\n")
	fmt.Printf("Example: app --add \"Groceries\" 50\n\n")

	fmt.Printf("• \033[1m--delete, -d\033[0m (id int) delete expense\n\n")
	fmt.Printf("Example: app --delete 1\n\n")

	fmt.Printf("• \033[1m--edit, -e\033[0m (id int, amount float) edit expense\n\n")
	fmt.Printf("Example: app --edit 1 50\n\n")

	fmt.Printf("• \033[1m--summary, -s\033[0m show list of expenses\n\n")
	fmt.Printf("Example: app --summary\n\n")

	fmt.Printf("• \033[1m--month_summary, -ms\033[0m (month string) show list of expenses by month\n\n")
	fmt.Printf("Example: app --month_summary \"January\"\n\n\n")

	fmt.Printf("Last version: \033[1m2.11.2025 github.com/golangworker/Expense-Tracker\033[0m\n\n")
}

// функция для добавления расхода
func add(description string, amount float64) error {
	expence := logic.InitExpence(description, amount)
	list := logic.InitList()
	if err := list.LoadFromFile(); err != nil {
		return err
	}
	list.Add(expence)
	if err := list.SaveToFile(); err != nil {
		return err
	}
	return nil
}

// функция для удаления расхода
func delete(id int) error {
	list := logic.InitList()
	list.LoadFromFile()
	if err := list.Delete(id); err != nil {
		return err
	}
	if err := list.SaveToFile(); err != nil {
		return err
	}
	return nil
}

// функция для редактирования расхода
func edit(id int, amount float64) error {
	list := logic.InitList()
	if err := list.LoadFromFile(); err != nil {
		return err
	}
	if err := list.Update(id, amount); err != nil {
		return err
	}
	if err := list.SaveToFile(); err != nil {
		return err
	}
	return nil
}

// функция для просмотра расходов
func summary() error {
	list := logic.InitList()
	if err := list.LoadFromFile(); err != nil {
		return err
	}
	list.ShowAllExpenses()
	return nil
}

// функция для просмотра расходов за месяц
func monthSummary(month string) error {
	list := logic.InitList()
	if err := list.LoadFromFile(); err != nil {
		return err
	}
	if err := list.ShowForMonthExpenses(month); err != nil {
		return err
	}
	return nil
}
