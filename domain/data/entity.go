package data

import "time"

const (
	Key_ID          = "ID"
	Key_UserID      = "UserID"
	Key_FileUrl     = "FileUrl"
	Key_Title       = "Title"
	Key_Description = "Description"
	Key_IsPublic    = "IsPublic"
	Key_CreatedAt   = "CreatedAt"
	Key_UpdatedAt   = "UpdatedAt"
)

type Data struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	FileUrl     string    `json:"file_url" db:"file_url"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	IsPublic    bool      `json:"is_public" db:"is_public"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (d *Data) Filter(keys []string) map[string]interface{} {
	result := make(map[string]interface{})

	if keys == nil {
		result["id"] = d.ID
		result["user_id"] = d.UserID
		result["file_url"] = d.FileUrl
		result["title"] = d.Title
		result["description"] = d.Description
		result["is_public"] = d.IsPublic
		result["created_at"] = d.CreatedAt
		result["updated_at"] = d.UpdatedAt
	} else {
		for _, key := range keys {
			if key == Key_ID {
				result["id"] = d.ID
			} else if key == Key_UserID {
				result["user_id"] = d.UserID
			} else if key == Key_FileUrl {
				result["file_url"] = d.FileUrl
			} else if key == Key_Title {
				result["title"] = d.Title
			} else if key == Key_Description {
				result["description"] = d.Description
			} else if key == Key_IsPublic {
				result["is_public"] = d.IsPublic
			} else if key == Key_CreatedAt {
				result["created_at"] = d.CreatedAt
			} else if key == Key_UpdatedAt {
				result["updated_at"] = d.UpdatedAt
			}
		}
	}

	return result
}
