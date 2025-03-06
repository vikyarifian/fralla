package models

import "time"

type User struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	Username  string     `gorm:"column:username;unique" json:"username,omitempty" form:"username"`
	Password  string     `gorm:"column:password" json:"password,omitempty" form:"password"`
	FirstName string     `gorm:"column:first_name;" json:"first_name,omitempty" form:"first_name"`
	LastName  string     `gorm:"column:last_name;" json:"last_name,omitempty" form:"last_name"`
	Email     string     `gorm:"column:email;unique" json:"email,omitempty" form:"email"`
	Phone     string     `gorm:"column:phone" json:"phone,omitempty" form:"phone"`
	Level     string     `gorm:"column:level" json:"level,omitempty" form:"level"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE users (
// 	no SERIAL PRIMARY KEY,
// 	id VARCHAR(128) NOT NULL DEFAULT '0',
// 	username VARCHAR(50) UNIQUE NOT NULL DEFAULT ' ',
// 	password VARCHAR(128) UNIQUE NOT NULL DEFAULT ' ',
// 	first_name VARCHAR(50) NOT NULL DEFAULT ' ',
// 	last_name VARCHAR(50) NOT NULL DEFAULT ' ',
// 	email VARCHAR(50) NOT NULL DEFAULT ' ',
// 	phone VARCHAR(25) NOT NULL DEFAULT ' ',
// 	level VARCHAR(10) NOT NULL DEFAULT 'USER',
// 	created_at TIMESTAMP DEFAULT NOW(),
// 	created_by VARCHAR(50) DEFAULT '',
// 	updated_at TIMESTAMP DEFAULT NOW(),
// 	updated_by VARCHAR(50) DEFAULT ''
// );
