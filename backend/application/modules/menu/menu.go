package menu

import (
	"context"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/shiwyang/snapify/backend/application/modules/menu/internal/controller"
)

var _menu = &MenuModule{}

type MenuModule struct {
	once       sync.Once
	controller controller.MenuModuleController
}

func GetMenuModule() *MenuModule {
	_menu.once.Do(func() {
		_menu = &MenuModule{
			controller: controller.MenuModuleController{},
		}
	})
	return _menu
}

func (m *MenuModule) Start(ctx context.Context) {
	log.Infof("MenuModule start")
}

func (m *MenuModule) Shoutdown(ctx context.Context) {
	log.Infof("MenuModule shoutdown")
}

func (m *MenuModule) Bind() []interface{} {
	return []interface{}{
		&m.controller,
	}
}