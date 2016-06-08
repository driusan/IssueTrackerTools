package bugapp

import (
	"fmt"
	"github.com/driusan/bug/bugs"
	"os"
	//	"io/ioutil"
)

func findTags(tags []string) {
	bugs := bugs.FindBugsByTag(tags)
	for index, element := range bugs {
		fmt.Printf("%s: %s\n", getBugName(element, index), element.Title(""))
	}
}

func findStatus(status string) {
	bugs := bugs.GetAllBugs()
	count := 0
	for _, b := range bugs {
		if b.Status() == status {
			count++
			fmt.Printf("%s: %s\n", getBugName(b, count), b.Title(""))
		}
	}
}
func Find(args ArgumentList) {
	fmt.Printf("%d\n", len(args))

	for i := range args {
		fmt.Printf("%d: %s\n", i, args[i])
	}

	if len(args) < 2 {
		fmt.Printf("Usage: %s find {(t)ag, (s)tatus, (p)riority, (m)ilestone} searchvalue\n", os.Args[0])
		return
	}

	switch args[0] {
	case "tag", "t":
		fmt.Printf("Running tag command\n")
		tags := []string{args[1]}
		findTags(tags)
	case "status", "s":
		fmt.Printf("Running status command\n")
		findStatus(args[1])
	case "priority", "p":
		fmt.Printf("Running priority command\n")
	case "milestone", "m":
		fmt.Printf("Running milestone command\n")
	default:
		fmt.Printf("Unknown command: %s %s %s\n", os.Args[0], os.Args[1], os.Args[2])
		return
	}
}
