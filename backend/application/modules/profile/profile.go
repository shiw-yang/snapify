package profile

import (
	"context"
	"sync"

	"github.com/labstack/gommon/log"
)

var _profile = &ProfileModule{}

type ProfileModule struct {
	once    sync.Once
	service ProfileModuleService
}

type ProfileModuleService struct {
}

func NewProfileModule() *ProfileModule {
	_profile.once.Do(func() {
		_profile = &ProfileModule{
			service: ProfileModuleService{},
		}
	})
	return _profile
}

func (m *ProfileModule) Start(ctx context.Context) {
	log.Infof("ProfileModule start")
}

func (m *ProfileModule) Shoutdown(ctx context.Context) {
	log.Infof("ProfileModule shoutdown")
}

func (m *ProfileModule) Bind() []interface{} {
	return []interface{}{
		&m.service,
	}
}

func (s *ProfileModuleService) Greet(name string) string {
	return "Hello" + name + ", It's show time!"
}
