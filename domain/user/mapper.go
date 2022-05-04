package user

type entityMapper struct{}

func (e entityMapper) toEntity(u *User) *User {
	return &User{
		ID:          u.ID,
		Email:       u.Email,
		OriginName:  u.OriginName,
		DisplayName: u.DisplayName,
		PhotoUrl:    u.PhotoUrl,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func (e entityMapper) toModel(u *User) *User {
	return &User{
		Email:       u.Email,
		OriginName:  u.OriginName,
		DisplayName: u.DisplayName,
		PhotoUrl:    u.PhotoUrl,
		UpdatedAt:   u.UpdatedAt,
	}
}
