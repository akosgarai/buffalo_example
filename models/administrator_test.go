package models

func (ms *ModelSuite) Test_Administrators_String() {
	var admin Administrator
	ms.Equal("{\"id\":\"00000000-0000-0000-0000-000000000000\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"name\":\"\",\"username\":\"\",\"password\":\"\",\"email\":\"\",\"Privs\":null}", admin.String())
}
func (ms *ModelSuite) Test_Administrator_Empty_String() {
	admins := &Administrators{}
	ms.Equal("[]", admins.String())
}
func (ms *ModelSuite) Test_Administrator_Not_Empty_String() {
	var admin Administrator
	var admins Administrators
	admins = append(admins, admin)
	ms.Equal("[{\"id\":\"00000000-0000-0000-0000-000000000000\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"name\":\"\",\"username\":\"\",\"password\":\"\",\"email\":\"\",\"Privs\":null}]", admins.String())
}
func (ms *ModelSuite) Test_Administrator_Create_Invalid_Data() {
	var admin Administrator
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}
func (ms *ModelSuite) Test_Administrator_Create_Invalid_Data_EmptyName() {
	var admin Administrator
	admin.Email = "admin@admin.com"
	admin.Username = "test admin username"
	admin.Pwd = "admin"
	admin.PwdConfirmation = "admin"
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}
func (ms *ModelSuite) Test_Administrator_Create_Invalid_Data_EmptyEmail() {
	var admin Administrator
	admin.Name = "test admin name"
	admin.Username = "test admin username"
	admin.Pwd = "admin"
	admin.PwdConfirmation = "admin"
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}
func (ms *ModelSuite) Test_Administrator_Create_Invalid_Data_EmptyPwd() {
	var admin Administrator
	admin.Email = "admin@admin.com"
	admin.Name = "test admin name"
	admin.Username = "test admin username"
	admin.PwdConfirmation = "admin"
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}
func (ms *ModelSuite) Test_Administrator_Create_Invalid_Data_EmptyPwdConfirmation() {
	var admin Administrator
	admin.Email = "admin@admin.com"
	admin.Name = "test admin name"
	admin.Username = "test admin username"
	admin.Pwd = "admin"
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}
func (ms *ModelSuite) Test_Administrator_Create_Invalid_Data_MissingPwd() {
	var admin Administrator
	admin.Email = "admin@admin.com"
	admin.Name = "test admin name"
	admin.Username = "test admin username"
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}
func (ms *ModelSuite) Test_Administrator_Create_Valid_Data() {
	var admin Administrator
	admin.Email = "admin@admin.com"
	admin.Name = "test admin name"
	admin.Username = "test admin username"
	admin.Pwd = "admin"
	admin.PwdConfirmation = "admin"
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_Administrator_HasPrivFor_EmptyLabel() {
	var admin Administrator
	var priv Privilege
	priv.Description = "test desc"
	priv.Label = "test label"
	admin.Privs = append(admin.Privs, priv)
	ms.False(admin.HasPrivFor(""))
}
func (ms *ModelSuite) Test_Administrator_HasPrivFor_MatchingLabel() {
	var admin Administrator
	var priv Privilege
	priv.Description = "test desc"
	priv.Label = "test label"
	admin.Privs = append(admin.Privs, priv)
	ms.True(admin.HasPrivFor("test label"))
}
func (ms *ModelSuite) Test_Administrator_HasPrivFor_NotMatchingLabel() {
	var admin Administrator
	var priv Privilege
	priv.Description = "test desc"
	priv.Label = "test label"
	admin.Privs = append(admin.Privs, priv)
	ms.False(admin.HasPrivFor("404"))
}
func (ms *ModelSuite) Test_Administrator_ValidateUpdate() {
	var admin Administrator
	admin.Email = "admin@admin.com"
	admin.Name = "test admin name"
	admin.Username = "test admin username"
	admin.Pwd = "admin"
	admin.PwdConfirmation = "admin"
	verrs, err := admin.ValidateUpdate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
