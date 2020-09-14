package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/wdhafin/eFishery-test/entity"
	"github.com/wdhafin/eFishery-test/schema"
)

// ApplicationName is for JWT Application Name
const ApplicationName = "eF"

// JWTSigningMethod is JWT's signing method
var jwtSigningMethod = jwt.SigningMethodHS256

// GenerateToken will generate both access and refresh token
// for current user.
// Access Token will be expired in 15 Minutes
// Refresh Token will be expired in 6 Months
func GenerateToken(user entity.User, expJWTAccess time.Time, expJWTRefresh time.Time) (token schema.Token, e error) {
	jwtToken, e := GenerateJWT(user, expJWTAccess)
	if e != nil {
		return
	}

	refreshToken, e := generateRefresh(user, expJWTRefresh)
	if e != nil {
		return
	}

	token = schema.Token{
		Access:  jwtToken,
		Refresh: refreshToken,
		TTL:     expJWTAccess.Format(time.UnixDate),
	}

	return
}

// GenerateJWT is
func GenerateJWT(user entity.User, expiration time.Time) (tokenString string, e error) {
	claims := schema.JWTMyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: expiration.Unix(),
		},
		Name:      user.Name,
		Phone:     user.Phone,
		Role:      user.Role,
		Timestamp: user.Timestamp,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString([]byte(viper.GetString("auth.accessSecret")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

func generateRefresh(user entity.User, expiration time.Time) (tokenString string, e error) {
	claims := schema.JWTMyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: expiration.Unix(),
		},
		Name:      user.Name,
		Phone:     user.Phone,
		Role:      user.Role,
		Timestamp: user.Timestamp,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := token.SignedString([]byte(viper.GetString("auth.refreshSecret")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

// GetAuthenticatedUser is
func GetAuthenticatedUser(token *jwt.Token) schema.JWTMyClaims {

	claims := token.Claims.(jwt.MapClaims)

	jClaim := schema.JWTMyClaims{
		Name:      claims["name"].(string),
		Phone:     claims["phone"].(string),
		Role:      claims["role"].(string),
		Timestamp: int64(claims["timestamp"].(float64)),
	}

	return jClaim
}
