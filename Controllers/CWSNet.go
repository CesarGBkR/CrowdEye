package Controllers

import (
  "fmt"
  "sync"
  "time"
  "encoding/json"
  "net/http"

  "CrowdEye/Interfaces"
  "github.com/gorilla/websocket"
)

var broadcastI = make(chan []Interfaces.Network) 
var clientsI = make(map[*websocket.Conn]bool)
var mutex = sync.Mutex{}
var firstClient = true 

var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool { return true},
}

func WSGetInterfacesWriter(){

  for {
    Interfaces := <-broadcastI
    data,_ := json.Marshal(Interfaces) 

    mutex.Lock()
    for clientI := range clientsI {
      err := clientI.WriteMessage(websocket.TextMessage, []byte(data)) 
      if err != nil {
        fmt.Printf("Error:%v", err)
      }
    }
    mutex.Unlock()
  }

}

func Foo(){
  for {
    if len(clientsI) > 0 {
      time.Sleep(2 * time.Second)
      data, err := GetInterfaces()
      if err != nil {
        continue
      }
      broadcastI <- data

    }
  }
}

func WSGetInterfaces(r *http.Request, w http.ResponseWriter) {
  
  conn, err := upgrader.Upgrade(w, r, nil) 
  if err != nil {
    fmt.Printf("\nError on Upgrade:\n%v", err)
    return
  } 
  defer func() {
    mutex.Lock()
    delete(clientsI, conn)
    conn.Close()
    mutex.Unlock()
  }()

  mutex.Lock()
  clientsI[conn] = true 
  mutex.Unlock()
  
  if firstClient {
    go Foo()
    firstClient = false 
  }

  for {
    _, _, err := conn.ReadMessage()  // Leer mensajes
    if err != nil {
        break
    }
  }
}
