package Character

type CharacterPublic struct {
	UUID           string `json:"uuid"`
	IdentifierName string `json:"identifier_name"`
	ServerUUID     string `json:"server_uuid"`
	ServerName     string `json:"server"`
	Level          int    `json:"level"`
}

func (c Character) Public() CharacterPublic {
	return CharacterPublic{
		c.UUID,
		c.IdentifierName,
		c.ServerUUID,
		c.Server.Name,
		c.Level,
	}
}

func PublicCharacters(characters []Character) []CharacterPublic {
	var publicCharacters []CharacterPublic
	for _, character := range characters {
		publicCharacters = append(publicCharacters, character.Public())
	}
	return publicCharacters
}
