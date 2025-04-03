package Orchestator

import (

  "CrowdEye/Controllers"
  "net/http"
)


func WSGetInterfaces(r *http.Request, w http.ResponseWriter) {
  Controllers.WSGetInterfaces(r, w)
}

func WSGetInterfacesWriter() {
 Controllers.WSGetInterfacesWriter()
}

