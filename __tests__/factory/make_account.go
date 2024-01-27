package factory_test

import "kryptify/entity"

func MakeAccount(args ...string) *entity.Account {
	var name, email, password string
	var accountNotVerified bool

	// handle optional arguments
	if len(args) > 0 {
		name = args[0]
	}

	if len(args) > 1 {
		email = args[1]
	}

	if len(args) > 2 {
		password = args[2]
	}

	// handle the "accountNotVerified" argument directly in the NewAccount function
	if len(args) > 3 && args[3] == "false" {
		accountNotVerified = false
	} else {
		accountNotVerified = true
	}

	// default values if optional arguments are empty
	if name == "" {
		name = "Jane Doe"
	}
	if email == "" {
		email = "jane@email.com"
	}
	if password == "" {
		password = "myVerySecurePassword"
	}

	account, err := entity.NewAccount(name, email, password)
	account.IsAccountVerified = accountNotVerified
	if err != nil {
		return nil
	}

	return account
}
