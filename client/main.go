package main

import (
	"log"
	"net/rpc"
)

type Args struct{}

func main() {
	host := "localhost:1122"

	var reply string

	args := Args{}

	client, err := rpc.DialHTTP("tcp", host)
	if err != nil {
		log.Fatal("tentando conexao: ", err)
	}

	err = client.Call("FortuneCookieMessageServer.GetFortuneMessage", args, &reply)
	if err != nil {
		log.Fatal("error", err)
	}

	log.Printf("%s\n", reply)
}
