package common

import "database/sql"

func IsExistEntity(id string, err error) (bool, error) {
	exist := false

	if err != nil {
		if err == sql.ErrNoRows {
			return exist, nil
		} else {
			return exist, err
		}
	}

	if id != "" {
		exist = true
	}

	return exist, nil
}
