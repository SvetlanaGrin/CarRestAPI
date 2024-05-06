package models

type (
	Params struct {
		Pagination
		Filtration
	}
	Pagination struct {
		Limit  int
		Offset int
	}
	Filtration struct {
		Mark string
	}
)
