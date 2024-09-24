package model

type Todo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"type:varchar(255)" validate:"required"`
	Done        bool   `gorm:"default:false"`
}

func (Todo) TableName() string {
	return "public.todos"
}
