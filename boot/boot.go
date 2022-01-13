package boot

type boot struct{}

var Boot = boot{}

func (*boot) Init() {
	// 初始化路由
	Route.Init()
}
