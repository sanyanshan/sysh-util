package SshUtil

import "fmt"

// OpenAndBindPortToLocal 打开SSH会话，并绑定远程端口到本地的一个随机端口
func OpenAndBindPortToLocal(connector Connector, remoteHost string, remotePort string) int {
	fmt.Println("go 语言自定义依赖测试")
	fmt.Println(connector, remoteHost, remotePort)
	return 222
}

type Connector struct {
	// 主机地址
	Host string
	// 端口
	Port int
	// 用户名
	User string
	// 密码
	Password string
}
