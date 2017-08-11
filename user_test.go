package main

import (
	"context"
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"os"

	"github.com/goadesign/goa"
	"github.com/JormungandrK/microservice-registration/app"
	"github.com/JormungandrK/microservice-registration/app/test"
)

var (
	service = goa.New("user-test")
	ctrl    = NewUserController(service, &MockMessage{})
)

// Call generated test helper, this checks that the returned media type is of the
// correct type (i.e. uses view "default") and validates the media type.
// Also, it ckecks the returned status code
func TestRegisterUserCreated(t *testing.T) {
	user := &app.UserPayload{
		Fullname:   "fullname",
		Username:   "username",
		Password:   "password",
		Email:      "example@mail.com",
		ExternalID: "qwerc461f9f8eb02aae053f3",
		Roles:      []string{"admin", "user"},
	}

	urlConfig, _ := UrlConfigFromFile("./urlConfig.json")

	gock.New(urlConfig.UserService).
		Post("/users/").
		Reply(201).
		JSON(map[string]interface{}{
			"id":         "59804b3c0000000000000000",
			"fullname":   user.Fullname,
			"username":   user.Username,
			"email":      user.Email,
			"externalId": "qwe04b3c000000qwertydgfsd",
			"roles":      []string{"admin", "user"},
			"active":     false,
		})

	gock.New(urlConfig.UserProfileService).
		Put(fmt.Sprintf("/users/%s/profile", "59804b3c0000000000000000")).
		Reply(204).
		JSON(map[string]interface{}{
			"fullname": user.Fullname,
			"email":    user.Email,
		})

	_, u := test.RegisterUserCreated(t, context.Background(), service, ctrl, user)

	if u == nil {
		t.Fatal("Nil user")
	}
}

// Call generated test helper, this checks that the returned media type is of the
// correct type (i.e. uses view "default") and validates the media type.
// Also, it ckecks the returned status code
func TestRegisterUserBadRequest(t *testing.T) {
	user := &app.UserPayload{
		Fullname:   "fu",
		Username:   "username",
		Password:   "password",
		Email:      "test",
		ExternalID: "qwerc461f9f8eb02aae053f3",
		Roles:      []string{"admin", "user"},
	}

	urlConfig, _ := UrlConfigFromFile("./urlConfig.json")

	gock.New(urlConfig.UserService).
		Post("/users/").
		Reply(400).
		JSON(map[string]interface{}{
			"id":         "59804b3c0000000000000000",
			"fullname":   user.Fullname,
			"username":   user.Username,
			"email":      user.Email,
			"externalId": "qwe04b3c000000qwertydgfsd",
			"roles":      []string{"admin", "user"},
			"active":     false,
		})

	gock.New(urlConfig.UserProfileService).
		Put(fmt.Sprintf("/users/%s/profile", "59804b3c0000000000000000")).
		Reply(400).
		JSON(map[string]interface{}{
			"fullname": user.Fullname,
			"email":    user.Email,
		})

	test.RegisterUserBadRequest(t, context.Background(), service, ctrl, user)
}

func TestEmailConfigFromFile(t *testing.T) {
	file := "./emailConfig.json"
	b, err := exists(file)

	if err != nil {
		t.Fatal()
	}

	if b ==true {
		_, err := EmailConfigFromFile(file)
		if err != nil {
			t.Fail()
		}
	}
}

func TestUrlConfigFromFile(t *testing.T) {
	file := "./urlConfig.json"
	b, err := exists(file)

	if err != nil {
		t.Fatal()
	}

	if b ==true {
		_, err := UrlConfigFromFile(file)
		if err != nil {
			t.Fail()
		}
	}
}

func TestSendEmail(t *testing.T) {
	type email struct {
		ID   string
		Name string
	}

	id := "20"
	username := "testusername"
	userEmail := email{id, username}
	template, errTemp := ParseTemplate("./emailTemplate.html", userEmail)
	
	if errTemp != nil {
		t.Fatal()
	}

	_, err := SendEmail(id, username, userEmail, template)

	if err != nil {
		t.Fail()
	}
}


// Returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}