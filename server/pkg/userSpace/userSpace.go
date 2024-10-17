package userSpace

import (
	"fmt"
	"os/user"
	"runtime"
)

type UserData struct {
	Username string
	UserID   string
	Name     string
	HomeDir  string
	OS       string
}

func GetUserData() (*UserData, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("error getting current user: %v", err)
	}

	userData := &UserData{
		Username: currentUser.Username,
		UserID:   currentUser.Uid,
		Name:     currentUser.Name,
		HomeDir:  currentUser.HomeDir,
		OS:       runtime.GOOS,
	}

	return userData, nil
}

func (u *UserData) PrintUserData() {
	fmt.Printf("Username: %s\n", u.Username)
	fmt.Printf("User ID: %s\n", u.UserID)
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Home Directory: %s\n", u.HomeDir)
	fmt.Printf("Operating System: %s\n", u.OS)
}
