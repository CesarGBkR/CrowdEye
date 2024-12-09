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
    return c.NoContent(400)
  }
  Orchestrator.CreateOne(network)
  return c.NoContent(204)
}

func CreateBulk(c echo.Context) error {

  var networks []Interfaces.ONet
  
  if err := c.Bind(&networks); err != nil {
    return c.NoContent(400)
  }
  Orchestrator.CreateBulk(networks)
  	return c.NoContent(204)
}

func ReadOne(c echo.Context) error {
  // network1 := new(Interfaces.ONet)
    // Orchestrator.ReadOne(*network1)
  	return c.NoContent(204)
}

func ReadAll(c echo.Context) error {
  res, err := Orchestrator.ReadAll()
  if err != nil {
    return c.NoContent(500) 
  }
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)
}

func UpdateOne(c echo.Context) error {
  var network Interfaces.ONet
  if err := c.Bind(&network); err != nil {
    return c.NoContent(400)
  }
  Orchestrator.UpdateOne(network)
  return c.NoContent(204)}

func UpdateBulk(c echo.Context) error {
  var networks []Interfaces.ONet  
  if err := c.Bind(&networks); err != nil {
    return c.NoContent(400)
  }
  Orchestrator.UpdateBulk(networks)
  	return c.NoContent(204)
}

func DeleteOne(c echo.Context) error {
  network1 := new(Interfaces.ONet)
    Orchestrator.DeleteOne(*network1)
  	return c.NoContent(204)
}

func DeleteBulk(c echo.Context) error {
  var networks []Interfaces.ONet
  if err := c.Bind(&networks); err != nil {
    return c.NoContent(400)
  }  
  Orchestrator.DeleteBulk(networks)
  return c.NoContent(204)
}

func GetInterfaces(c echo.Context) error {

  res, err := Orchestrator.GetInterfaces() 
  if err != nil {
    return c.NoContent(500)
  }
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)
}

func MonitorMode(c echo.Context) error {
  var network Interfaces.Network
  if err := c.Bind(&network); err != nil {
    return c.NoContent(400)
  }
  res, err := Orchestrator.MonitorMode(network)
  if err != nil {
    return c.NoContent(500)
  }

  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)

}

func CreateScann(c echo.Context) error {
  var network Interfaces.Network
  if err := c.Bind(&network); err != nil {
    return c.NoContent(400)
  }
  if err := Orchestrator.CreateScann(network); err != nil {
    return c.NoContent(500)
  }
  return c.NoContent(204)
}

func StopScann(c echo.Context) error {
  var ScanningInterface Interfaces.ScanningInterface
  if err := c.Bind(&ScanningInterface); err != nil {
    return c.NoContent(400)
  }
  if err := Orchestrator.StopScann(ScanningInterface); err != nil {
    return c.NoContent(500)
  }
  return c.NoContent(204)
}

func GetScannProcess(c echo.Context) error {
  res, err := Orchestrator.GetScannProcess()
  if err != nil {
    fmt.Printf("\n%v\n", err)
    return c.NoContent(500)
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

  // Server Starter
  e.Logger.Fatal(e.Start(":3333"))	
}
