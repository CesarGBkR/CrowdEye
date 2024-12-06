package Orchestrator

import (

  "CrowdEye/Controllers"
  "CrowdEye/Interfaces"
)

func GetInterfaces() ([]Interfaces.Network, error) {
  res, err := Controllers.GetInterfaces()
  if err != nil{
    return res, err
  }
  return res, nil
}

func MonitorMode(Network Interfaces.Network) (Interfaces.Network, error) {
  res, err := Controllers.MonitorMode(Network)
  if err != nil {
    return res, err
  }
  return res, nil
}
