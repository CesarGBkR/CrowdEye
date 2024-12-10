package Orchestrator

import (

  "CrowdEye/Controllers"
  "CrowdEye/Interfaces"
)

func GetInterfaces() ([]Interfaces.Network, error) {
  res, err := Controllers.GetInterfaces()
  return res, err 
}

func MonitorMode(Network Interfaces.Network) (Interfaces.Network, error) {
  res, err := Controllers.MonitorMode(Network)
  return res, err 
}

func CreateScann(Network Interfaces.Network) error {
  err := Controllers.CreateScann(Network)
  return err
}

func StopScann(ScannInterface Interfaces.ScanningInterface) error { 
  if err := Controllers.StopScann(ScannInterface); err != nil { 
    return err
  } 
  return nil 
}

func GetScannProcess() ([]Interfaces.ResScanningInterface, error) {
  res, err := Controllers.GetScannProcess()
  return res, err
}

func GetCurrentNetworks() ([]Interfaces.Network, error) {
  res, err := Controllers.GetCurrentNetworks()
  return res, err
}
