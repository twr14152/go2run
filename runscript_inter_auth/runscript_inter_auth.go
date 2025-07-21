package runscript_inter_auth

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var hostList []string

func loginHosts(hostfile string) {
	hostList = nil
	hf, err := os.Open(hostfile)
	if err != nil {
		log.Fatal("Failed to open host file: ", err)
	}
	defer hf.Close()

	scanner := bufio.NewScanner(hf)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
}

func Connect(user, pass, hostfile string) {
	loginHosts(hostfile)

	auth := ssh.KeyboardInteractive(
		func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			answers := make([]string, len(questions))
			for i := range questions {
				answers[i] = pass
			}
			return answers, nil
		},
	)

	for _, host := range hostList {
		fmt.Printf("\n--- Connecting to %s ---\n", host)

		config := &ssh.ClientConfig{
			User:            user,
			Auth:            []ssh.AuthMethod{auth},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         5 * time.Second,
		}

		conn, err := ssh.Dial("tcp", host, config)
		if err != nil {
			log.Printf("Failed to dial %s: %v\n", host, err)
			continue
		}

		sess, err := conn.NewSession()
		if err != nil {
			log.Printf("Failed to create session for %s: %v\n", host, err)
			conn.Close()
			continue
		}

		modes := ssh.TerminalModes{
			ssh.ECHO:          1,
			ssh.TTY_OP_ISPEED: 14400,
			ssh.TTY_OP_OSPEED: 14400,
		}

		err = sess.RequestPty("xterm", 80, 40, modes)
		if err != nil {
			log.Fatalf("PTY request failed: %v", err)
		}

		stdin, _ := sess.StdinPipe()
		stdout, _ := sess.StdoutPipe()
		sess.Stderr = os.Stderr

		err = sess.Shell()
		if err != nil {
			log.Fatalf("Failed to start shell: %v", err)
		}

		reader := stdout

		// Login prompts
		waitForPrompt(reader, ">")
		fmt.Fprintf(stdin, "enable\n")
		waitForPrompt(reader, "#")

		fmt.Fprintf(stdin, "term len 0\n")
		waitForPrompt(reader, "#")

		// Read commands from per-host file
		cfgFile := "file_" + host + ".cfg"
		fmt.Printf("Sending commands from: %s\n", cfgFile)

		cmdFile, err := os.Open(cfgFile)
		if err != nil {
			log.Printf("Could not open config file %s: %v\n", cfgFile, err)
			sess.Close()
			conn.Close()
			continue
		}

		var commands []string
		scanner := bufio.NewScanner(cmdFile)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				commands = append(commands, line)
			}
		}
		cmdFile.Close()

		for _, cmd := range commands {
			fmt.Printf("[SENDING] %s\n", cmd)
			fmt.Fprintf(stdin, "%s\n", cmd)
			time.Sleep(200 * time.Millisecond)
			output := readUntilPrompt(reader, "#")
			fmt.Println("[OUTPUT]")
			fmt.Println(output)
		//}
			// Cleanly exit config mode
			fmt.Fprintf(stdin, "end\n")
			waitForPrompt(reader, "#")
			fmt.Fprintf(stdin, "exit\n")
			waitForPrompt(reader, ">")
			//fmt.Fprintf(stdin, "exit\n")
			stdin.Close()

		// Ensure session terminates cleanly
		        err = sess.Wait()
		        if err != nil {
				log.Printf("Session ended with error on %s: %v", host, err)
			}
			}

		sess.Close()
		conn.Close()
	}
}
