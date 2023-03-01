package main

type Error struct {
	Deployment string
	Error      error
}

type Errors []*Error

func (errs Errors) Error() string {
	return ""
}
