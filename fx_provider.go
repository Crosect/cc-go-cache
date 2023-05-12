package cccache

import (
	"github.com/crosect/cc-go"
	"go.uber.org/fx"
)

func EnableCache() fx.Option {
	return fx.Options(
		ccgo.ProvideProps(NewCacheProperties),
		fx.Provide(NewCache),
	)
}
