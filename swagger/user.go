package swagger

type user struct {
	Id     	string `json:"id" example:"1"`
	Name    string `json:"name" example:"name"`
	Email    string `json:"email" example:"email@email.com"`
}

type createdUser struct {
	Token string `json:"token" example:"1|V96CSH7oe3Pjr174sL4s24rCcnjZhWWw0IBUUjeh"`
	User 	user `json:"user"`
}

type ResponseCreatedUser struct {
	Data createdUser
}

type ResponseGetUser struct {
	Data user
}