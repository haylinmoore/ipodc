package main

import (
  "encoding/base64"
  "flag"
  "log"
  "github.com/songgao/water"
  "os"
  "bufio"
  "strings"
  "fmt"
)

func createInterface(id int) *water.Interface {
  config := water.Config{
    DeviceType: water.TUN,
  }
  config.Name = fmt.Sprint("IPoDC", id)

  // setup the tun interface
  ifce, err := water.New(config)

  if err != nil {
    if (strings.Contains(fmt.Sprint(err), "device or resource busy")){
      log.Printf(fmt.Sprint(err))
      return createInterface(id+1)
    } else {
      log.Fatal(err)
      os.Exit(3)
    }
  }

  return ifce
}

func main() {

  flag.Parse()

  ifce := createInterface(0)

  // just log the interface name
  log.Printf("if: %s\n", ifce.Name())

  go reader(ifce)

  listener(ifce)
}

func reader(ifce *water.Interface){
  reader := bufio.NewReader(os.Stdin)
  for {
    text, _ := reader.ReadString('\n')
    cleaned := strings.TrimSuffix(text, "\n")
    packet, err := base64.StdEncoding.DecodeString(cleaned)
    log.Printf("Packet recv: % x\n", packet)
    _, err = ifce.Write(packet)
    if err != nil {
      log.Println("failed to write to if")
    }
  }
}

func listener(ifce *water.Interface){

  log.Printf("LISTENING")
  // listen for packets to send to irc
  packet := make([]byte, 2000)
  for {
    n, err := ifce.Read(packet)
    if err != nil {
      log.Fatal(err)
    }
    log.Printf("PACKET %s", base64.StdEncoding.EncodeToString(packet[:n]))
  }
}
