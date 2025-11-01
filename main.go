package main

import (
	"app/logic"
)

func main() {
	expence := logic.InitExpence("Sosal", 6.66)
	list := logic.InitList()
	list.LoadFromFile()
	list.Add(expence)
	list.ShowAllExpenses()
	list.SaveToFile()
}
