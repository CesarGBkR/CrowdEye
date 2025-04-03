package Orchestator

import (

  "CrowdEye/Controllers"
  "CrowdEye/Interfaces"
)

// Net

//CRUD
func CreateOne(network Interfaces.ONet) error {
   err := Controllers.CreateOne(network)
   return err
}

func CreateBulk(networks []Interfaces.ONet) []error {
  err := Controllers.CreateBulk(networks)
  return err
}

// func ReadOne(network Interfaces.ONet)  {
//   Controllers.ReadOne(network)
// }

func ReadAll() ([]Interfaces.ONet, error) {
  res, err := Controllers.ReadAll()
  return res, err
}

func UpdateOne(network Interfaces.ONet) (Interfaces.ONet, error) {
  res, err := Controllers.UpdateOne(network)
  return res, err

}

func UpdateBulk(networks []Interfaces.ONet)  {
 Controllers.UpdateBulk(networks) 
}

func DeleteOne(network Interfaces.ONet) {
  Controllers.DeleteOne(network)
}

func DeleteBulk(networks []Interfaces.ONet)  {
  Controllers.DeleteBulk(networks)
}
