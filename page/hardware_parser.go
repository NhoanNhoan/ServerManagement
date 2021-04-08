package page

import (
	"CURD/entity/server/hardware"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type HardwareParser struct {
	c *gin.Context
}

func (p HardwareParser) post_parse(name string) string {
	return p.c.PostForm(name)
}

func (p HardwareParser) HardwareConfig() hardware.HardwareConfig {
	return hardware.HardwareConfig{
		ChassisId: p.post_parse("cbChassis"),
		ClusterId: p.post_parse("cbCluster"),
	}
}

// Parse Hardware CPU
func (p HardwareParser) HardwareCPUArray() []hardware.HardwareCpu {
	raw := p.post_parse("txtAllCPUs")
	return p.hardwareCPUArray_by_string(raw)
}

func (p HardwareParser) hardwareCPUArray_by_string(raw string) []hardware.HardwareCpu {
	raws := strings.Split(raw, ",")
	return p.hardwareCPUArray_by_string_arr(raws)
}

func (p HardwareParser) hardwareCPUArray_by_string_arr(raws []string) []hardware.HardwareCpu {
	list := make([]hardware.HardwareCpu, len(raws))
	for i := range list {
		list[i] = *p.hardwareCPU_by_string(raws[i])
	}
	return list
}

func (p HardwareParser) hardwareCPU_by_string(content string) (hw *hardware.HardwareCpu) {
	return p.hardwareCPU_by_values(strings.Split(content, "-")...)
}

func (p HardwareParser) hardwareCPU_by_values(values ...string) (hw *hardware.HardwareCpu) {
	var err error
	hw = &hardware.HardwareCpu{}
	if hw.NumberCpu, err = strconv.Atoi(values[1]); nil != err {
		return nil
	}

	hw.CpuId = values[0]
	return
}

// Parse Hardware RAM
func (p HardwareParser) HardwareRAMArray() []hardware.HardwareRam {
	raw := p.post_parse("txtAllRAMs")
	return p.hardwareRAMArray_by_string(raw)
}

func (p HardwareParser) hardwareRAMArray_by_string(raw string) []hardware.HardwareRam {
	raws := strings.Split(raw, ",")
	return p.hardwareRAMArray_by_string_arr(raws)
}

func (p HardwareParser) hardwareRAMArray_by_string_arr(raws []string) []hardware.HardwareRam {
	list := make([]hardware.HardwareRam, len(raws))
	for i := range list {
		list[i] = *p.hardwareRAM_by_string(raws[i])
	}
	return list
}

func (p HardwareParser) hardwareRAM_by_string(content string) (hw *hardware.HardwareRam) {
	return p.hardwareRAM_by_values(strings.Split(content, "-")...)
}

func (p HardwareParser) hardwareRAM_by_values(values ...string) (hw *hardware.HardwareRam) {
	var err error
	hw = &hardware.HardwareRam{}
	if hw.NumberRam, err = strconv.Atoi(values[1]); nil != err {
		return nil
	}

	hw.RamId = values[0]
	return
}

// Parse Hardware Disk
func (p HardwareParser) HardwareDiskArray() []hardware.HardwareDisk {
	raw := p.post_parse("txtAllDisks")
	return p.hardwareDiskArray_by_string(raw)
}

func (p HardwareParser) hardwareDiskArray_by_string(raw string) []hardware.HardwareDisk {
	raws := strings.Split(raw, ",")
	return p.hardwareDiskArray_by_string_arr(raws)
}

func (p HardwareParser) hardwareDiskArray_by_string_arr(raws []string) []hardware.HardwareDisk {
	list := make([]hardware.HardwareDisk, len(raws))
	for i := range list {
		list[i] = *p.hardwareDisk_by_string(raws[i])
	}
	return list
}

func (p HardwareParser) hardwareDisk_by_string(content string) (hw *hardware.HardwareDisk) {
	return p.hardwareDisk_by_values(strings.Split(content, "-")...)
}

func (p HardwareParser) hardwareDisk_by_values(values ...string) (hw *hardware.HardwareDisk) {
	var err error
	hw = &hardware.HardwareDisk{}
	if hw.NumberDisk, err = strconv.Atoi(values[1]); nil != err {
		return nil
	}

	hw.DiskId = values[0]
	return
}

// Parse Hardware Raid
func (p HardwareParser) HardwareRaidArray() []hardware.HardwareRaid {
	raw := p.post_parse("txtAllRaids")
	return p.hardwareRaidArray_by_string(raw)
}

func (p HardwareParser) hardwareRaidArray_by_string(raw string) []hardware.HardwareRaid {
	raws := strings.Split(raw, ",")
	return p.hardwareRaidArray_by_string_arr(raws)
}

func (p HardwareParser) hardwareRaidArray_by_string_arr(raws []string) []hardware.HardwareRaid {
	list := make([]hardware.HardwareRaid, len(raws))
	for i := range list {
		list[i] = *p.hardwareRaid_by_string(raws[i])
	}
	return list
}

func (p HardwareParser) hardwareRaid_by_string(content string) (hw *hardware.HardwareRaid) {
	return p.hardwareRaid_by_values(strings.Split(content, "-")...)
}

func (p HardwareParser) hardwareRaid_by_values(values ...string) (hw *hardware.HardwareRaid) {
	var err error
	hw = &hardware.HardwareRaid{}
	if hw.NumberRaid, err = strconv.Atoi(values[1]); nil != err {
		return nil
	}

	hw.RaidId = values[0]
	return
}

// Parse Hardware NIC
func (p HardwareParser) HardwareNicArray() []hardware.HardwareNic {
	raw := p.post_parse("txtAllNICs")
	return p.hardwareNicArray_by_string(raw)
}

func (p HardwareParser) hardwareNicArray_by_string(raw string) []hardware.HardwareNic {
	raws := strings.Split(raw, ",")
	return p.hardwareNicArray_by_string_arr(raws)
}

func (p HardwareParser) hardwareNicArray_by_string_arr(raws []string) []hardware.HardwareNic {
	list := make([]hardware.HardwareNic, len(raws))
	for i := range list {
		list[i] = *p.hardwareNic_by_string(raws[i])
	}
	return list
}

func (p HardwareParser) hardwareNic_by_string(content string) (hw *hardware.HardwareNic) {
	return p.hardwareNic_by_values(strings.Split(content, "-")...)
}

func (p HardwareParser) hardwareNic_by_values(values ...string) (hw *hardware.HardwareNic) {
	var err error
	hw = &hardware.HardwareNic{}
	if hw.NumberNic, err = strconv.Atoi(values[1]); nil != err {
		return nil
	}

	hw.NicId = values[0]
	return
}

// Parse Hardware PSU
func (p HardwareParser) HardwarePsuArray() []hardware.HardwarePsu {
	raw := p.post_parse("txtAllPSUs")
	return p.hardwarePsuArray_by_string(raw)
}

func (p HardwareParser) hardwarePsuArray_by_string(raw string) []hardware.HardwarePsu {
	raws := strings.Split(raw, ",")
	return p.hardwarePsuArray_by_string_arr(raws)
}

func (p HardwareParser) hardwarePsuArray_by_string_arr(raws []string) []hardware.HardwarePsu {
	list := make([]hardware.HardwarePsu, len(raws))
	for i := range list {
		list[i] = *p.hardwarePsu_by_string(raws[i])
	}
	return list
}

func (p HardwareParser) hardwarePsu_by_string(content string) (hw *hardware.HardwarePsu) {
	return p.hardwarePsu_by_values(strings.Split(content, "-")...)
}

func (p HardwareParser) hardwarePsu_by_values(values ...string) (hw *hardware.HardwarePsu) {
	var err error
	hw = &hardware.HardwarePsu{}
	if hw.NumberPsu, err = strconv.Atoi(values[1]); nil != err {
		return nil
	}

	hw.PsuId = values[0]
	return
}

// Parse Hardware MNT
func (p HardwareParser) HardwareMntArray() []hardware.HardwareManagement {
	raw := p.post_parse("txtAllMNTs")
	return p.hardwareMntArray_by_string(raw)
}

func (p HardwareParser) hardwareMntArray_by_string(raw string) []hardware.HardwareManagement {
	raws := strings.Split(raw, ",")
	return p.hardwareMntArray_by_string_arr(raws)
}

func (p HardwareParser) hardwareMntArray_by_string_arr(raws []string) []hardware.HardwareManagement {
	list := make([]hardware.HardwareManagement, len(raws))
	for i := range list {
		list[i] = *p.hardwareMnt_by_string(raws[i])
	}
	return list
}

func (p HardwareParser) hardwareMnt_by_string(content string) (hw *hardware.HardwareManagement) {
	return p.hardwareMnt_by_values(strings.Split(content, "-")...)
}

func (p HardwareParser) hardwareMnt_by_values(values ...string) (hw *hardware.HardwareManagement) {
	hw = &hardware.HardwareManagement{}
	var err error
	if hw.NumberManagement, err = strconv.Atoi(values[1]); nil != err {
		return nil
	}

	hw.ManagementId = values[0]
	return
}
