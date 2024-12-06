package Controllers

import (
  "fmt"
  "os/exec"
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

func UpInterface(link netlink.Link, Network Interfaces.Network) (Interfaces.Network, error) {
  if err := netlink.LinkSetUp(link); err != nil {
    return Network, err
  }  

  Network.State = 1 

  return Network, nil
}

func DownInterface(link netlink.Link, Network Interfaces.Network) (Interfaces.Network, error) {
  if err := netlink.LinkSetDown(link); err != nil {
    return Network, err
  }  

  Network.State = 0

  return Network, nil
}

func MonitorMode(Network Interfaces.Network) (Interfaces.Network, error) {

  link, err := netlink.LinkByName(Network.Name)
  if err != nil {
    return Network, err
  }
  fmt.Printf("%v", link)

  Network, err = DownInterface(link, Network) 
  if err != nil {
    return Network, err
  }
  cmd := exec.Command("sudo", "iw", Network.Name, "set", "type", "monitor") 
  if err := cmd.Run(); err != nil {
    return Network, err
  }

  Network.Mode = "Monitor"

  Network, err = UpInterface(link, Network) 
  if err != nil {
    return Network, err
  }
  return Network, nil
}

