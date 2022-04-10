package user

type UserFormater struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Occuption string `json:"occuption"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func FormatUser(user User, token string) UserFormater {
	Formater := UserFormater{
		ID:        user.ID,
		Name:      user.Name,
		Occuption: user.Occuption,
		Email:     user.Email,
		Token:     token,
	}
	return Formater
}
