package user

type CreateUserParams struct {
	UUID        string `json:"uuid" db:"uuid"`
	Email       string `json:"email" db:"email"`
	Username    string `json:"username" db:"username"`
	DisplayName string `json:"display_name" db:"display_name"`
	PhotoUrl    string `json:"photo_url" db:"photo_url"`
}
