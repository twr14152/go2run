// Code snippet for not standard logins

package main

func main() {
	var user string
	var pass string

	hostData := commands()

	for host, commands := range hostData {
		if host == "nxos:8181" {
			user = "username"
			pass = "password1"
		} else if host == "fastxe:22" || host == "slowxe:8181" {
			user = "username"
			pass = "password2"

		}

	}
}
