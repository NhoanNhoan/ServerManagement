package entity

type IpAddress struct {
	IpNet
	IpHost string
}

func (obj *IpAddress) New(ipNetValue string, ipHost string) {
	obj.IpNet = IpNet {Value: ipNetValue}
	obj.IpHost = ipHost
}

func (obj *IpAddress) String() string {
	return obj.IpNet.Value + obj.IpHost
}
