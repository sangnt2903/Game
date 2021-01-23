package Database

import (
	"MiniGameAPI/Entities/Character"
	"MiniGameAPI/Entities/Item"
	"MiniGameAPI/Entities/Server"
	"MiniGameAPI/Entities/Specification"
	"MiniGameAPI/Logging/Error"
)

func AutoMigrate() {
	db := GetDatabase()
	if err := db.AutoMigrate(
		&Server.Server{},
		&Character.Character{},
		&Item.Item{},
		&Specification.Specification{},
	); Error.ErrorService.HasError(err) {
		panic(err)
	}
}
