package menu

import (
	"context"

	"github.com/labstack/gommon/log"
)

type MenuModule struct {
}

func NewMenuModule() *MenuModule {
	return &MenuModule{}
}

func (m *MenuModule) Start(ctx context.Context) {
	log.Infof("MenuModule start")
}

func (m *MenuModule) Shoutdown(ctx context.Context) {
	log.Infof("MenuModule shoutdown")
}

func (m *MenuModule) GetMenusName() []string {
	return []string{"Snap", "Profile", "Setting"}
}
