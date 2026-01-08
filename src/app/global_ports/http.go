package global_ports

type HTTPPort interface {
	Start(addr string) error
	RegisterControllers(controllers ...HTTPRouteRegister)
}

type HTTPRouteRegister interface {
	Register(router any)
}
