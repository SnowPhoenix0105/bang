package with

type DeepCopyOption int32

const (

	// 1x Interface

	InterfaceBitwiseCopy    = 10
	InterfaceSetNil         = 11
	InterfaceDeepCopyUnsafe = 12

	// 2x Map

	MapBitwiseCopyKey = 20
	MapDeepCopyKey    = 21
)
