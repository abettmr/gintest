package model
import ("time")

type BaseModel struct {
    Id int64 `gorm:"primaryKey"`
	CreateTime time.Time `gorm: "autoCreateTime"`
	DelTime time.Time `gorm: "autoCreateTime"`
}

type User struct {
	BaseModel
	Name string `gorm:"type:varchar(32);not null"`
	Pwd string `gorm:"type:varchar(255);not null"`
	Remark string `gorm:"type:varchar(255)"`
}

type Teacher struct {
	User
}

type Student struct{
	User
	ClassID int
	
	Class Class `gorm:""`
}

type Course struct{
	
}
type Class struct{
	StuNum int 
	Tutor Teacher `gorm:"constraint:OnDelete:CASCADE"`
}
