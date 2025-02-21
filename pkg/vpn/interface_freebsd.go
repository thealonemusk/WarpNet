package vpn

import (
	"fmt"
	"github.com/mudler/water"
	"os/exec"
)

func createInterface(c *Config) (*water.Interface, error) {
	config := water.Config{
		DeviceType: c.DeviceType,
	}
	config.Name = c.InterfaceName

	return water.New(config)
}

func prepareInterface(c *Config) error {
	err := sh(fmt.Sprintf("ifconfig %s create", c.InterfaceName))
	if err != nil {
		return err
	}
	err = sh(fmt.Sprintf("ifconfig %s inet %s %s netmask %s", c.InterfaceName, c.InterfaceAddress, c.InterfaceAddress, "255.255.255.0"))
	if err != nil {
		return err
	}
	return sh(fmt.Sprintf("ifconfig %s up", c.InterfaceName))
}

func sh(c string) (err error) {
	_, err = exec.Command("/bin/sh", "-c", c).CombinedOutput()
	return
}
