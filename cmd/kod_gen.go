// Code generated by "kod generate". DO NOT EDIT.
//go:build !ignoreKodGen

package cmd

import (
	"context"
	"github.com/go-kod/kod"
	"github.com/go-kod/kod/interceptor"
	"reflect"
)

func init() {
	kod.Register(&kod.Registration{
		Name:      "github.com/leopku/kodtest/cmd/Demo1",
		Interface: reflect.TypeOf((*Demo1)(nil)).Elem(),
		Impl:      reflect.TypeOf(demo1Impl{}),
		Refs:      ``,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return demo1_local_stub{
				impl:        info.Impl.(Demo1),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/leopku/kodtest/cmd/Demo2",
		Interface: reflect.TypeOf((*Demo2)(nil)).Elem(),
		Impl:      reflect.TypeOf(demo2Impl{}),
		Refs:      `⟦73ce26ff:KoDeDgE:github.com/leopku/kodtest/cmd/Demo2→github.com/leopku/kodtest/pkg/migrate/Migrate⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return demo2_local_stub{
				impl:        info.Impl.(Demo2),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
	kod.Register(&kod.Registration{
		Name:      "github.com/go-kod/kod/Main",
		Interface: reflect.TypeOf((*kod.Main)(nil)).Elem(),
		Impl:      reflect.TypeOf(app{}),
		Refs: `⟦d9632c20:KoDeDgE:github.com/go-kod/kod/Main→github.com/leopku/kodtest/cmd/Demo1⟧,
⟦11d40cc0:KoDeDgE:github.com/go-kod/kod/Main→github.com/leopku/kodtest/cmd/Demo2⟧`,
		LocalStubFn: func(ctx context.Context, info *kod.LocalStubFnInfo) any {
			interceptors := info.Interceptors
			if h, ok := info.Impl.(interface {
				Interceptors() []interceptor.Interceptor
			}); ok {
				interceptors = append(interceptors, h.Interceptors()...)
			}

			return main_local_stub{
				impl:        info.Impl.(kod.Main),
				interceptor: interceptor.Chain(interceptors),
				name:        info.Name,
			}
		},
	})
}

// kod.InstanceOf checks.
var _ kod.InstanceOf[Demo1] = (*demo1Impl)(nil)
var _ kod.InstanceOf[Demo2] = (*demo2Impl)(nil)
var _ kod.InstanceOf[kod.Main] = (*app)(nil)

// Local stub implementations.

type demo1_local_stub struct {
	impl        Demo1
	name        string
	interceptor interceptor.Interceptor
}

// Check that demo1_local_stub implements the Demo1 interface.
var _ Demo1 = (*demo1_local_stub)(nil)

type demo2_local_stub struct {
	impl        Demo2
	name        string
	interceptor interceptor.Interceptor
}

// Check that demo2_local_stub implements the Demo2 interface.
var _ Demo2 = (*demo2_local_stub)(nil)

func (s demo2_local_stub) Migrate(ctx context.Context, a1 int) (err error) {

	if s.interceptor == nil {
		err = s.impl.Migrate(ctx, a1)
		return
	}

	call := func(ctx context.Context, info interceptor.CallInfo, req, res []any) (err error) {
		err = s.impl.Migrate(ctx, a1)
		return
	}

	info := interceptor.CallInfo{
		Impl:       s.impl,
		Component:  s.name,
		FullMethod: "github.com/leopku/kodtest/cmd/Demo2.Migrate",
		Method:     "Migrate",
	}

	err = s.interceptor(ctx, info, []any{a1}, []any{}, call)
	return
}

type main_local_stub struct {
	impl        kod.Main
	name        string
	interceptor interceptor.Interceptor
}

// Check that main_local_stub implements the kod.Main interface.
var _ kod.Main = (*main_local_stub)(nil)

