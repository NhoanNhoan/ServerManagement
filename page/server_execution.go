package page

type ServerExecution interface {
	ExecuteServer() error
	ExecuteTags() error
	HardwareExecution
}
