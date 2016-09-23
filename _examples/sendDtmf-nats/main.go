package main

import (
	"os"
	"sync"
	"time"

	"net/http"

	"golang.org/x/net/context"

	log15 "gopkg.in/inconshreveable/log15.v2"

	"github.com/CyCoreSystems/ari"
	"github.com/CyCoreSystems/ari/client/nc"
)

var log = log15.New()

func main() {

	<-time.After(1 * time.Second)

	if i := run(); i != 0 {
		os.Exit(i)
	}
}

func channelHandler(cl *ari.Client, h *ari.ChannelHandle) {
	log.Info("Running channel handler")

	stateChange := h.Subscribe(ari.Events.ChannelStateChange)
	defer stateChange.Cancel()

	dtmfSub := h.Subscribe(ari.Events.ChannelDtmfReceived)
	defer dtmfSub.Cancel()

	data, err := h.Data()
	if err != nil {
		log.Error("Error getting data", "error", err)
		return
	}
	log.Info("Channel State", "state", data.State)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Info("Waiting for channel events")

		defer wg.Done()

		for {
			select {
			case <-time.After(5 * time.Second):
				log.Error("Timeout waiting for channel UP and all 4 DTMF digits")
				return
			case <-stateChange.Events():
				log.Info("Got state change request")

				data, err = h.Data()
				if err != nil {
					log.Error("Error getting data", "error", err)
					return
				}
				log.Info("New Channel State", "state", data.State)

				if data.State == "Up" {
					stateChange.Cancel() // stop subscription to state change events
					log.Info("Sending DTMF to channel")
					h.SendDTMF("1234", nil)
				}
			case evt := <-dtmfSub.Events():
				dtmf := evt.(*ari.ChannelDtmfReceived)
				log.Info("Got DTMF digit", "digit", dtmf.Digit)
				if dtmf.Digit == "4" {
					return
				}
			}
		}

	}()

	log.Info("Answering call", "error", h.Answer())

	wg.Wait()

	log.Info("Hangup up call", "error", h.Hangup())
	log.Info("Leaving channel handler")
}

func run() int {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup logging

	nc.Logger = log15.New()

	// connect

	cl, err := connect(ctx)
	if err != nil {
		log.Error("Failed to build nc ARI client", "error", err)
		return -1
	}

	// setup app

	log.Info("Starting listener app")

	go listenApp(ctx, cl, channelHandler)

	// start call start listener

	log.Info("Starting HTTP Handler")

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// make call
		log.Info("Make sample call")
		h, err := createCall(cl)
		if err != nil {
			log.Error("Failed to create call", "error", err)
			w.WriteHeader(500)
			w.Write([]byte("Failed to create call: " + err.Error()))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(h.ID()))
	}))

	log.Info("Listening for requests on port 9991")
	http.ListenAndServe(":9991", nil)

	return 0
}

func listenApp(ctx context.Context, cl *ari.Client, handler func(cl *ari.Client, h *ari.ChannelHandle)) {
	sub := cl.Bus.Subscribe("StasisStart")

	for {
		select {
		case e := <-sub.Events():
			log.Info("Got stasis start")
			stasisStartEvent := e.(*ari.StasisStart)
			go handler(cl, cl.Channel.Get(stasisStartEvent.Channel.ID))
		case <-ctx.Done():
			return
		}
	}

}

func createCall(cl *ari.Client) (h *ari.ChannelHandle, err error) {
	h, err = cl.Channel.Create(ari.OriginateRequest{
		Endpoint: "Local/1000",
		App:      "example",
	})

	return
}

func connect(ctx context.Context) (cl *ari.Client, err error) {

	opts := nc.Options{
		URL: "nats://nats:4222",
	}

	log.Info("Connecting")

	cl, err = nc.New(opts)
	return
}
