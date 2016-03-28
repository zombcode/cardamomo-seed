package routes

import (
  "github.com/zombcode/cardamomo"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

type Server struct {
  Config map[string]interface{}
}

func NewServer() Server {
  configFile, err := ioutil.ReadFile("carfig.json")
  if err != nil {
    // Error
    fmt.Printf("\n\nError reading carfig.json: %s\n\n", err)

    return Server{}
  }

  var config map[string]interface{}
  json.Unmarshal(configFile, &config)

  return Server{Config:config}
}

func Start() {
  // Create new server
  s := NewServer()

  // Instanciate "cardamomo"
  c := cardamomo.Instance(s.Config["port"].(string))

  // Generate routes
  c.Base("/", HomeRoute)

  // Run cardamomo
  c.Run()
}
