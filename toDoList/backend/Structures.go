package main

type List struct {
	ID			int		`json:"id"`
	Task		string	`json:"task"`
	Completed	bool	`json:"completed"`
}
