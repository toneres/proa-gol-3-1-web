package employees

type Employee struct {
	ID     int64
	Name   string `validate:"required,min=3,max=160"`
	Gender string `validate:"required,oneof=male female"`
}
