package Controllers

import (
  "fmt"

  "github.com/vishvananda/netlink"
)

func getInterfaces() error {
  links, err := netlink.LinkList()
  if err != nil {
    return err
  }
  fmt.Prinln("Interfaces: ")
  for _, link := range links {
    fmt.Printf("\n Name: %s    Tipo: %s", link.Attrs().Name, link.Type())
  }
  return nil
}

func monitorMode(interfaceName string) error {
  link, err := netlink.LinkByname(interfaceName)
  if err != nil {
    return err
  }
  if err := netlink.LinkSetDown(link) err != nil {
    return err
  }

  err = netlink.LinkSetType(link, "monitor")
  if err != nil {
    return err
  }

  if err := netlink.LinkSetUo(link) err != nil {
    return err
  }
  return nil
}

