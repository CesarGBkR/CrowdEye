package Controllers

import (
  "CrowdEye/Interfaces"
  "github.com/vishvananda/netlink"
)

func GetInterfaces() ([]Interfaces.Network, error) {
  var Networks []Interfaces.Network
  links, err := netlink.LinkList()
  if err != nil {
    return Networks, err
  }
  // Constructor

  for _, link := range links {
    attrs := link.Attrs()
   
    Network := &Interfaces.Network {
      Name: attrs.Name,
      Mac: attrs.HardwareAddr.String(), 
      Type: link.Type(),
    } 
    Networks = append(Networks, *Network)
  }
  return Networks, nil
}

func MonitorMode(interfaceName string) error {
  link, err := netlink.LinkByName(interfaceName)
  if err != nil {
    return err
  }
  if err := netlink.LinkSetDown(link); err != nil {
    return err
  }

  // err = netlink.LinkSetType(link, "monitor")
  // if err != nil {
  //   return err
  // }

  if err := netlink.LinkSetUp(link); err != nil {
    return err
  }
  return nil
}

