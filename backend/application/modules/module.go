package module

import "context"


type IModule interface {
	Start(ctx context.Context)
	Shoutdown(ctx context.Context)
	Bind() []interface{}
}
