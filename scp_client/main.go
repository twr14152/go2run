package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"log"
	"net"
	"os"

	"github.com/tmc/scp"
)

func getAgent() (agent.Agent, error) {
	agentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	return agent.NewClient(agentConn), err
}

func main() {
	src, err := os.Open("testfile001.txt")
	if err != nil {
		fmt.Println(err)
	}
	src.Close()
	user := "developer"
	pass := "C1sco12345"
	targethost := "sandbox-iosxe-latest-1.cisco.com:22"

	client, err := ssh.Dial("tcp", targethost, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // FIXME: please be more secure in checking host keys
	})
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalln("Failed to create session: " + err.Error())
	}

	dest := src.Name()
	err = scp.CopyPath(src.Name(), dest, session)
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", dest)
	} else {
		fmt.Println("success")
	}
}
