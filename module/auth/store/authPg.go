package store

import (
	"fmt"

	"github.com/wdhafin/eFishery-test/entity"
	"github.com/wdhafin/eFishery-test/internal/db"
	"github.com/wdhafin/eFishery-test/module/auth"
)

type pgRepository struct {
	Conn db.PgDB
}

// NewPgRepository will create an object that represents the auth.Repository interface
func NewPgRepository(Conn db.PgDB) auth.Repository {
	return &pgRepository{Conn}
}

func (pg *pgRepository) Register(eUser entity.User) (res *entity.User, err error) {
	fmt.Println(eUser)
	tx, err := pg.Conn.DB.Begin()

	if err != nil {
		tx.Rollback()
		return &entity.User{}, err
	}

	err = tx.Insert(&eUser)
	if err != nil {
		tx.Rollback()
		return &entity.User{}, err
	}

	tx.Commit()

	return &eUser, err
}

func (pg *pgRepository) CheckUserExist(phone string) (res *entity.User, err error) {

	user := &entity.User{}

	err = pg.Conn.DB.Model(user).Where("phone = ?", phone).Select()

	return user, err
}
