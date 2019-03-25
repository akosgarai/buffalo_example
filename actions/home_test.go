package actions

import "github.com/akosgarai/buffalo_example/models"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	// Should redirect to /login page
	as.Equal(302, res.Code)
	as.Contains(res.Body.String(), "<a href=\"/login\">Found</a>")
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	u := &models.Administrator{
		Email:           "test-admin@admin.com",
		Pwd:             "admin",
		PwdConfirmation: "admin",
		Name:            "test admin",
		Username:        "testAdmin",
		Password:        "$2a$10$zm/lRrBz8kObRwHcF1erHOc6Ac2o7Cog.KP8fVen388EtvNaASJHW",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Set("current_admin_id", u.ID)

	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "logout")

	as.Session.Clear()
	// Should redirect to /login page
	res = as.HTML("/").Get()
	as.Equal(302, res.Code)
	as.Contains(res.Body.String(), "<a href=\"/login\">Found</a>")
}
