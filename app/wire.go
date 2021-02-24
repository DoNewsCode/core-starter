//+build wireinject

package app

import (
    "github.com/DoNewsCode/core/contract"
    "github.com/google/wire"
)

func InjectKernel(env contract.Env) (Kernel, error) {
    panic(
        wire.Build(
            wire.Struct(new(Kernel), "*"),
        ),
    )
}
