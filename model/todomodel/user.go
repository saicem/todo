package todomodel

type User struct {
	Uid       uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	NickName  string
	Password  string
	TodoItems []TodoItem `gorm:"foreignKey:Uid"`
}
