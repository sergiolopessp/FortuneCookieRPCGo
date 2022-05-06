package main

import (
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct{}

type FortuneCookieMessageServer string

func (t *FortuneCookieMessageServer) GetFortuneMessage(args *Args, reply *string) (err error) {

	frases := [...]string{
		"Thousands of candles can be lighted from a single candle, and the life of the candle will not be shortened. Happiness never decreases by being shared. - Buddha",
		"There are more things to alarm us than to harm us, and we suffer more often in apprehension than reality. - Seneca",
		"Time you enjoy wasting is not wasted time - Time you enjoy wasting is not wasted time - Marthe Troly-Curtin",
		"When one door of happiness closes, another opens, but often we look so long at the closed door that we do not see the one that has been opened for us. - Helen Keller",
		"Life is not measured by the number of breaths we take, but by the moments that take our breath away. - Maya Angelou ",
		"The pleasure which we most rarely experience gives us greatest delight. - Epictetus",
		"Do not spoil what you have by desiring what you have not; remember that what you now have was once among the things you only hoped for. - Epicurus",
		"Just because it didn’t last forever, doesn’t mean it wasn’t worth your while. - Unknown",
	}
	*reply = frases[rand.Intn(len(frases))]
	return nil
}

func main() {
	fortuneServer := new(FortuneCookieMessageServer)
	rpc.Register(fortuneServer)
	rpc.HandleHTTP()

	port := ":1122"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Erro ao escutar:", err)
	}

	log.Printf("Ouvindo o servidor na porta %s", port)
	http.Serve(listener, nil)
}
