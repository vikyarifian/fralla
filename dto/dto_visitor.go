package dto

type Visitor struct {
	No        int    `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	Username  string `gorm:"column:username;unique" json:"username,omitempty" form:"username"`
	FirstName string `gorm:"column:first_name;" json:"first_name,omitempty" form:"first_name"`
	LastName  string `gorm:"column:last_name;" json:"last_name,omitempty" form:"last_name"`
	Email     string `gorm:"column:email;unique" json:"email,omitempty" form:"email"`
	Phone     string `gorm:"column:phone;unique" json:"phone,omitempty" form:"phone"`
}
