package page

type HardwareExecution interface {
	ExecuteHardwareConfig() error
	ExecuteHardwareCPU() error
	ExecuteHardwareRAM() error
	ExecuteHardwareDisk() error
	ExecuteHardwareNIC() error
	ExecuteHardwareRaid() error
	ExecuteHardwarePSU() error
	ExecuteHardwareMNT() error
}
