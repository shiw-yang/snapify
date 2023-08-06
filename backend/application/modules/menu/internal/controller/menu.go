package controller

type MenuModuleController struct {
}

func (c MenuModuleController) GetMenusName() []string {
	return []string{"Snap", "Profile", "Setting"}
}
