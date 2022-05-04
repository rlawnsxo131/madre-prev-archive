package user

import (
	"database/sql"
	"regexp"
	"time"
)

const (
	Key_ID          = "ID"
	Key_Email       = "Email"
	Key_OriginName  = "OriginName"
	Key_DisplayName = "DisplayName"
	Key_PhotoUrl    = "PhotoUrl"
	Key_CreatedAt   = "CreatedAt"
	Key_UpdatedAt   = "UpdatedAt"
)

type User struct {
	ID          string         `json:"id" db:"id"`
	Email       string         `json:"email" db:"email"`
	OriginName  sql.NullString `json:"origin_name" db:"origin_name"`
	DisplayName string         `json:"display_name" db:"display_name"`
	PhotoUrl    sql.NullString `json:"photo_url" db:"photo_url"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

func (u *User) Filter(keys []string) map[string]interface{} {
	result := make(map[string]interface{})

	if keys == nil {
		result["id"] = u.ID
		result["email"] = u.Email
		result["display_name"] = u.DisplayName
		result["created_at"] = u.CreatedAt
		result["updated_at"] = u.UpdatedAt

		if u.OriginName.Valid {
			result["origin_name"] = u.OriginName.String
		} else {
			result["origin_name"] = nil
		}

		if u.PhotoUrl.Valid {
			result["photo_url"] = u.PhotoUrl.String
		} else {
			result["photo_url"] = nil
		}
	} else {
		for _, key := range keys {
			if key == Key_ID {
				result["id"] = u.ID
			} else if key == Key_Email {
				result["email"] = u.Email
			} else if key == Key_OriginName {
				if u.OriginName.Valid {
					result["origin_name"] = u.OriginName.String
				} else {
					result["origin_name"] = nil
				}
			} else if key == Key_DisplayName {
				result["display_name"] = u.DisplayName
			} else if key == Key_PhotoUrl {
				if u.PhotoUrl.Valid {
					result["photo_url"] = u.PhotoUrl.String
				} else {
					result["photo_url"] = nil
				}
			} else if key == Key_CreatedAt {
				result["created_at"] = u.CreatedAt
			} else if key == Key_UpdatedAt {
				result["updated_at"] = u.UpdatedAt
			}
		}
	}

	return result
}

func (u *User) ValidateDisplayName() (bool, error) {
	match, err := regexp.MatchString("^[a-zA-Z0-9가-힣]{1,16}$", u.DisplayName)
	if err != nil {
		return false, err
	}
	return match, nil
}
