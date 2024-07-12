package main

import "github.com/sanyanshan/sysh-util/SshUtil"

func main() {
	SshUtil.OpenAndBindPortToLocal(SshUtil.Connector{Host: "sdf", Port: 123, User: "", Password: ""}, "", "")
}
