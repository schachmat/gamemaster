package main

import "fmt"
import "os"
import "syscall"

type Rule struct {
	uid int
	gid int
	cmd string
	path string
}

var rules = []Rule{
	{1000, -1, "whoami", "/usr/bin/whoami" },
	{1000, -1, "ifconfig", "/sbin/ifconfig" },
	{1000, -1, "ls", "/bin/ls" },
	{1000, -1, "wifi", "/root/wifi.sh" },
	{1000, -1, "cp", "*"}, // allow to run this program in PATH
	{1000, -1, "*", "*"}, // allow to run any program in PATH
}

func die(exitcode int, msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(exitcode)
}

func main() {
	fmt.Println(os.Args)
	fmt.Println(rules)

	switch {
	case len(os.Args) < 2 || os.Args[1] == "-h":
		die(0, "usage: gm [-hlv] [cmd ..]")
	case os.Args[1] == "-v":
		die(0, "gm v0.1 teichm@in.tum.de")
	case os.Args[1] == "-l":
		for _, r := range rules {
			fmt.Println(r.uid, r.gid, r.cmd, r.path)
		}
		os.Exit(0)
	}

	uid, gid := syscall.Getuid(), syscall.Getgid()
	for _, r := range rules {
		if r.cmd == "*" || r.cmd == os.Args[1] {
			
		}
	}
}
