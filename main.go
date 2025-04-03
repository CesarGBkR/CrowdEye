package main

import (
  "io"
  "fmt"
	"html/template"
  "net/http"
  "encoding/json"
  "CrowdEye/Orchestator"
  "CrowdEye/Interfaces"
	
  "github.com/labstack/echo/v4"

)

func errResConstructor(c echo.Context, scode int, res error) error {
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(scode)
  return json.NewEncoder(c.Response()).Encode(res)
}

func resConstructor(c echo.Context, scode int, res any) error {
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(scode)
  return json.NewEncoder(c.Response()).Encode(res)
}

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Handlers Definition

func CreateOne(c echo.Context) error {

  var network Interfaces.ONet

  if err := c.Bind(&network); err != nil {
    return errResConstructor(c, 400, nil)
  }
  if err := Orchestator.CreateOne(network); err != nil {

    fmt.Printf("\nError:\n%v", err)
    return errResConstructor(c, 500, err)
  }
  return resConstructor(c, 204, nil )
}

func CreateBulk(c echo.Context) error {

  var networks []Interfaces.ONet
  
  if err := c.Bind(&networks); err != nil {
    return errResConstructor(c, 400, nil)
  }
  if err := Orchestator.CreateBulk(networks); err != nil {

    fmt.Printf("\nError:\n%v", err)
  	return errResConstructor(c, 500, nil)
  }
  return resConstructor(c, 204, nil)
}

func ReadOne(c echo.Context) error {
  // network1 := new(Interfaces.ONet)
    // Orchestator.ReadOne(*network1)
  	return errResConstructor(c, 400, nil)
}

func ReadAll(c echo.Context) error {
  res, err := Orchestator.ReadAll()
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
    return errResConstructor(c, 500, err)  
  }
  return resConstructor(c, 200, res) 
}

func UpdateOne(c echo.Context) error {
  var network Interfaces.ONet
  if err := c.Bind(&network); err != nil {
    return errResConstructor(c, 400, nil)
  }
  Orchestator.UpdateOne(network)
  return resConstructor(c, 500, nil)}

func UpdateBulk(c echo.Context) error {
  var networks []Interfaces.ONet  
  if err := c.Bind(&networks); err != nil {
    return errResConstructor(c, 400, nil)
  }
  Orchestator.UpdateBulk(networks)
  	return resConstructor(c, 500, nil)
}

func DeleteOne(c echo.Context) error {
  network1 := new(Interfaces.ONet)
    Orchestator.DeleteOne(*network1)
  	return resConstructor(c, 500, nil)
}

func DeleteBulk(c echo.Context) error {
  var networks []Interfaces.ONet
  if err := c.Bind(&networks); err != nil {
    return errResConstructor(c, 400, nil)
  }  
  Orchestator.DeleteBulk(networks)
  return resConstructor(c, 500, nil)
}

func GetInterfaces(c echo.Context) error {

  res, err := Orchestator.GetInterfaces() 
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
    return errResConstructor(c, 500, err)
  }
  return resConstructor(c, 200, res)
}

func MonitorMode(c echo.Context) error {
  var network Interfaces.Network
  if err := c.Bind(&network); err != nil {
    return resConstructor(c, 400, err)
  }
  res, err := Orchestator.MonitorMode(network)
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
    return errResConstructor(c, 500, err)
  }
  return resConstructor(c, 200, res)
}

func CreateScann(c echo.Context) error {
  var network Interfaces.Network
  if err := c.Bind(&network); err != nil {
    return resConstructor(c, 400, err)
  }
  if err := Orchestator.CreateScann(network); err != nil {
    fmt.Printf("\nError:\n%v", err)
    return errResConstructor(c, 500, err)
  }
  return resConstructor(c, 204, nil)
}

func StopScann(c echo.Context) error {
  var ScanningInterface Interfaces.ScanningInterface
  if err := c.Bind(&ScanningInterface); err != nil {
    return errResConstructor(c, 400, nil)
  }
  if err := Orchestator.StopScann(ScanningInterface); err != nil {
    fmt.Printf("\nError:\n%v", err)
  }
  return resConstructor(c, 204, nil)
}

func GetScannProcess(c echo.Context) error {
  res, err := Orchestator.GetScannProcess()
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
  }

  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)

}

func GetCurrentNetworks(c echo.Context) error {
  res, err := Orchestator.GetCurrentNetworks()
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
  }

  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)
}


// WS MNGMNT
func WSGetInterfaces(c echo.Context) error {

  r := c.Request()
  w := c.Response().Writer 

  Orchestator.WSGetInterfaces(r, w)

  c.Response().Header().Set("Connection", "upgrade")
  return nil
}

func main() {
  // Render Templates
  e := echo.New()
  
  render := &TemplateRender{
    templates: template.Must(template.ParseGlob("Views/*.html")),
  }
  e.Renderer = render
  e.Static("/static", "Views/static")
  
  // Broadcast 
  go Orchestator.WSGetInterfacesWriter() 

  // Router
	e.GET("/", func(c echo.Context) error {
    return c.Render(200, "index.html", map[string]interface{}{
      "title": "Owo",
    })
  })
  
  // Routes
    //CRUD

  e.POST("/v1/CreateOne", CreateOne)
  e.POST("/v1/CreateBulk", CreateBulk)


  e.GET("/v1/ReadAll", ReadAll)
  
  e.PUT("/v1/Update", UpdateOne)

    // Interface MNGMNT
  e.GET("/v1/GetInterfaces", GetInterfaces)
  e.POST("/v1/MonitorMode", MonitorMode)

  e.POST("/v1/CreateScann", CreateScann)
  e.POST("/v1/StopScann", StopScann)
  e.GET("/v1/GetScannProcess", GetScannProcess)

  e.GET("/v1/GetCurrentNetworks", GetCurrentNetworks)

  
  // WS MNGMNT

  e.GET("/v1/wsGetInterfaces", WSGetInterfaces)
  e.POST("/v1/wsGetInterfaces", WSGetInterfaces)

  // Server Starter
  e.Logger.Fatal(e.Start(":3333"))	
}
