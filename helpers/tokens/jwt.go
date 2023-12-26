package tokens

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(secret string, idUser uint) (string, error) {
	if secret == "" {
		return "", errors.New("invalid token secret")
	}

	if idUser == 0 {
		return "", errors.New("invalid user id")
	}

	var claim = jwt.MapClaims{}
	claim["id"] = idUser
	claim["iat"] = time.Now().UnixMilli()
	claim["exp"] = time.Now().Add(time.Hour * 2).UnixMilli()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	strToken, err := token.SignedString([]byte("altamantul"))
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func ExtractToken(secret string, t *jwt.Token) (uint, error) {
	if secret == "" {
		return 0, errors.New("invalid token secret")
	}

	if t == nil {
		return 0, errors.New("invalid jwt token")
	}

	var userID uint
	expiredTime, err := t.Claims.GetExpirationTime()
	if err != nil {
		return 0, err
	}

	var eTime = *expiredTime

	if t.Valid && eTime.Compare(time.Now()) > 0 {
		var tokenClaims = t.Claims.(jwt.MapClaims)

		if tokenClaims["id"] == nil {
			return 0, errors.New("id user tidak ditemukan")
		}

		userID = uint(tokenClaims["id"].(float64))

		return userID, nil
	}

	return 0, errors.New("token tidak valid")
}
