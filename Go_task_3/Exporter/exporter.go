package Exporter

import(
  "errors"
  "fmt"
  "encoding/json"
  "os"
  "Go_task_3"
)


func JSONExport(sm Go_task_3.SocialMedia, filename string ) error {
  f, err := os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0755)
  if err != nil {
    return errors.New("An error occured opening the file: " + err.Error())
  }
  type Message struct{
    First string
    Second string
  }
  file := sm.Feed()
  m := Message{file[0], file[1]}
  b, err := json.Marshal(m)
  if err != nil {
    return errors.New("An error occured while encodind to JSON file: " + err.Error())
  }
  n, err := f.Write([]byte(b))
  if err != nil {
    return errors.New("An error occured writng to the file: " + err.Error())
  }
  fmt.Printf("Wrote %d bytes\n", n)
  return nil
}
