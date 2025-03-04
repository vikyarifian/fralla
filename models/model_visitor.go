package models

import "time"

type Visitor struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	IP        string     `gorm:"column:ip" json:"ip,omitempty" form:"ip"`
	Token     string     `gorm:"column:token" json:"auth_token"`
	IsAuth    bool       `gorm:"is_auth" json:"is_auth"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE visitors (
// 	no SERIAL PRIMARY KEY,
// 	id VARCHAR(128) NOT NULL DEFAULT '0',
// 	token VARCHAR(256) NOT NULL DEFAULT ' ',
// 	is_auth BOOLEAN DEFAULT FALSE,
// 	created_at TIMESTAMP DEFAULT NOW(),
// 	created_by VARCHAR(50) DEFAULT '',
// 	updated_at TIMESTAMP DEFAULT NOW(),
// 	updated_by VARCHAR(50) DEFAULT ''
// );
