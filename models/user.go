package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Age      int    `json:"age"`
}

// Name The Table
func (User) TableName() string {
	return "users"
}
