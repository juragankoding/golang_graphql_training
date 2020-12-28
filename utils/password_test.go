package utils

import "testing"

func TestPassword(t *testing.T) {
	password := "testing1"
	passwordEncrpte, err := HashPassword(password)

	if err != nil {
		t.Error(err.Error())
	}

	if passwordEncrpte == "" {
		t.Error("passsword is nil")
	}

	t.Logf("password generate is : %s", passwordEncrpte)

	passwordDecrypted, err := Decrypt(passwordEncrpte)

	if err != nil {
		t.Error(err.Error())
	}

	if passwordDecrypted == "" {
		t.Error("passsword is nil")
	}

	t.Logf("password decrypt is : %s", passwordDecrypted)

	if passwordDecrypted != password {
		panic("pasword is not same")
	}

}
