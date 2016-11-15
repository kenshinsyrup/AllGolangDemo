package main

import (
	"github.com/jinzhu/gorm"
)

type Data struct {
	ID      int    `gorm:"primary_key"`
	Content string `gorm:"content"`
}

type Manager struct {
	DB *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{DB: db}
}

func (m *Manager) insert(data *Data) error {
	return m.DB.Create(data).Error
}
