package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	srvc "github.com/kevingimbel/srvc"
)

var port string
var versionFlag bool

var version string
var buildDate string
var commit string

var usage = `USAGE: %s [-port]

Arguments:
- port		Port number, e.g. 1919, 1313, 1337
`

// osSignal captchers signals sent by the OS. This is used to close / exit the program
func osSignal(err chan<- error) {
	osc := make(chan os.Signal)
	signal.Notify(osc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	err <- fmt.Errorf("%s", <-osc)
}

func init() {
	flag.StringVar(&port, "port", "8080", "Assign a port to serve to")
	flag.BoolVar(&versionFlag, "version", false, "Show version")
}

func main() {
	flag.Parse()

	if versionFlag {
		fmt.Printf("Version %s", version)
		if buildDate != "" {
			fmt.Printf("Build date: %s\n", buildDate)
		}

		if commit != "" {
			fmt.Printf("Commit: %s", commit)
		}
		os.Exit(0)
	}

	// If we have args there's something wrong.
	// The only argument is "-port XXXX" which is (removed?) from
	// flag.Args() or doesn't count towards it.
	if len(flag.Args()) > 0 {
		// os.Args[0] is the executable name
		fmt.Printf(usage, os.Args[0])
		os.Exit(1)
	}

	p := ":" + port
	srv := srvc.New(p)
	srv.CreateConfiguredHandlers()

	errch := make(chan error)

	go osSignal(errch)

	go func() {
		fmt.Println("Server is running on port", port)
		errch <- srv.Run()
	}()

	exit := <-errch
	fmt.Println("Stopping Service. Reason:", exit)
}
