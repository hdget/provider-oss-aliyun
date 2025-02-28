package oss_aliyun

import (
	"github.com/hdget/common/types"
	"github.com/hdget/provider-oss-aliyun/pkg"
	"go.uber.org/fx"
)

const (
	providerName = "oss-aliyun"
)

var Capability = &types.Capability{
	Category: types.ProviderCategoryOss,
	Name:     providerName,
	Module: fx.Module(
		providerName,
		fx.Provide(pkg.New),
	),
}
