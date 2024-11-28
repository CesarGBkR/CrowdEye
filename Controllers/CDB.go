package Controllers 

import (
  "fmt"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  
  "CrowdEye/Interfaces"
)

var (
  db *sql.DB
)

func ConnectDB() (*sql.DB, error){
  key := os.Getenv("skey")
  uri := fmt.Sprintf("root:%s@tcp(127.0.0.1:3308)/crowdeye", key)
  var err error
  db, err = sql.Open("mysql", uri)
   
  return db, err
}

func CreateOne(network Interfaces.ONet) error{
  db, err := ConnectDB()
  if err != nil {
    return err
  }
  defer db.Close()
  q := `INSERT INTO networks (ssid, mac, hidden)
  VALUES (?, ?, ?)`
  result, err := db.Exec(q, network.SSID, network.Mac, network.Hidden)
  if err != nil {
    return err
  }

  rowsAffected, err := result.RowsAffected()
  if err != nil {
    return err
  }
  fmt.Printf("\nSuccessful Insert on networks, Affected Rows: %d\n", rowsAffected)
  
  sec := network.NetSec
  LastID, err := result.LastInsertId()
  if err != nil {
    return err
  }
  q = `INSERT INTO  netsecurity (netid, pass, type)
  VALUES (?, ?, ?)`
  result, err = db.Exec(q, LastID, sec.Pass, sec.Type)
  if err != nil {
    return err
  }

  rowsAffected, err = result.RowsAffected()
  if err != nil {
    return err
  }
  fmt.Printf("\nSuccessful Insert on netsecurity, Affected Rows: %d\n", rowsAffected)
  return err
}

func CreateBulk(networks []Interfaces.ONet) []error {
  var errors []error
  for _, network := range networks {
    err := CreateOne(network)
    if err != nil {
      errors = append(errors, err)
    }
  }
  return errors
}

func ReadOne(network Interfaces.ONet)  {
  
}

func ReadAll() ([]Interfaces.ONet, error) {
  var res []Interfaces.ONet
  db, err := ConnectDB()
  if err != nil {
    return res, err
  }
  defer db.Close()
  q := `SELECT * 
FROM networks n
LEFT JOIN netsecurity ns
ON n.netid = ns.netid
WHERE ns.netid > 1 
AND n.netid > 1
ORDER BY n.netid DESC;`
  rows, err := db.Query(q)
  if err != nil {
    return res, err
  }
  defer rows.Close()

  for rows.Next() {
    var (
      netid int
      ssid string
      mac int
      hidden bool
      secid int
      netid2 int
      pass string
      Type  int

    )
    if err := rows.Scan(&netid, &ssid, &mac, &hidden, &secid, &netid2, &pass, &Type); err != nil {
      return res, err
    }

    // Constructor
    nsec:= &Interfaces.ONetSec{
      SecID: secid,
      Pass: pass,
      Type: Type,
    }
    network := &Interfaces.ONet{
      NetID:  netid,
      SSID: ssid,
      Mac: mac,
      Hidden: hidden,
      NetSec: *nsec, 
    }

    // Res Collector
    res = append(res, *network)
    
  }
  if !rows.NextResultSet() {
    db.Close()
    return res, err
	}
  return res, err
}

func UpdateOne(network Interfaces.ONet) {
  
}

func UpdateBulk(network []Interfaces.ONet) {
  
}

func DeleteOne(network Interfaces.ONet) {

}

func DeleteBulk(network []Interfaces.ONet)  {
  
}
