package main

import (
	"checks"
	"flag"
	"fmt"
	"notify"
	"os"
)

var toAddr = flag.String("to-address", "jareksedlacek@gmail.com", "Email address to send to")
var fromAddr = flag.String("from-address", "updatechecks@jsedlacek.info", "Email address to send from")
var host = flag.String("email-host", "localhost", "Host to send email from")
var port = flag.String("email-port", "25", "Port to connect to host on")

func main() {
	outOfDatePrograms := checks.GetOutOfDatePrograms()
	message := notify.BuildOutOfDateMessage(outOfDatePrograms)
	if message != "" {
		fmt.Printf("%v", message)
		err := notify.SendMail(*toAddr, *fromAddr, *host+":"+*port, "Program updates found!", message)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error sending mail: %v\n", err)
		}
	} else {
		fmt.Printf("No updates found\n")
	}
}
