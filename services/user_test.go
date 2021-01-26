package services_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vmkevv/duiztapi/ent"
	"github.com/vmkevv/duiztapi/serverr"
	"github.com/vmkevv/duiztapi/services"
)

type userActions struct {
	users []*ent.User
}

func (u *userActions) Register(name, email string) (*ent.User, error) {
	user := ent.User{
		ID:    len(u.users) + 1,
		Name:  name,
		Email: email,
	}
	u.users = append(u.users, &user)
	return &user, nil
}
func (u userActions) SendEmailToken(email string) error {
	return nil
}
func (u userActions) GenerateToken(ID int) (string, error) {
	return "tokenstring", nil
}
func (u userActions) ExistsEmail(email string) bool {
	for _, u := range u.users {
		if u.Email == email {
			return true
		}
	}
	return false
}
func (u userActions) Login(token string) (*ent.User, error) {
	for _, u := range u.users {
		if u.Email == token {
			return u, nil
		}
	}
	return nil, errors.New("error when login")
}

func TestRegisterService(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: serverr.Handler,
	})
	appV1 := app.Group("/api/v1")
	userActionsMock := &userActions{}
	validator := validator.New()
	services.SetupUserServices(userActionsMock, validator).ServeRoutes(appV1)

	t.Run("It should register correctly", func(t *testing.T) {
		reqStr, _ := json.Marshal(services.RegisterReq{
			Name:  "Kevv",
			Email: "kevv@gmail.com",
		})
		req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewReader(reqStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be OK")

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err, "No errors expected by reading body")

		respData := services.RegisterRes{}

		err = json.Unmarshal(body, &respData)
		assert.Nil(t, err, "No errors expected when unmarshalling resp body")

		assert.NotEqual(t, 0, respData.User.ID, "User ID should not be 0")
		assert.Equal(t, "tokenstring", respData.Token, "Token must coincide")
	})

	t.Run("It should not be possible to register with the same email", func(t *testing.T) {
		reqStr, _ := json.Marshal(services.RegisterReq{
			Name:  "Kevv",
			Email: "kevv@gmail.com",
		})
		req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewReader(reqStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "It should return a bad request error")
	})

	t.Run("It should return error if name is not valid", func(t *testing.T) {
		reqStr, _ := json.Marshal(services.RegisterReq{
			Name:  "K",
			Email: "kevv2@gmail.com",
		})
		req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewReader(reqStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "It should return a bad request error")
	})

	t.Run("It should return error if email is not valid", func(t *testing.T) {
		reqStr, _ := json.Marshal(services.RegisterReq{
			Name:  "Kevin",
			Email: "kevvgmail.com",
		})
		req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewReader(reqStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "It should return a bad request error")
	})
}

func TestSendEmail(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: serverr.Handler,
	})
	appV1 := app.Group("/api/v1")
	userActionsMock := &userActions{}
	userActionsMock.Register("Kevv", "kevv@gmail.com")
	validator := validator.New()
	services.SetupUserServices(userActionsMock, validator).ServeRoutes(appV1)

	t.Run("It shoud handle request without error for existing users", func(t *testing.T) {
		reqStr, _ := json.Marshal(services.SendEmailReq{
			Email: "kevv@gmail.com",
		})
		req, _ := http.NewRequest("POST", "/api/v1/email", bytes.NewReader(reqStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "It should return OK status")

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err, "No errors expecten by reading body")

		respData := services.SendEmailRes{}
		err = json.Unmarshal(body, &respData)
		assert.Nil(t, err, "No errors expected when unmarshalling resp body")

		assert.NotEqual(t, "", respData.Message, "Expected non empty response")
	})

	t.Run("It should return 400 error for unknown email", func(t *testing.T) {
		reqStr, _ := json.Marshal(services.SendEmailReq{"unknown@gmail.com"})

		req, _ := http.NewRequest("POST", "/api/v1/email", bytes.NewReader(reqStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "reponse status code should be a 400 error")
	})
}

func TestLogin(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: serverr.Handler,
	})
	appV1 := app.Group("/api/v1")
	userActionsMock := &userActions{}
	userActionsMock.Register("Kevv", "kevv@gmail.com")
	validator := validator.New()
	services.SetupUserServices(userActionsMock, validator).ServeRoutes(appV1)

	t.Run("It should verify a correct token string", func(t *testing.T) {
		reqStr, _ := json.Marshal(services.LoginReq{"kevv@gmail.com"})

		req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewReader(reqStr))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode, "response code should be 200")
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err, "No errors expected by reading response body")
		respData := services.LoginRes{}
		err = json.Unmarshal(body, &respData)
		assert.Nil(t, err, "No errors expected by unmarshalling response")
	})

}
