package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type File struct {
    UID         string      
    ID          int64       
    Name        string      
    Size        string      
    Owner       string      
    Visibility  string      
    Description string      
    CreatedAt   time.Time   
    UpdatedAt   time.Time   
    DeletedAt   *time.Time  
}

func (f *File) BeforeCreate(t *gorm.DB) (err error) {
    f.UID = uuid.New().String()
    return
}

