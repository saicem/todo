package todomodel

type User struct {
	Uid       uint `gorm:"primaryKey"`
	NickName  string
	Password  string
	TodoItems []TodoItem `gorm:"foreignKey:Uid"`
}
