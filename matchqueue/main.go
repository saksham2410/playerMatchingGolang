package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/saksham2410/matchqueue/agent"
	"github.com/saksham2410/matchqueue/matcher"
	"github.com/valyala/fasthttp"
)

var maxTime int
var maxScore int
var scoreGroupLen int
var matchCount int

func init() {
	flag.IntVar(&maxTime, "max_time", 180, "maxtime")
	flag.IntVar(&maxScore, "max_score", 300, "maxscore")
	flag.IntVar(&scoreGroupLen, "score_group_len", 10, "scoreGroupLen")
	flag.IntVar(&matchCount, "match_count", 25, "matchCount")
}

func main() {
	flag.Parse()
	matchingServer := agent.NewHttpMatchingServer(matcher.Time(maxTime), matcher.PlayerScore(maxScore), scoreGroupLen)
	server := fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			if string(ctx.Request.URI().Path()) == "/" {
				ctx.SetStatusCode(200)
				ctx.SetContentType("text/html; charset=utf-8")
				return
			}
			matchingServer.HandleHTTP(ctx)
		},
		Name: "saksham",
	}
	isRun := true

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("Server is shutting down.")
		isRun = false
		if err := server.Shutdown(); err != nil {
			log.Fatal(err)
		}
		log.Println("Server shutdown finished.")
	}()

	go func() {
		log.Println("Matching service started.")
		for isRun {
			matchingServer.Match(matcher.Time(time.Now().Unix()), matchCount)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		log.Println("Sweeping service started.")
		for isRun {
			matchingServer.Sweep(matcher.Time(int(time.Now().Unix()) - maxTime*2))
			time.Sleep(time.Duration(maxTime) * time.Second)
		}
	}()

	addr := ":8000"
	if len(os.Args) >= 2 {
		addr = os.Args[1]
	}
	log.Println("HTTP server listening " + addr)
	if err := server.ListenAndServe(addr); err != nil {
		log.Fatal(err)
	}
	log.Println("Server main exit.")
}
