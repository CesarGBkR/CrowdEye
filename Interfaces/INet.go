package Interfaces

type Network struct {
  Name string
  Mac string 
  Type string
  Mode string
  State int
}

type ScanningInterface struct {
  ID string
  NetName string
  Mac string
  Chann chan struct{}
}

type ResScanningInterface struct {
  ID string
  NetName string
  Mac string
}
