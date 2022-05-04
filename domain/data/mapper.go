package data

type entityMapper struct{}

func (e entityMapper) toEntity(d *Data) *Data {
	return &Data{
		ID:          d.ID,
		UserID:      d.UserID,
		FileUrl:     d.FileUrl,
		Title:       d.Title,
		Description: d.Description,
		IsPublic:    d.IsPublic,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func (e entityMapper) toModel(d *Data) *Data {
	return &Data{
		UserID:      d.UserID,
		FileUrl:     d.FileUrl,
		Title:       d.Title,
		Description: d.Description,
		IsPublic:    d.IsPublic,
		UpdatedAt:   d.UpdatedAt,
	}
}
