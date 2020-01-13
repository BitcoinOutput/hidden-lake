package settings

import (
	"../utils"
	"errors"
)

func CheckLifetimeToken(token string) error {
	user := Users[token]
	tokenTime := utils.ParseTime(user.Session.Time)
	currTime := utils.ParseTime(utils.CurrentTime())
	if tokenTime.Add(LIFETIME_TOKEN).Before(currTime) {
		delete(Tokens, Users[token].Hashname)
		delete(Users, token)
		return errors.New("Token lifetime is over")
	}
	return nil
}
