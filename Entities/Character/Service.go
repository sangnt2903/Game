package Character

import (
	"MiniGameAPI/Commons/Paging"
)

type CharacterServiceInterface interface {
	CreateCharacter(Character *Character) error
	GetCharacters(Characters *[]Character, paginator ...*Paging.Paginator) error
	GetCharacter(Character *Character, uuid string) error
}

type CharacterService struct {
	sr *CharacterRepository
}

func (ss *CharacterService) GetCharacter(Character *Character, uuid string) error {
	return ss.sr.GetInstanceByID(Character, uuid)
}

func (ss *CharacterService) GetCharacters(Characters *[]Character, paginator ...*Paging.Paginator) error {
	if len(paginator) > 0 {
		return paginator[0].Paging(ss.sr.Db, Characters)
	}
	return ss.sr.GetInstancesWithConditions(Characters, "")
}

func (ss *CharacterService) CreateCharacter(Character *Character) error {
	return ss.sr.CreateInstance(Character)
}

var _ CharacterServiceInterface = &CharacterService{}

func NewCharacterService(base *CharacterRepository) *CharacterService {
	return &CharacterService{base}
}
