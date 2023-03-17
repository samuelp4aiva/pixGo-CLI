package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PixKey struct {
	ID    uint   `gorm:"primaryKey"`
	Key   string `gorm:"unique"`
	Alias string
}

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("pixkeys.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&PixKey{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func addPixKey(db *gorm.DB, key, alias string) error {

	pixKey := PixKey{Key: key, Alias: alias}
	result := db.Create(&pixKey)
	return result.Error
}

func listPixKeys(db *gorm.DB) ([]PixKey, error) {

	var pixKeys []PixKey
	result := db.Find(&pixKeys)

	return pixKeys, result.Error
}
