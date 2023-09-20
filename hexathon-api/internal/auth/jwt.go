package auth

import (
	"errors"
	"strings"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm/clause"
)

func CreateJWTToken(team_name string, role string, jwtKey string) (string, error) {
	claims := jwt.MapClaims{}
	claims["team_name"] = team_name
	claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

func getJWTClaims(tokenString string, jwtKey string) (jwt.MapClaims, error) {
	var err error
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid JWT Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims, nil
	}
	return claims, errors.New("invalid claims")
}

func GetTeamFromJWTToken(tokenString string, jwtKey string) (models.Team, error) {
	authorizationString := strings.Split(tokenString, " ")

	if len(authorizationString) != 2 {
		return models.Team{}, errors.New("invalid Authorization Header")
	}

	token := authorizationString[1]

	if token == "" {
		return models.Team{}, errors.New("not logged in")
	}

	claims, err := getJWTClaims(token, jwtKey)

	var team models.Team

	if err != nil {
		return team, err
	}
	if database.DB.Where("id = ?", claims["team_id"]).First(&team).RowsAffected == 0 {
		return team, errors.New("request made by invalid team")
	}

	database.DB.Where("id = ?", claims["team_id"]).Preload(clause.Associations).First(&team)
	return team, nil
}
