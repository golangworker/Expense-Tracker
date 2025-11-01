package term

import (
	"app/logic"
	"fmt"
	"strconv"
)

var ErrUseHelp error = fmt.Errorf("use --help or -h to show commands")

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
	default:
		return ErrUseHelp
	}
	return nil
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  --help, -h")
	fmt.Println("  --add, -a")
	fmt.Println("  --delete, -d")
	fmt.Println("  --edit, -e")
	fmt.Println("  --summary, -s")
	fmt.Println("  --month_summary, -ms")
}

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

func summary() error {
	list := logic.InitList()
	if err := list.LoadFromFile(); err != nil {
		return err
	}
	list.ShowAllExpenses()
	return nil
}
