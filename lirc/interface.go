package lirc

// ClientAPI represents LIRC Client
type ClientAPI interface {
	Version() (version string, err error)
	List(remote string, code ...string) (replies []string, err error)
	SendOnce(remote string, code ...string) (err error)
	SendStart(remote string, code string) (err error)
	SendStop(remote string, code string) (err error)
	Close() error

	send(cmd, remote string, code ...string) ([]Reply, error)
	read()
}
