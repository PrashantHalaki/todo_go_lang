// models/todo.go

package models

import (
	"time"
)

type Task struct {
	ID        int       `gorm:"primaryKey:autoIncrement"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Status    int       `json:"status" gorm:"DEFAULT:1"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"DEFAULT:NULL"`
}
