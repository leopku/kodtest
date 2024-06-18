// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package migrate

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod/interceptor"
	"reflect"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/leopku/kodtest/pkg/migrate/IMigrate",
		Interface: reflect.TypeOf((*IMigrate)(nil)).Elem(),
		Impl:      reflect.TypeOf(Migrate{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return iMigrate_local_stub{
				impl:        info.Impl.(IMigrate),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[IMigrate] = (*Migrate)(nil)

// Local stub implementations.

type iMigrate_local_stub struct {
	impl        IMigrate
	name        string
	interceptor interceptor.Interceptor
}

// Check that iMigrate_local_stub implements the IMigrate interface.
var _ IMigrate = (*iMigrate_local_stub)(nil)

func (s iMigrate_local_stub) Hello() {
	// Because the first argument is not context.Context, so interceptors are not supported.
	s.impl.Hello()
	return
}
