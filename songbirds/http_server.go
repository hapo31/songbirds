package songbirds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type CheckDetectSwitchAPResponse struct {
	Detected bool   `json:"detected"`
	ESSID    string `json:"essid"`
}

type ConnectSwitchAPRequest struct {
	WpaPassword string `json:"wpa_password"`
}

func HTTPServer(port int, detectSwitchHandler func() (bool, string, error)) (ctx context.Context, wpaPasswordCh chan string, err error) {

	wpaPasswordCh = make(chan string)

	http.HandleFunc("/check_detect_switch_ap", func(w http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("method not allowed"))
			return
		}

		detected, essid, err := detectSwitchHandler()

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("error on scan wifi"))
			return
		}

		response, err := json.Marshal(CheckDetectSwitchAPResponse{Detected: detected, ESSID: essid})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("error on marshal JSON"))
			return
		}
		w.Write(response)
	})

	http.HandleFunc("/connect_switch_ap", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			w.WriteHeader(405)
			w.Write([]byte("method not allowed"))
			return
		}

		bodyBuf := make([]byte, req.ContentLength)
		if _, err := req.Body.Read(bodyBuf); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("bad request"))
			return
		}

		var request ConnectSwitchAPRequest
		if err := json.Unmarshal([]byte(bodyBuf), &request); err != nil {
			w.WriteHeader(400)
			w.Write([]byte("bad request"))
			return
		}

		if request.WpaPassword == "" {
			w.WriteHeader(400)
			w.Write([]byte("bad request"))
			return
		}

		wpaPasswordCh <- request.WpaPassword

		w.WriteHeader(200)
		w.Write([]byte("ok"))

		req.Body.Close()
		return
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)

	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go server.ListenAndServe()

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		cancel()
		stop()
		server.Shutdown(ctx)
	}()

	return
}
