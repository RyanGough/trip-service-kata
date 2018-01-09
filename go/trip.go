package trip

import (
	"errors"
)

type Trip struct{}

type User struct {
	Friends []*User
	Trips   []*Trip
}

func getLoggedUser() (*User, error) {
	return nil, errors.New("getLoggerUser should not be called in unit test")
}

func findTripsByUser(user *User) ([]*Trip, error) {
	return nil, errors.New("findTripsByUser should not be called in unit test")
}

func GetTripsByUser(user *User) ([]*Trip, error) {
	var trips []*Trip
	loggedUser, err := getLoggedUser()
	if err != nil {
		return nil, err
	}
	if loggedUser != nil {
		isFriend := false
		for _, friend := range user.Friends {
			if loggedUser == friend {
				isFriend = true
				break
			}
		}
		if isFriend {
			return findTripsByUser(user)
		} else {
			return trips, nil
		}
	} else {
		return nil, errors.New("User Not Logged In")
	}
}
