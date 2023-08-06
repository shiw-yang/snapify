package profile

import (
	"context"

	"github.com/labstack/gommon/log"
)

type ProfileModule struct {
}

func NewProfileModule() *ProfileModule {
	return &ProfileModule{}
}

func (m *ProfileModule) Start(ctx context.Context) {
	log.Infof("ProfileModule start")
}

func (m *ProfileModule) Shoutdown(ctx context.Context) {
	log.Infof("ProfileModule shoutdown")
}

func (m *ProfileModule) Greet(name string) string {
	return "Hello" + name + ", It's show time!"
}
