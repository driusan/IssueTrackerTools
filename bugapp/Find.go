package bugapp

import (
	"fmt"
	"github.com/driusan/bug/bugs"
	"io/ioutil"
	"os"
)

func find(findType, findValue string) {
	issues, _ := ioutil.ReadDir(string(bugs.GetIssuesDir()))
	for idx, issue := range issues {
		if issue.IsDir() != true {
			continue
		}
		var dir bugs.Directory = bugs.GetIssuesDir() + bugs.Directory(issue.Name())
		b := bugs.Bug{Dir: dir}
		name := getBugName(b, idx)
		var value string
		switch findType {
		case "tags":
		case "status":
			value = b.Status()
		case "priority":
			value = b.Priority()
		case "milestone":
			value = b.Milestone()
		default:
			fmt.Printf("Unknown find type: %s\n", findType)
			return
		}
		if findType == "tags" {
			if b.HasTag(bugs.Tag(findValue)) {
				fmt.Printf("%s: %s\n", name, b.Title(findType))
			}
		} else {
			if value == findValue {
				fmt.Printf("%s: %s\n", name, b.Title(findType))
			}
		}
	}
}

func Find(args ArgumentList) {
	if len(args) < 2 {
		fmt.Printf("Usage: %s find {tag, status, priority, milestone} findvalue\n", os.Args[0])
		return
	}
	switch args[0] {
	case "tags":
		fallthrough
	case "status":
		fallthrough
	case "priority":
		fallthrough
	case "milestone":
		find(args[0], args[1])
	default:
		fmt.Printf("Unknown command: %s %s %s\n", os.Args[0], os.Args[1], os.Args[2])
		return
	}
}
