package usecase

import (
	"errors"
	"time"

	"github.com/spf13/viper"
	"github.com/thanhpk/randstr"
	"github.com/wdhafin/eFishery-test/entity"
	"github.com/wdhafin/eFishery-test/module/auth"
	"github.com/wdhafin/eFishery-test/pkg/helper"
	"github.com/wdhafin/eFishery-test/pkg/utils"
	"github.com/wdhafin/eFishery-test/schema"
)

// AuthUsecase will create a usecase with its required repo
type AuthUsecase struct {
	authRepo auth.Repository
}

// NewAuthUsecase will create new an contactUsecase object representation of auth.Usecase
func NewAuthUsecase(ar auth.Repository) AuthUsecase {
	return AuthUsecase{
		authRepo: ar,
	}
}

// Register returns a user data
func (u *AuthUsecase) Register(eUser entity.User) (*entity.User, error) {

	user, err := u.authRepo.CheckUserExist(eUser.Phone)

	if user.ID == 0 {
		eUser.Password = randstr.String(4)
		eUser.Timestamp = time.Now().Unix()

		user, err = u.authRepo.Register(eUser)
		if err != nil {
			return &entity.User{}, err
		}
	} else {
		err = errors.New("user already exist")
	}

	return user, err
}

// Login returns a user data
func (u *AuthUsecase) Login(eUser entity.User) (*schema.Token, error) {

	user, err := u.authRepo.CheckUserExist(eUser.Phone)
	if err != nil {
		return &schema.Token{}, errors.New("User not found")
	}

	valid := utils.CheckPassword(user.Password, eUser.Password)
	if valid != true {
		return &schema.Token{}, errors.New("Wrong Password")
	}

	expAccess := viper.GetDuration("auth.accessExpiry")
	expRefresh := viper.GetDuration("auth.refreshExpiry")

	expJWTAccess := time.Now().Add(expAccess)
	expJWTRefresh := time.Now().Add(expRefresh)
	token, err := helper.GenerateToken(*user, expJWTAccess, expJWTRefresh)
	if err != nil {
		return &schema.Token{}, err
	}

	return &token, nil
}
