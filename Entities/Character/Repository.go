package Character

import (
	"MiniGameAPI/Entities/Base"
)

type CharacterRepository struct {
	*Base.RepositoryBase
}

func NewCharacterRepository(base *Base.RepositoryBase) *CharacterRepository {
	return &CharacterRepository{base}
}
