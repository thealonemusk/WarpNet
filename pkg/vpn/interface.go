//go:build !windows && !darwin && !freebsd
// +build !windows,!darwin,!freebsd

package vpn

import (
	"github.com/mudler/water"
	"github.com/vishvananda/netlink"
)

func createInterface(c *Config) (*water.Interface, error) {
	config := water.Config{
		DeviceType:             c.DeviceType,
		PlatformSpecificParams: water.PlatformSpecificParams{Persist: !c.NetLinkBootstrap},
	}
	config.Name = c.InterfaceName

	return water.New(config)
}

func prepareInterface(c *Config) error {
	link, err := netlink.LinkByName(c.InterfaceName)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(c.InterfaceAddress)
	if err != nil {
		return err
	}

	err = netlink.LinkSetMTU(link, c.InterfaceMTU)
	if err != nil {
		return err
	}

	err = netlink.AddrAdd(link, addr)
	if err != nil {
		return err
	}

	err = netlink.LinkSetUp(link)
	if err != nil {
		return err
	}
	return nil
}
