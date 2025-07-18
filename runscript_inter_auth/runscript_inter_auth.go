/* DO NOT USE */

package runscript_inter_auth

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"time"
)

var hostList []string
var user string
var pass string
var hostfile string

func loginHosts(hostfile string) {
	hf, err := os.Open(hostfile)
        if err != nil{
		log.Fatal("Failed to Open file: ", err)
	}
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
	fmt.Println(hostList)
}

func Connect(user, pass, hostfile string) {
	interactiveAuth := ssh.KeyboardInteractive(
		func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			answers := make([]string, len(questions))
			for i := range answers {
				answers[i] = pass
			}

			return answers, nil
		},
	)
	loginHosts(hostfile)
	for _, host := range hostList {
		config := &ssh.ClientConfig{
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			User:            user,
			Auth:            []ssh.AuthMethod{interactiveAuth},
		}
		fmt.Println("++++++++++++++++++++++++++++++++")
                fmt.Println("Connected to: ", host)
                fmt.Println("++++++++++++++++++++++++++++++++")		
		conn, err := ssh.Dial("tcp", host, config)
		time.Sleep(1)
		if err != nil {
			log.Fatal("Failed to dial: ", err)
		}
		defer conn.Close()
		sess, err := conn.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}
		fmt.Println("Started session...")
		defer sess.Close()
		fmt.Println("Stating stdin...")
		stdin, err := sess.StdinPipe()
		if err != nil {
                        log.Fatal("Failed to connect to remote devices stdin: ", err)
		}
		sess.Stdout = os.Stdout
		sess.Stderr = os.Stderr
		sess.Shell()
		// cmds file should use host.cfg name standard
		fmt.Println("\n\nThis is the config file named:" + "file_" + host + ".cfg")
		fmt.Printf("\n\n\n\n")
		cmds, err := os.Open("file_" + host + ".cfg")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(cmds)
		scanner.Split(bufio.ScanLines)
		var lines []string
		fmt.Fprintf(stdin, "enable\n")
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		cmds.Close()
		//commands := strings.Join(lines, ";")
		for _, line := range lines {
			fmt.Fprintf(stdin, "%s\n", line)
		}
		//fmt.Fprintf(stdin, line)
		fmt.Fprintf(stdin, "exit\n")
		fmt.Fprintf(stdin, "exit\n")
		stdin.Close()
		sess.Wait()
		sess.Close()
	}
	hostList = nil
}
