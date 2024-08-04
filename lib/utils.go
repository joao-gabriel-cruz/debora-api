package lib

import (
	"github.com/joao-gabriel-cruz/debora-api/prisma/db"
)

func ErrorData[T *db.UserModel | []db.UserModel](data T, err error) (T, error) {

	if err != nil {
		return data, err
	}

	if data != nil {
		return data, nil
	}

	return data, nil
}

func Error(err error) error {
	if err != nil {
		return err
	}

	return nil
}
