package SshUtil

// OpenAndBindPortToLocal 打开SSH会话，并绑定远程端口到本地的一个随机端口
func OpenAndBindPortToLocal(connector Connector, remoteHost string, remotePort string) int {
	println("go 语言自定义依赖测试")
	println(connector, remoteHost, remotePort)
	return 222
}

type Connector struct {
	// 主机地址
	host string
	// 端口
	port int
	// 用户名
	user string
	// 密码
	password string
}
