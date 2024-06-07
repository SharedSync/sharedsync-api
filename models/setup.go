package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectionDB() {
    db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/sharedsync_api"))
    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&File{})
    database = db
}

