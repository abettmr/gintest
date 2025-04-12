package model

import (
	"time"
)

type BaseModel struct {
	Id         int64     `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(32);not null"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoCreateTime"`
}

type User struct {
	BaseModel
	Pwd    string `gorm:"type:varchar(255);not null"`
	Remark string `gorm:"type:varchar(255)"`
}

type Teacher struct {
	User
}

type Student struct {
	User
	// 多对一
	ClassID int
	Class   Class `gorm:"foreignKey:ClassID"`
	// 多对多
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}

type Course struct {
	BaseModel
}

type Class struct {
	BaseModel
	StuNum  int
	TutorID int
	Tutor   Teacher `gorm:"foreignKey:TutorID;constraint:OnDelete:CASCADE;"`
}
