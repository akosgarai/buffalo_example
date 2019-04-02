package models

func (ms *ModelSuite) Test_Privilege_String() {
	p := &Privilege{}
	ms.Equal("{\"privilege_id\":\"00000000-0000-0000-0000-000000000000\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"label\":\"\",\"description\":\"\"}", p.String())
}
func (ms *ModelSuite) Test_Privileges_Empty_String() {
	p := &Privileges{}
	ms.Equal("[]", p.String())
}
func (ms *ModelSuite) Test_Privileges_Not_Empty_String() {
	var p Privileges
	var priv Privilege
	p = append(p, priv)
	ms.Equal("[{\"privilege_id\":\"00000000-0000-0000-0000-000000000000\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"label\":\"\",\"description\":\"\"}]", p.String())
}
func (ms *ModelSuite) Test_Privilege_Validate_No_Label() {
	var priv Privilege
	priv.Description = "test desc"
	verrs, err := priv.Validate(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}
func (ms *ModelSuite) Test_Privilege_Validate_With_Label() {
	var priv Privilege
	priv.Description = "test desc"
	priv.Label = "test label"
	verrs, err := priv.Validate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_Privilege_ValidateCreate_No_Label() {
	var priv Privilege
	priv.Description = "test desc"
	verrs, err := priv.ValidateCreate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_Privilege_ValidateCreate_With_Label() {
	var priv Privilege
	priv.Description = "test desc"
	priv.Label = "test label"
	verrs, err := priv.ValidateCreate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_Privilege_ValidateUpdate_No_Label() {
	var priv Privilege
	priv.Description = "test desc"
	verrs, err := priv.ValidateUpdate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_Privilege_ValidateUpdate_With_Label() {
	var priv Privilege
	priv.Description = "test desc"
	priv.Label = "test label"
	verrs, err := priv.ValidateUpdate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
