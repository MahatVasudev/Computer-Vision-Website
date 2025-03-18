package permission

import "fmt"

func ProfilePermission(s string) string {
	return fmt.Sprintf("permission:%s", s)
}

var Editpermission = ProfilePermission("edit")

var Viewpermission = ProfilePermission("view")

var Analyticspermission = ProfilePermission("analytics")

var Deletepermission = ProfilePermission("delete")

type PermissionEntity struct {
	permissions []string
	status      string
}

var AdminControl = []string{
	Editpermission,
	Viewpermission,
	Analyticspermission,
	Deletepermission,
}

var StaffControl = []string{
	Viewpermission,
	Analyticspermission,
}

var UserControl = []string{
	Editpermission,
	Viewpermission,
	Analyticspermission,
	Deletepermission,
}
