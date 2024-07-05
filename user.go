package timetracker

type User struct {
	Id            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name" binding:"required"`
	PasportNumber string `json:"pasportNumber" db:"pasportnumber" binding:"required"`
}

type UpdateUserInput struct {
	Name          *string `json:"name" example:"Иван"`
	PasportNumber *string `json:"pasportNumber" example:"1234 123456"`
}

type CreateUserInput struct {
	Name          string `json:"name" binding:"required" example:"Иван"`
	PasportNumber string `json:"pasportNumber" binding:"required" example:"1234 123456"`
}
