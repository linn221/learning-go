package models

func Migrate() {
	DB.AutoMigrate(&Student{})
}
