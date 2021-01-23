package Server

type ServerPublic struct {
	UUID string `json:"uuid"`
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"`
}

func (s Server) Public() ServerPublic {
	return ServerPublic{
		UUID: s.UUID,
		Host: s.Host,
		Port: s.Port,
		Name: s.Name,
	}
}

func PublicServers(servers []Server) []ServerPublic {
	var publicServers []ServerPublic
	for _, server := range servers {
		publicServers = append(publicServers, server.Public())
	}
	return publicServers
}
