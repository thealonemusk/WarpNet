package vpn

import (
	"net/netip"

	"github.com/mudler/water"
	"golang.org/x/sys/windows"
	"golang.zx2c4.com/wireguard/windows/tunnel/winipcfg"
)

func prepareInterface(c *Config) error {
	// find interface created by water
	guid, err := windows.GUIDFromString("{00000000-FFFF-FFFF-FFE9-76E58C74063E}")
	if err != nil {
		return err
	}
	luid, err := winipcfg.LUIDFromGUID(&guid)
	if err != nil {
		return err
	}

	prefix, err := netip.ParsePrefix(c.InterfaceAddress)
	if err != nil {
		return err
	}
	addresses := append([]netip.Prefix{}, prefix)
	if err := luid.SetIPAddresses(addresses); err != nil {
		return err
	}

	iface, err := luid.IPInterface(windows.AF_INET)
	if err != nil {
		return err
	}
	iface.NLMTU = uint32(c.InterfaceMTU)
	if err := iface.Set(); err != nil {
		return err
	}
	return nil
}

func createInterface(c *Config) (*water.Interface, error) {
	config := water.Config{
		DeviceType: c.DeviceType,
	}
	config.Name = c.InterfaceName
	return water.New(config)
}
