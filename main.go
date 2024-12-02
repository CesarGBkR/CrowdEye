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
  network1 := new(Interfaces.ONet)
  Orchestrator.CreateOne(*network1)
  return c.NoContent(204)
}
func CreateBulk(c echo.Context) error {
  var networks []Interfaces.ONet

  sec := new(Interfaces.ONetSec)
  sec.Pass = "Password123"
  sec.Type = 1 // WPA2
  
  network1 := new(Interfaces.ONet)
  network1.SSID = "Wifi1"
  network1.Mac = 1282926237310
  network1.Hidden = false
  network1.NetSec = *sec

  network2 := new(Interfaces.ONet)
  network2.SSID = "Wifi2"
  network2.Mac = 1282926237310
  network2.Hidden = false
  network2.NetSec = *sec
  
  networks = append(networks, *network1, *network2)
    Orchestrator.CreateBulk(networks)
  	return c.NoContent(204)
}
func ReadOne(c echo.Context) error {
  network1 := new(Interfaces.ONet)
    Orchestrator.ReadOne(*network1)
  	return c.NoContent(204)
}
func ReadBulk(c echo.Context) error {
  res, err := Orchestrator.ReadAll()
  if err != nil {
    fmt.Printf("%v", err)
    return c.NoContent(500) 
  }
  c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
  c.Response().WriteHeader(http.StatusOK)
  return json.NewEncoder(c.Response()).Encode(res)
}
func UpdateOne(c echo.Context) error {
  network1 := new(Interfaces.ONet)
    Orchestrator.UpdateOne(*network1)
  	return c.NoContent(204)
}
func UpdateBulk(c echo.Context) error {
  var networks []Interfaces.ONet  
  network1 := new(Interfaces.ONet)
  network2 := new(Interfaces.ONet)
  networks = append(networks, *network1, *network2)
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
  network1 := new(Interfaces.ONet)
  network2 := new(Interfaces.ONet)
  networks = append(networks, *network1, *network2)

  Orchestrator.DeleteBulk(networks)
  return c.NoContent(204)
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

  e.POST("/v1/CreateOne", CreateOne)
  e.POST("/v1/CreateBulk", CreateBulk)
  e.GET("/v1/ReadBulk", ReadBulk)
  // Server Starter
  e.Logger.Fatal(e.Start(":3333"))	
}
