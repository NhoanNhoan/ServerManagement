package page

import (
	"CURD/entity/server/hardware"
	"CURD/repo/server/hardware_repo"
)

type HardwareInsertion struct {
	hardware.HardwareConfig
	HardwareCPU  []hardware.HardwareCpu
	HardwareRAM  []hardware.HardwareRam
	HardwareDisk []hardware.HardwareDisk
	HardwareNIC  []hardware.HardwareNic
	HardwareRaid []hardware.HardwareRaid
	HardwarePSU  []hardware.HardwarePsu
	HardwareMNT  []hardware.HardwareManagement
}

func (ins HardwareInsertion) ExecuteHardwareConfig() error {
	repo := hardware_repo.HardwareConfigRepo{}
	return repo.Insert(ins.HardwareConfig)
}

func (ins HardwareInsertion) ExecuteHardwareCPU() error {
	repo := hardware_repo.HardwareCPURepo{}
	return repo.Insert(ins.HardwareConfig.Id, ins.HardwareCPU...)
}

func (ins HardwareInsertion) ExecuteHardwareRAM() error {
	repo := hardware_repo.HardwareRamRepo{}
	return repo.Insert(ins.HardwareConfig.Id, ins.HardwareRAM...)
}

func (ins HardwareInsertion) ExecuteHardwareDisk() error {
	repo := hardware_repo.HardwareDiskRepo{}
	return repo.Insert(ins.HardwareConfig.Id, ins.HardwareDisk...)
}

func (ins HardwareInsertion) ExecuteHardwareNIC() error {
	repo := hardware_repo.HardwareNicRepo{}
	return repo.Insert(ins.HardwareConfig.Id, ins.HardwareNIC...)
}

func (ins HardwareInsertion) ExecuteHardwareRaid() error {
	repo := hardware_repo.HardwareRaidRepo{}
	return repo.Insert(ins.HardwareConfig.Id, ins.HardwareRaid...)
}

func (ins HardwareInsertion) ExecuteHardwarePSU() error {
	repo := hardware_repo.HardwarePsuRepo{}
	return repo.Insert(ins.HardwareConfig.Id, ins.HardwarePSU...)
}

func (ins HardwareInsertion) ExecuteHardwareMNT() error {
	repo := hardware_repo.HardwareManagementRepo{}
	return repo.Insert(ins.HardwareConfig.Id, ins.HardwareMNT...)
}