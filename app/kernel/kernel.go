package kernel

import "github.com/google/wire"

type Kernel struct {
}

var kernelSet = wire.NewSet(
	wire.Struct(new(Kernel), "*"),
)
