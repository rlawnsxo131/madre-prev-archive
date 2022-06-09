package user

type userEntityMapper struct{}

func (e userEntityMapper) toEntity(u *User) *User {
	return &User{
		ID:         u.ID,
		Email:      u.Email,
		OriginName: u.OriginName,
		Username:   u.Username,
		PhotoUrl:   u.PhotoUrl,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (e userEntityMapper) toModel(u *User) *User {
	return &User{
		Email:      u.Email,
		OriginName: u.OriginName,
		Username:   u.Username,
		PhotoUrl:   u.PhotoUrl,
		UpdatedAt:  u.UpdatedAt,
	}
}
