package main

import (
      "github.com/tarm/goserial"
      "log"
)

func main() {
  c1 := new(serial.Config)
  c1.Name = "/dev/tty.usbmodem1421"
  c1.Baud = 9600
  
  s, err := serial.OpenPort(c1)
  if err != nil {
    log.Fatal(err)
  }

  n := 0
  buf := make([]byte, 256)
  n, err = s.Read(buf)
  if err != nil {
    log.Fatal(err)
  }
      
  log.Print(n)

  s.Close()
}