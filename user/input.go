package user

type RegisterUserInput struct {
	Name      string `json:"name" binding:"required"`
	Occuption string `json:"occuption" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//input dari user
//handler, mapping input dari user -> struct input
//service : melakukan mapping dari struct input ke struct entity
//repository
//db
