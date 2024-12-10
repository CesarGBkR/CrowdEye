package main

import (
  "io"
  "fmt"
	"html/template"
  "net/http"
  "encoding/json"
  "CrowdEye/Orchestrator"
  "CrowdEye/Interfaces"
	
  "github.com/labstack/echo/v4"

)

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
    return resConstructor(c, 400, nil)
  }
  if err := Orchestrator.CreateOne(network); err != nil {

    fmt.Printf("\nError:\n%v", err)
    return resConstructor(c, 500, err)
  }
  return resConstructor(c, 204, nil )
}

func CreateBulk(c echo.Context) error {

  var networks []Interfaces.ONet
  
  if err := c.Bind(&networks); err != nil {
    return resConstructor(c, 400, nil)
  }
  if err := Orchestrator.CreateBulk(networks); err != nil {

    fmt.Printf("\nError:\n%v", err)
  	return resConstructor(c, 500, err)
  }
  return resConstructor(c, 204, nil)
}

func ReadOne(c echo.Context) error {
  // network1 := new(Interfaces.ONet)
    // Orchestrator.ReadOne(*network1)
  	return resConstructor(c, 400, nil)
}

func ReadAll(c echo.Context) error {
  res, err := Orchestrator.ReadAll()
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
    return resConstructor(c, 500, err)  
  }
  return resConstructor(c, 200, res) 
}

func UpdateOne(c echo.Context) error {
  var network Interfaces.ONet
  if err := c.Bind(&network); err != nil {
    return resConstructor(c, 400, nil)
  }
  Orchestrator.UpdateOne(network)
  return resConstructor(c, 500, nil)}

func UpdateBulk(c echo.Context) error {
  var networks []Interfaces.ONet  
  if err := c.Bind(&networks); err != nil {
    return resConstructor(c, 400, nil)
  }
  Orchestrator.UpdateBulk(networks)
  	return resConstructor(c, 500, nil)
}

func DeleteOne(c echo.Context) error {
  network1 := new(Interfaces.ONet)
    Orchestrator.DeleteOne(*network1)
  	return resConstructor(c, 500, nil)
}

func DeleteBulk(c echo.Context) error {
  var networks []Interfaces.ONet
  if err := c.Bind(&networks); err != nil {
    return resConstructor(c, 400, nil)
  }  
  Orchestrator.DeleteBulk(networks)
  return resConstructor(c, 500, nil)
}

func GetInterfaces(c echo.Context) error {

  res, err := Orchestrator.GetInterfaces() 
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
    return resConstructor(c, 500, err)
  }
  return resConstructor(c, 200, res)
}

func MonitorMode(c echo.Context) error {
  var network Interfaces.Network
  if err := c.Bind(&network); err != nil {
    return resConstructor(c, 400, err)
  }
  res, err := Orchestrator.MonitorMode(network)
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
    return resConstructor(c, 500, err)
  }
  return resConstructor(c, 200, res)
}

func CreateScann(c echo.Context) error {
  var network Interfaces.Network
  if err := c.Bind(&network); err != nil {
    return resConstructor(c, 400, err)
  }
  if err := Orchestrator.CreateScann(network); err != nil {
    fmt.Printf("\nError:\n%v", err)
    return resConstructor(c, 500, err)
  }
  return resConstructor(c, 204, nil)
}

func StopScann(c echo.Context) error {
  var ScanningInterface Interfaces.ScanningInterface
  if err := c.Bind(&ScanningInterface); err != nil {
    return resConstructor(c, 400, nil)
  }
  if err := Orchestrator.StopScann(ScanningInterface); err != nil {
    fmt.Printf("\nError:\n%v", err)
  }
  return resConstructor(c, 204, nil)
}

func GetScannProcess(c echo.Context) error {
  res, err := Orchestrator.GetScannProcess()
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
  }

  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)

}

func GetCurrentNetworks(c echo.Context) error {
  res, err := Orchestrator.GetCurrentNetworks()
  if err != nil {
    fmt.Printf("\nError:\n%v", err)
  }

  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)
}

func main() {
  // Foo
  
  // Render Templates
  e := echo.New()
  
  render := &TemplateRender{
    templates: template.Must(template.ParseGlob("Views/*.html")),
  }
  e.Renderer = render
  e.Static("/static", "Views/static")

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
  // Server Starter
  e.Logger.Fatal(e.Start(":3333"))	
}
