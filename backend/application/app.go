package app

import (
	"context"
	"sync"

	"github.com/labstack/gommon/log"
	module "github.com/shiwyang/snapify/backend/application/modules"
	"github.com/shiwyang/snapify/backend/application/modules/menu"
	"github.com/shiwyang/snapify/backend/application/modules/profile"
	"github.com/shiwyang/snapify/backend/runtime/config"
)

var _app = &app{}

// app struct
type app struct {
	once    sync.Once
	ctx     context.Context
	cfg     config.AppConfig
	Modules []module.IModule
}

// App creates a new App application struct
func App() *app {
	_app.once.Do(func() {
		// load config
		_app.cfg.Load()

		// load modules
		_app.Modules = make([]module.IModule, 0)
		// add modules
		_app.addModule(profile.NewProfileModule())
		_app.addModule(menu.GetMenuModule())
	})
	return _app
}

// OnStartup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *app) OnStartup(ctx context.Context) {
	a.ctx = ctx

}

// addModule add module
func (a *app) addModule(module module.IModule) {
	module.Start(a.ctx)
	a.Modules = append(a.Modules, module)
	log.Info("Add module:", module)
}

func (a *app) Bind() []interface{} {
	bind := make([]interface{}, 0)
	for _, module := range a.Modules {
		bind = append(bind, module.Bind()...)
	}
	return bind
}
