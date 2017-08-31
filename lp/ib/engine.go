package ib

import (
	"flag"

	"os"

	agent "github.com/gofinance/ib"
	log "github.com/sirupsen/logrus"
)

var (
	gwURL = flag.String("gw", "", "Gateway URL")
)

func getGatewayURL() string {
	if *gwURL != "" {
		return *gwURL
	}
	if url := os.Getenv("GATEWAY_URL"); url != "" {
		return url
	}
	return "127.0.0.1:4001"
}

func newEngine(url string) *agent.Engine {
	opts := agent.EngineOptions{Gateway: url}
	if os.Getenv("CI") != "" || os.Getenv("IB_ENGINE_DUMP") != "" {
		opts.DumpConversation = true
	}

	engine, err := agent.NewEngine(opts)
	if err != nil {
		log.Fatalf("cannot connect engine: %s", err)
	}

	if engine.State() != agent.EngineReady {
		log.Fatalf("engine %s not ready (did a prior test Stop() rather than ConditionalStop() ?)", engine.ConnectionInfo())
	}

	log.Printf("engine %s; state: %v", engine.ConnectionInfo(), engine.State())
	return engine
}

/*
func (e *Engine) expect(t *testing.T, seconds int, ch chan Reply, expected []IncomingMessageID) (Reply, error) {
	for {
		select {
		case <-time.After(time.Duration(seconds) * time.Second):
			return nil, errors.New("Timeout waiting")
		case v := <-ch:
			if v.code() == 0 {
				t.Fatalf("don't know message '%v'", v)
			}
			for _, code := range expected {
				if v.code() == code {
					return v, nil
				}
			}
			// wrong message received
			t.Logf("received message '%v' of type '%v'\n",
				v, reflect.ValueOf(v).Type())
		}
	}
}
*/
