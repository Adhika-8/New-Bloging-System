package app

import (
	table "new-bloging-system/model/entities"
)

func DBmigrate() {
	err := db.AutoMigrate(&table.CreateBlogs{})
	if err != nil {
		panic("Database Migration failed")
	}

}
