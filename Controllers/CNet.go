package Controllers

import (
  "fmt"
  "time"
  "os/exec"
  // "sync"
  "errors"
  "github.com/google/uuid"
  

  "CrowdEye/Interfaces"

  "github.com/vishvananda/netlink"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
  "github.com/google/gopacket/layers"
)

var ScannProcess []Interfaces.ScanningInterface 
var CurrentNetworks []Interfaces.Network

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

func PacketCapturer(Network Interfaces.Network, stopChan <-chan struct{}) error {
  // Define time of promiscuous mode
  wt := time.Duration(5 * float64(time.Second))

  handle, err := pcap.OpenLive(Network.Name, 1024, true, wt)
  if err != nil {
    return err
  }

  defer handle.Close()

  filter := "type mgt subtype beacon or subtype probe-req or subtype probe-resp"// Capture only Beacon packets
	if err := handle.SetBPFFilter(filter); err != nil {
		 return err
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())


  for packet := range packetSource.Packets() {
    if packet != nil {
      dot11Layer := packet.Layer(layers.LayerTypeDot11InformationElement)
      if dot11Layer == nil {
        continue
      }
      // Conver in to *layers.Dot11InformationElement
      dot11info, ok := dot11Layer.(*layers.Dot11InformationElement)
      if !ok {
          continue
      }
      if dot11info.ID == layers.Dot11InformationElementIDSSID {
          fmt.Printf("SSID: %q\n", dot11info.Info)
      }

    }
  }
  for {
      select {
      case <- stopChan :
          fmt.Printf("Capture Stopped.")
          return nil
      }
    }
}

func CreateScann(Network Interfaces.Network) (error) {

  
  // Monitor mode validation
  if Network.Mode != "Monitor" {
    err := errors.New("Network not in Monitor Mode")
    fmt.Printf("\n%v\n", err)
    return err 
  }

  stopChan := make(chan struct{})

  go PacketCapturer(Network, stopChan) 


  // Scanning Process Identifyer Constructor
  ScannIdentify := &Interfaces.ScanningInterface{
    ID: uuid.New().String(),
    NetName: Network.Name,
    Mac: Network.Mac,
    Chann: stopChan,
  } 

  ScannProcess = append(ScannProcess, *ScannIdentify)
  return nil 
  
}

func StopScann(ScanningInterface Interfaces.ScanningInterface) error {
  err := errors.New("Process Not Found")
  if len(ScannProcess) > 0 {
    for _, Proc := range ScannProcess {
      if Proc.ID == ScanningInterface.ID {
        close(Proc.Chann)
        return nil
      }
    }
    return err
  }
  return err
}

func GetScannProcess() ([]Interfaces.ResScanningInterface, error) {

  var res []Interfaces.ResScanningInterface

  err := errors.New("Process Not Found")

  if len(ScannProcess) > 0 {
    for _, Proc := range ScannProcess{
      resProc := &Interfaces.ResScanningInterface{
      ID: Proc.ID,
      NetName: Proc.NetName,
      Mac: Proc.Mac,
      }
      res = append(res, *resProc)
    }
    return res, nil
  }

  return res, err
}


// TODO
// func SaveScann(){
//   writer, err := pcap.OpenOffline("output.pcap")
//   if err != nil {
//       log.Fatalf("Error al crear archivo PCAP: %v", err)
//   }
//   defer writer.Close()
//
//   for packet := range packetSource.Packets() {
//       writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
//   }
// }
