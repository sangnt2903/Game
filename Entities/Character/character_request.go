package Character

type CharacterRequest struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}
