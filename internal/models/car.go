package models


type Answer struct {
	Status string
}
type People struct {
	Name       string `json:"name" db:"name" binding:"required"`
	Surname    string `json:"surname" db:"surname" binding:"required"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}

type Car struct {
	RegNum string `json:"regNum" db:"regnum" binding:"required"`
	Mark   string `json:"mark" db:"mark" binding:"required"`
	Model  string `json:"model" db:"model" binding:"required"`
	Year   int    `json:"year" db:"year"`
	Owner  People
}

type UpdatePeople struct {
	Name       *string `json:"name" db:"name"`
	Surname    *string `json:"surname" db:"surname" `
	Patronymic *string `json:"patronymic" db:"patronymic"`
}

type UpdateCar struct {
	RegNum *string `json:"regNum" db:"regNum" `
	Mark   *string `json:"mark" db:"mark"`
	Model  *string `json:"model" db:"model"`
	Year   *int    `json:"year" db:"year"`
	Owner  *UpdatePeople
}

type GetRegNums struct {
	RegNums []string `json:"regNums"`
}
type GetAllCarResponse struct {
	Data []Car `json:"data"`
}
