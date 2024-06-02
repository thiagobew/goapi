package tools

import (
	"time"
)

var mockLoginDetails = map[string]*LoginDetails{
	"thiagobew": {
		AuthToken: "123456",
		Username: "thiagobew",
	},
	"testuser": {
		AuthToken:	"654321",
		Username: "testuser",
	},
}

var mockCoinDetails = map[string]*CoinDetails{
	"thiagobew": {
		Coins: 100,
		Username: "thiagobew",
	},
	"testuser": {
		Coins: 50,
		Username: "testuser",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
