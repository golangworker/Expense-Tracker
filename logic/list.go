package logic

type exp map[int]expence

type list struct {
	Expenses exp `json:"expenses"`
}

func InitList() list {
	return list{}
}
