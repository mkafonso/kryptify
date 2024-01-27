package factory_test

import "kryptify/entity"

func MakeCredential(args ...string) *entity.Credential {
	var email, password, website, owner_id string

	// handle optional arguments
	if len(args) > 0 {
		email = args[0]
	}

	if len(args) > 1 {
		password = args[1]
	}

	if len(args) > 2 {
		website = args[2]
	}

	if len(args) > 3 {
		owner_id = args[3]
	}

	// default values if optional arguments are empty
	if email == "" {
		email = "john@email.com"
	}
	if password == "" {
		password = "test1234"
	}
	if website == "" {
		website = "https://my-website.com"
	}
	if owner_id == "" {
		owner_id = "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5"
	}

	credential, err := entity.NewCredential(email, password, website, owner_id)
	if err != nil {
		return nil
	}

	return credential
}
