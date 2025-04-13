package main

import (
	"fmt"
	. "ginprac/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func addRecord(db *gorm.DB) {
	//
	// 添加老师对象
	t1 := Teacher{User: User{BaseModel: BaseModel{Name: "张珺"}, Pwd: "123456"}}
	cl := []Class{
		{BaseModel: BaseModel{Name: "2017级化工班"}, TutorID: 1, StuNum: 0},
		{BaseModel: BaseModel{Name: "2017级电子班"}, TutorID: 1, StuNum: 0},
		{BaseModel: BaseModel{Name: "2017级软件班"}, TutorID: 1, StuNum: 0},
		{BaseModel: BaseModel{Name: "2017级土木班"}, TutorID: 1, StuNum: 0},
	}
	db.Create(&t1)
	db.Create(&cl)
	for i, v := range cl {
		fmt.Println(i, v)
	}
}

func SelectRecord(db *gorm.DB) {
	cl := []Class{}
	db.Find(&cl)
	for i, v := range cl {
		fmt.Println(i, v.Tutor)
	}
	var count int64
	db.Model(Class{}).Where("tutor_id between ? and ?", 0, 1).Count(&count)
	fmt.Println(count)
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			// SlowThreshold: time.Second,
			LogLevel: logger.Info,
		})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库连接失败")
	}
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Student{})
	// addRecord(db)
	SelectRecord(db)
}
