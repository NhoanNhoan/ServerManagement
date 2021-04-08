package page

import (
	"CURD/entity"
	"CURD/repo/server"
)

// ============================================

type NetworkPortionsItem struct {
	entity.NetworkPortion
	NumAvailableHosts int
	NumUsedHosts	int
}

type ListNetworkPortions struct {
	Portions []NetworkPortionsItem
}

func (p *ListNetworkPortions) New() (err error) {
	portions, err := server.NetworkPortionRepo{}.FetchAll()
	if nil != err {
		return err
	}

	p.Portions = make([]NetworkPortionsItem, len(portions))
	for i := range portions {
		p.Portions[i].NetworkPortion = portions[i]
		p.Portions[i].NumAvailableHosts, err = server.IpRepo{}.CountStates(p.Portions[i].Id, "available")
		if nil != err {
			return err
		}

		p.Portions[i].NumUsedHosts, err = server.IpRepo{}.CountStates(p.Portions[i].Id, "used")
		if nil != err {
			return err
		}
	}

	return nil
}