package jwt_test

import (
	"testing"

	"github.com/ivansevryukov1995/url-shortening-service/pkg/jwt"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJwt("7y7c8bo5cRl60XVvR4PrJw/8DPA/fyqy+2CuPxGvPAg=")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
		return
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
		return
	}

}
