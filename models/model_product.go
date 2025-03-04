package models

import "time"

type Product struct {
	No          int              `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID          string           `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	Title       string           `gorm:"column:title" json:"title,omitempty" form:"title"`
	Description string           `gorm:"column:description" json:"description,omitempty" form:"description"`
	Photos      []ProductPhoto   `gorm:"foreignKey:ID;references:ProductID" json:"photos,omitempty"`
	Variants    []ProductVariant `gorm:"foreignKey:ID;references:ProductID" json:"variants,omitempty"`
	CategoryID  string           `gorm:"column:category_id" json:"-" form:"category_id"`
	Category    Category         `gorm:"foreignKey:CategoryID;references:ID" json:"category,omitempty"`
	GenderID    string           `gorm:"column:gender_id" json:"-" form:"gender_id"`
	Gender      Gender           `gorm:"foreignKey:GenderID;references:ID" json:"gender,omitempty"`
	Condition   int              `gorm:"column:condition" json:"condition,omitempty" form:"condition"`
	CurrencyID  string           `gorm:"column:currency_id" json:"-" form:"currency_id"`
	Currency    Currency         `gorm:"foreignKey:CurrencyID;references:ID" json:"currency,omitempty"`
	Status      string           `gorm:"column:status" json:"status,omitempty" form:"status"`
	CreatedAt   *time.Time       `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty" form:"created_at"`
	CreatedBy   string           `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt   *time.Time       `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy   string           `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE products (
// 	no SERIAL PRIMARY KEY NOT NULL,
//    id VARCHAR(128) NOT NULL DEFAULT '0',
//    title VARCHAR(75) NOT NULL DEFAULT ' ',
//    description TEXT NOT NULL ,
//    category_id VARCHAR(128) NOT NULL DEFAULT ' ',
//    gender_id VARCHAR(128) NOT NULL DEFAULT ' ',
//    currency_id VARCHAR(128) NOT NULL DEFAULT ' ',
//    condition INT NOT NULL DEFAULT 1,
//    status VARCHAR(5) NOT NULL DEFAULT 'A',
//    created_at TIMESTAMP DEFAULT NOW(),
//    created_by VARCHAR(50) DEFAULT '',
//    updated_at TIMESTAMP DEFAULT NOW(),
//    updated_by VARCHAR(50) DEFAULT ''
// );

type ProductVariant struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	ProductID string     `gorm:"column:product_id;" json:"-" form:"product_id"`
	ColorID   string     `gorm:"column:color_id" json:"-" form:"color_id"`
	Color     Color      `gorm:"foreignKey:ColorID;references:ID" json:"color,omitempty"`
	SizeID    string     `gorm:"column:size_id" json:"-" form:"size_id"`
	Size      Color      `gorm:"foreignKey:SizeID;references:ID" json:"size,omitempty"`
	Price     float64    `gorm:"column:price" json:"price,omitempty" form:"price"`
	Stock     float64    `gorm:"column:stock" json:"stock,omitempty" form:"stock"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE product_variants (
// 	no SERIAL PRIMARY KEY NOT NULL,
//    id VARCHAR(128) NOT NULL DEFAULT '0',
//    product_id VARCHAR(128) NOT NULL DEFAULT ' ',
//    color_id VARCHAR(128) NOT NULL DEFAULT ' ',
//    size_id VARCHAR(128) NOT NULL DEFAULT ' ',
//    price NUMERIC(12,6) NOT NULL DEFAULT 0,
//    stock NUMERIC(12,0) NOT NULL DEFAULT 0,
//    created_at TIMESTAMP DEFAULT NOW(),
//    created_by VARCHAR(50) DEFAULT '',
//    updated_at TIMESTAMP DEFAULT NOW(),
//    updated_by VARCHAR(50) DEFAULT ''
// );

type ProductPhoto struct {
	No        int        `gorm:"column:no;primaryKey" json:"-,omitempty" form:"-"`
	ID        string     `gorm:"column:id;unique" json:"id,omitempty" form:"id"`
	ProductID string     `gorm:"column:product_id;" json:"-" form:"product_id"`
	Name      string     `gorm:"column:name;" json:"name" form:"name"`
	Photo     string     `gorm:"column:photo" json:"photo,omitempty" form:"photo"`
	CreatedAt *time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty" form:"created_at"`
	CreatedBy string     `gorm:"column:created_by" json:"created_by,omitempty" form:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP" json:"updated_at,omitempty" form:"updated_at"`
	UpdatedBy string     `gorm:"column:updated_by" json:"updated_by,omitempty" form:"updated_by"`
}

// CREATE TABLE product_photos (
// 	no SERIAL PRIMARY KEY NOT NULL,
//    id VARCHAR(128) NOT NULL DEFAULT '0',
//    product_id VARCHAR(128) NOT NULL DEFAULT ' ',
//    name VARCHAR(50) NOT NULL DEFAULT ' ',
//    photo VARCHAR(128) NOT NULL DEFAULT ' ',
//    created_at TIMESTAMP DEFAULT NOW(),
//    created_by VARCHAR(50) DEFAULT '',
//    updated_at TIMESTAMP DEFAULT NOW(),
//    updated_by VARCHAR(50) DEFAULT ''
// );
