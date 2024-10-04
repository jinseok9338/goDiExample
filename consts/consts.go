package consts

import "gorm.io/gorm"

// Define an interface
type DbHandlerInterface interface {
	GetDB() *gorm.DB
}

// Modify DbHandlers to implement this interface
type DbHandlers struct {
	DB *gorm.DB
}

func (d DbHandlers) GetDB() *gorm.DB {
	return d.DB
}

func ProvideDB(DB *gorm.DB) DbHandlers {
	return DbHandlers{DB: DB}
}
