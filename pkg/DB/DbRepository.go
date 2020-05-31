package DB

import "GolangTest/model"

type DbInterface interface {
	UpdatetUser(model.User)
	InsertUser(model.User)
	GetUserByFields(map[string]string) []*model.User
	CreateDB()
	FillDB()
}

func GetDb() *MongoDB {
	return new(MongoDB)
}
