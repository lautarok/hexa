package global_ports

type EnvPort interface {
	Load() error
	GetStr(key string) (string, error)
}
