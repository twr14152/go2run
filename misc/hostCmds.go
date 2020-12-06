// This code snippet will allow me to create a map with host:cmds functionality

package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func commands() map[string][]string {
  var count int 
  var host string
  fmt.Print("Number of hosts: ")
  fmt.Scanf("%d", &count)
  hostCmds := make(map[string][]string)
  for i := 1; i <= count; i++ {
    fmt.Print("Hostname: ")
    fmt.Scanf("%s", &host)
    fmt.Println()
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("cmds: ")
    cmds, _ := reader.ReadString('\n')
    s := strings.Split(cmds, ",")
    fmt.Println()
    hostCmds[host] = s 
  }
  for k,v := range hostCmds {
    fmt.Println("hosts is ", k)
    fmt.Println("commands being entered are: ", v)
 }
 return hostCmds
}

func main() {
  hostlist := commands()
  fmt.Printf("%T", hostlist)
}
