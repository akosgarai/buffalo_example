package actions

import (
	"github.com/akosgarai/buffalo_example/models"
)

func getTestAdmin() *models.Administrator {
	return &models.Administrator{
		Email:           "test-admin@admin.com",
		Password:        "admin",
		Pwd:             "admin",
		PwdConfirmation: "admin",
		Name:            "test admin",
		Username:        "testAdmin",
	}
}
func getTestAdminWithBadPassword() *models.Administrator {
	return &models.Administrator{
		Email:    "test-admin@admin.com",
		Password: "bad",
	}
}
func (as *ActionSuite) Test_Auth_New() {
	res := as.HTML("/login").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
	as.Contains(res.Body.String(), "Sign In!")
	as.Contains(res.Body.String(), "Email")
	as.Contains(res.Body.String(), "Pwd")
}

func (as *ActionSuite) Test_Auth_Login() {
	u := getTestAdmin()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/login").Post(u)
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Auth_Login_Redirect() {
	as.Session.Clear()
	u := getTestAdmin()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	as.Session.Set("redirectURL", "/")

	res := as.HTML("/login").Post(u)
	as.Equal(302, res.Code)
	as.Equal(res.Location(), "/")
}

func (as *ActionSuite) Test_Auth_Login_UnknownUser() {
	as.Session.Clear()
	u := getTestAdmin()

	res := as.HTML("/login").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Login_BadPassword() {
	u := getTestAdmin()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Clear()

	wrongAdmin := getTestAdminWithBadPassword()
	res := as.HTML("/login").Post(wrongAdmin)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}
