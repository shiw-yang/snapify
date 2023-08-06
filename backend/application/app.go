package app

import (
	"context"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/shiwyang/snapify/backend/application/modules/menu"
	"github.com/shiwyang/snapify/backend/application/modules/profile"
	"github.com/shiwyang/snapify/backend/runtime/config"
)

type IMoudle interface {
	Start(ctx context.Context)
	Shoutdown(ctx context.Context)
}

var _app = &app{}

// app struct
type app struct {
	once    sync.Once
	ctx     context.Context
	cfg     config.AppConfig
	Modules []IMoudle
}

// App creates a new App application struct
func App() *app {
	_app.once.Do(func() {
		// load config
		_app.cfg.Load()

		// load modules
		_app.Modules = make([]IMoudle, 0)

		_app.AddModule(menu.NewMenuModule())
		_app.AddModule(profile.NewProfileModule())
	})
	return _app
}

// OnStartup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *app) OnStartup(ctx context.Context) {
	a.ctx = ctx

}

// AddModule add module
func (a *app) AddModule(module IMoudle) {
	if a.Modules == nil {
		a.Modules = make([]IMoudle, 2)
	}
	module.Start(a.ctx)
	a.Modules = append(a.Modules, module)
	log.Info("Add module:", module)
}

func (a *app) Bind() []interface{} {
	bind := make([]interface{}, 0)
	for _, module := range a.Modules {
		bind = append(bind, module)
	}
	return bind
}
