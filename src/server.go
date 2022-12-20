package main

// 启动 HTTP 服务
func StartHttpServer() {
	// 创建应用
	app := NewApp()

	// 启动服务
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
