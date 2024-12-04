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
