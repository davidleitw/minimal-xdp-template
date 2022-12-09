package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/dropbox/goebpf"
)

func main() {
	bpf := goebpf.NewDefaultEbpfSystem()
	err := bpf.LoadElf("bpf/xdp_main.o")
	if err != nil {
		log.Fatalln(err)
	}
	printXdpProgramInfo(bpf)

	xdp := bpf.GetProgramByName("xdp_root")
	if xdp == nil {
		log.Fatalln("Program 'xdp_main' not found.")
	}

	err = xdp.Load()
	if err != nil {
		log.Fatalln(err)
	}

	err = xdp.Attach("lo")
	if err != nil {
		log.Fatalln(err)
	}
	defer xdp.Detach()

	// Add CTRL+C handler
	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt)

	fmt.Println("XDP program successfully loaded and attached.")
	fmt.Println("Press CTRL+C to stop.")
	fmt.Println()

	for {
		<-ctrlC
		fmt.Println("\nDetaching program and exit")
		return
	}
}

func printXdpProgramInfo(bpfProgram goebpf.System) {
	fmt.Println("Maps:")
	for _, item := range bpfProgram.GetMaps() {
		fmt.Printf("\t%s: %v, Fd %v\n", item.GetName(), item.GetType(), item.GetFd())
	}
	fmt.Println("\nPrograms:")
	for _, prog := range bpfProgram.GetPrograms() {
		fmt.Printf("\t%s: %v, size %d, license \"%s\"\n",
			prog.GetName(), prog.GetType(), prog.GetSize(), prog.GetLicense(),
		)
	}
	fmt.Println()
}
