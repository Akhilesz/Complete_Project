package main

import "math"

func AddServer(port string) {
	//var server Server
	//server.Url = port
	//server.IsAlive = true
	//server.RequestsCount = 0
	server := &Server{
		Url:           port,
		IsAlive:       true,
		RequestsCount: 0,
	}
	Serverpool.Servers = append(Serverpool.Servers, server)
}

func getLeastUsedServer(serverPool *ServerPool) *Server {
	var currCount uint64 = math.MaxUint64
	var bestServer *Server
	for _, server := range serverPool.Servers {
		if server.RequestsCount < currCount {
			bestServer = server
			currCount = server.RequestsCount
		}
	}
	return bestServer
}

func getServerPool() *ServerPool {
	return Serverpool
}
func GetBestServer() *Server {
	return getLeastUsedServer(Serverpool)
}
