package models

import (
	"github.com/gobuffalo/uuid"
)

func getTestAdminWithPriv(email string, privName string) *Administrator {
	var privList Privileges
	if privName == "Privileges" || privName == "All" {
		uid, _ := uuid.FromString("d65b85e0-a5ac-492a-92c9-9040bba1a981")
		privList = append(privList, Privilege{ID: uid})
	}
	if privName == "Users" || privName == "All" {
		uid, _ := uuid.FromString("3d41e152-7f33-49e6-829e-411446d495e8")
		privList = append(privList, Privilege{ID: uid})
	}
	return &Administrator{
		Email:           email,
		Password:        "admin",
		Pwd:             "admin",
		PwdConfirmation: "admin",
		Name:            "test admin",
		Username:        "testAdmin",
		Privs:           privList,
	}
}
func (ms *ModelSuite) Test_AdminPriv_String() {
	ap := &AdminPriv{}
	ms.Equal("{\"id\":\"00000000-0000-0000-0000-000000000000\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"administrator_id\":\"00000000-0000-0000-0000-000000000000\",\"privilege_id\":\"00000000-0000-0000-0000-000000000000\"}", ap.String())
}
func (ms *ModelSuite) Test_AdminPriv_ValidateCreate_InvalidAdminAndPriv() {
	var ap AdminPriv
	verrs, err := ap.ValidateCreate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_AdminPriv_ValidateCreate_RandomAdminAndPriv() {
	var ap AdminPriv
	ap.AdministratorID = uuid.Must(uuid.NewV4())
	ap.PrivilegeID = uuid.Must(uuid.NewV4())
	verrs, err := ap.ValidateCreate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_AdminPriv_ValidateCreate_Valid() {
	ms.LoadFixture("insert privileges")
	admin := getTestAdminWithPriv("test-admin-non-validatecreate@admin.com", "")
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	lastAdmin := Administrator{}
	err = ms.DB.Last(&lastAdmin)
	ms.NoError(err)
	ms.DB.Load(&lastAdmin)
	ms.Equal(0, len(lastAdmin.Privs))
	var ap AdminPriv
	ap.AdministratorID = lastAdmin.ID
	ap.PrivilegeID = uuid.Must(uuid.FromString("d65b85e0-a5ac-492a-92c9-9040bba1a981"))
	verrs, err = ap.ValidateCreate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_AdminPriv_ValidateUpdate_RandomAdminAndPrivWoDB() {
	var ap AdminPriv
	ap.AdministratorID = uuid.Must(uuid.NewV4())
	ap.PrivilegeID = uuid.Must(uuid.NewV4())
	verrs, err := ap.ValidateUpdate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_AdminPriv_ValidateUpdate_RandomAdminAndPriv() {
	ms.LoadFixture("insert privileges")
	var ap AdminPriv
	ap.AdministratorID = uuid.Must(uuid.NewV4())
	ap.PrivilegeID = uuid.Must(uuid.NewV4())
	verrs, err := ap.ValidateUpdate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_AdminPriv_Validate_RandomAdminAndPrivWoDB() {
	var ap AdminPriv
	ap.AdministratorID = uuid.Must(uuid.NewV4())
	ap.PrivilegeID = uuid.Must(uuid.NewV4())
	verrs, err := ap.Validate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_AdminPriv_Validate_RandomAdminAndPriv() {
	ms.LoadFixture("insert privileges")
	var ap AdminPriv
	ap.AdministratorID = uuid.Must(uuid.NewV4())
	ap.PrivilegeID = uuid.Must(uuid.NewV4())
	verrs, err := ap.Validate(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
}
func (ms *ModelSuite) Test_AdminPrivs_String_Empty() {
	var aps AdminPrivs
	ms.Equal("{\"AdministratorID\":\"00000000-0000-0000-0000-000000000000\",\"PrivilegeIDs\":null}", aps.String())
}
func (ms *ModelSuite) Test_AdminPrivs_Update_AddPriv() {
	ms.LoadFixture("insert privileges")
	admin := getTestAdminWithPriv("test-admin-all-delete@admin.com", "")
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	lastAdmin := Administrator{}
	err = ms.DB.Last(&lastAdmin)
	ms.NoError(err)
	ms.DB.Load(&lastAdmin)
	ms.Equal(0, len(lastAdmin.Privs))
	aps := AdminPrivs{}
	aps.AdministratorID = lastAdmin.ID
	uid, _ := uuid.FromString("d65b85e0-a5ac-492a-92c9-9040bba1a981")
	aps.PrivilegeIDs = append(aps.PrivilegeIDs, uid)
	verrs, err = aps.Update(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.DB.Load(&lastAdmin)
	ms.Equal(1, len(lastAdmin.Privs))
}
func (ms *ModelSuite) Test_AdminPrivs_Update_RevokePriv() {
	ms.LoadFixture("insert privileges")
	admin := getTestAdminWithPriv("test-admin-all-delete@admin.com", "All")
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	lastAdmin := Administrator{}
	err = ms.DB.Last(&lastAdmin)
	ms.NoError(err)
	ms.DB.Load(&lastAdmin)
	ms.Equal(2, len(lastAdmin.Privs))
	aps := AdminPrivs{}
	aps.AdministratorID = lastAdmin.ID
	uid, _ := uuid.FromString("d65b85e0-a5ac-492a-92c9-9040bba1a981")
	aps.PrivilegeIDs = append(aps.PrivilegeIDs, uid)
	verrs, err = aps.Update(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.DB.Load(&lastAdmin)
	ms.Equal(1, len(lastAdmin.Privs))
}
func (ms *ModelSuite) Test_AdminPrivs_Delete() {
	ms.LoadFixture("insert privileges")
	admin := getTestAdminWithPriv("test-admin-all-delete@admin.com", "All")
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	lastAdmin := Administrator{}
	err = ms.DB.Last(&lastAdmin)
	ms.NoError(err)
	ms.DB.Load(&lastAdmin)
	ms.Equal(2, len(lastAdmin.Privs))
	aps := AdminPrivs{}
	aps.AdministratorID = lastAdmin.ID
	verrs, err = aps.Delete(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.DB.Load(&lastAdmin)
	ms.Equal(0, len(lastAdmin.Privs))
}
func (ms *ModelSuite) Test_AdminPrivs_deletePrivileges_FakeAdmin() {
	ms.LoadFixture("insert privileges")
	aps := AdminPrivs{}
	uid, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	aps.AdministratorID = uid
	err := aps.deletePrivileges(ms.DB)
	ms.NoError(err)
}
func (ms *ModelSuite) Test_AdminPrivs_deletePrivileges_ValidAdmin() {
	ms.LoadFixture("insert privileges")
	admin := getTestAdminWithPriv("test-admin-all@admin.com", "All")
	verrs, err := admin.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	lastAdmin := Administrator{}
	err = ms.DB.Last(&lastAdmin)
	ms.NoError(err)
	ms.DB.Load(&lastAdmin)
	ms.Equal(2, len(lastAdmin.Privs))
	aps := AdminPrivs{}
	aps.AdministratorID = lastAdmin.ID
	err = aps.deletePrivileges(ms.DB)
	ms.NoError(err)
	ms.DB.Load(&lastAdmin)
	ms.Equal(0, len(lastAdmin.Privs))
}
