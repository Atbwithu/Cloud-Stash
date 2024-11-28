package pojo

type User struct {
	Id         int    `gorm:"primary_key"`
	Username   string `gorm:"not null"`
	Sex        string
	Age        int
	Password   string `gorm:"not null"`
	CreateTime string
}

func (user User) TableName() string {
	return "user"
}
