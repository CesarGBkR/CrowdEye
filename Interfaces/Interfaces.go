package Interfaces 

type ONetSec struct {
  SecID int
  NetID int 
  Pass string 
  Type int
}

type OGeo struct {
  Lat float32 
  Long float32
}

type ONet struct {
  NetID int
  SSID string 
  Mac int
  Hidden bool
  NetSec ONetSec
  Geo OGeo
}


