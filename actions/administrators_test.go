package actions

func (as *ActionSuite) Test_AdministratorsResource_List_Without_Logged_In_User() {
	res := as.HTML("/administrators/").Get()
	as.Equal(302, res.Code)
	as.Equal("/login", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_List() {
	u := getTestAdmin()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/login").Post(u)

	res = as.HTML("/administrators/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Administrators")
}

func (as *ActionSuite) Test_AdministratorsResource_New_Without_Logged_In_User() {
	res := as.HTML("/administrators/new/").Get()
	as.Equal(302, res.Code)
	as.Equal("/login", res.Location())
}

func (as *ActionSuite) Test_AdministratorsResource_New_Without_Privileges() {
	u := getTestAdmin()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/login").Post(u)

	res = as.HTML("/administrators/new/").Get()
	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}

func (as *ActionSuite) Test_AdministratorsResource_New_With_All_Privileges() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPrivs()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/login").Post(u)

	res = as.HTML("/administrators/new/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "New Administrator")
	as.Contains(res.Body.String(), "Setup privileges")
}

func (as *ActionSuite) Test_AdministratorsResource_Show_With_All_Privileges() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPrivs()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/" + u.ID.String() + "/").Get()
	as.Equal(200, res.Code)

	// Page header
	as.Contains(res.Body.String(), "Administrator#Show")
	// Back button
	as.Contains(res.Body.String(), "Back to all Administrators")
	// Edit button (due to all priv.)
	as.Contains(res.Body.String(), "Edit")
	// Delete button
	as.Contains(res.Body.String(), "Destroy")
	// Name label
	as.Contains(res.Body.String(), "<strong>Name</strong>")
	// Admin name
	as.Contains(res.Body.String(), u.Name)
	// Username label
	as.Contains(res.Body.String(), "<strong>Username</strong>")
	// Admin username
	as.Contains(res.Body.String(), u.Username)
	// Email label
	as.Contains(res.Body.String(), "<strong>Email</strong>")
	// Admin email
	as.Contains(res.Body.String(), u.Email)
	// Privs label
	as.Contains(res.Body.String(), "<strong>Privs</strong>:")
}
func (as *ActionSuite) Test_AdministratorsResource_Show_Without_Admin_Privileges() {
	as.LoadFixture("insert privileges")
	u := getTestAdmin()

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/" + u.ID.String() + "/").Get()
	as.Equal(200, res.Code)

	// Page header
	as.Contains(res.Body.String(), "Administrator#Show")
	// Back button
	as.Contains(res.Body.String(), "Back to all Administrators")
	// Edit button
	as.NotContains(res.Body.String(), "Edit")
	// Delete button
	as.NotContains(res.Body.String(), "Destroy")
	// Name label
	as.Contains(res.Body.String(), "<strong>Name</strong>")
	// Admin name
	as.Contains(res.Body.String(), u.Name)
	// Username label
	as.Contains(res.Body.String(), "<strong>Username</strong>")
	// Admin username
	as.Contains(res.Body.String(), u.Username)
	// Email label
	as.Contains(res.Body.String(), "<strong>Email</strong>")
	// Admin email
	as.Contains(res.Body.String(), u.Email)
	// Privs label
	as.Contains(res.Body.String(), "<strong>Privs</strong>:")
}

func (as *ActionSuite) Test_AdministratorsResource_Create_Without_Logged_In_User() {
	res := as.HTML("/administrators/").Post(nil)
	as.Equal(302, res.Code)
	as.Equal("/login", res.Location())
}

func (as *ActionSuite) Test_AdministratorsResource_Create_Without_Privilege() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/").Post(nil)
	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}

func (as *ActionSuite) Test_AdministratorsResource_Create_With_Priv_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "Privileges")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/").Post(nil)
	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}

func (as *ActionSuite) Test_AdministratorsResource_Create_With_Admin_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "Users")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/").Post(nil)
	as.Equal(422, res.Code)
}

func (as *ActionSuite) Test_AdministratorsResource_Create_With_All_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/").Post(nil)
	as.Equal(500, res.Code)
}

func (as *ActionSuite) Test_AdministratorsResource_Create_With_Admin_Priv_Priv_Data() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")
	newAdmin := getTestAdminWithPriv("new-test-admin@admin.com", "")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/").Post(newAdmin)
	as.Equal(302, res.Code)
	as.Contains(res.Location(), "/administrators/")
}
func (as *ActionSuite) Test_AdministratorsResource_Create_Email_Duplication() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	other := getTestAdminWithPriv("other-test-admin@admin.com", "All")
	verrs, err = other.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/").Post(other)
	as.Equal(500, res.Code)
}

func (as *ActionSuite) Test_AdministratorsResource_Edit_Without_Logged_In_User() {
	res := as.HTML("/administrators/invalid-id/edit/").Get()
	as.Equal(302, res.Code)
	as.Equal("/login", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Edit_Without_Privilege() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "Privileges")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/invalid-id/edit/").Get()

	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Edit_With_Priv_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "Privileges")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/invalid-id/edit/").Get()

	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Edit_With_Admin_Priv_Fake_User() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/invalid-id/edit/").Get()

	as.Equal(404, res.Code)
}

func (as *ActionSuite) Test_AdministratorsResource_Edit_With_Admin_Priv_Valid_User() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	other := getTestAdminWithPriv("other-test-admin@admin.com", "All")
	verrs, err = other.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/" + other.ID.String() + "/edit/").Get()

	as.Equal(200, res.Code)
	// Page header
	as.Contains(res.Body.String(), "<h1>Edit Administrator</h1>")
	// Form
	as.Contains(res.Body.String(), "<form action=\"/administrators/"+other.ID.String()+"/\" id=\"administrator-form\" method=\"POST\">")
	// Name label
	as.Contains(res.Body.String(), "<label>Name</label>")
	// Admin name
	as.Contains(res.Body.String(), other.Name)
	// Username label
	as.Contains(res.Body.String(), "<label>Username</label>")
	// Admin username
	as.Contains(res.Body.String(), other.Username)
	// Email label
	as.Contains(res.Body.String(), "<label>Email</label>")
	// Admin email
	as.Contains(res.Body.String(), other.Email)
	// Privs label
	as.Contains(res.Body.String(), "Setup privileges")
}

func (as *ActionSuite) Test_AdministratorsResource_Update_Without_Logged_In_User() {
	res := as.HTML("/administrators/fakeif/").Put(nil)
	as.Equal(302, res.Code)
	as.Equal("/login", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Update_Without_Privilege() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/fakeid/").Put(nil)
	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Update_With_Priv_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "Privileges")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/fakeid/").Put(nil)
	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Update_With_Admin_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "Users")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/fakeid/").Put(nil)
	as.Equal(404, res.Code)
}
func (as *ActionSuite) Test_AdministratorsResource_Update_With_All_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/fakeid/").Put(nil)
	as.Equal(404, res.Code)
}
func (as *ActionSuite) Test_AdministratorsResource_Update_With_Admin_Priv_Priv_Data() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	other := getTestAdminWithPriv("other-test-admin@admin.com", "All")
	verrs, err = other.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/" + other.ID.String() + "/").Put(other)
	as.Equal(302, res.Code)
	as.Contains(res.Location(), "/administrators/")
}

func (as *ActionSuite) Test_AdministratorsResource_Destroy_Without_Logged_In_User() {
	res := as.HTML("/administrators/fakeif/").Delete()
	as.Equal(302, res.Code)
	as.Equal("/login", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Destroy_Without_Privilege() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/fakeif/").Delete()
	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}
func (as *ActionSuite) Test_AdministratorsResource_Destroy_With_Priv() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/fakeif/").Delete()
	as.Equal(404, res.Code)
}
func (as *ActionSuite) Test_AdministratorsResource_Destroy_With_Priv_Valid_User() {
	as.LoadFixture("insert privileges")
	u := getTestAdminWithPriv("test-admin@admin.com", "All")

	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	other := getTestAdminWithPriv("other-test-admin@admin.com", "All")
	verrs, err = other.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	_ = as.HTML("/login").Post(u)
	res := as.HTML("/administrators/" + other.ID.String() + "/").Delete()
	as.Equal(302, res.Code)
	as.Equal("/administrators/", res.Location())
}
