package base

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func InitDirector(firstName string, lastName string) Director{
	return Director{FirstName: firstName, LastName: lastName}
}