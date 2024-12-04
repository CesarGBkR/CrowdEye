package Interfaces 

type ONetSec struct {
  SecID int `json:"SecID"`
  NetID int `json:"NetID"`
  Pass string `json:"Pass"`
  Type int `json:"Type" validate:"required"`
}

type OGeo struct {
  Lat float32 `json:"Lat" validate:"required"`
  Long float32 `json:"Long" validate:"required"`
}

type ONet struct {
  NetID int `json:"NetID"`
  SSID string `json:"SSID"`
  Mac int `json:"Mac" validate:"required"`
  Hidden bool `json:"Hidden" validate:"required"`
  NetSec ONetSec `json:"NetSec" validate:"required"`
  Geo OGeo `json:"Geo" validate:"required"`
}


