// internal/infrastructure/kis/websocket.go
package kis

import (
	// project
	"KIS_HTS/internal/config"
	"crypto/tls"
	"net/http"
	"time"

	// standard
	"context"
	"encoding/json"
	"log"

	// external
	"github.com/gorilla/websocket"
)

// WSClient
// 한국투자증권 WebSocket Client
type WSClient struct {
	conn *websocket.Conn
}

// KIS Header 구조
type KISHeader struct {
	TrID string `json:"tr_id"`
}

// KIS Message 구조
type KISMessage struct {
	Header KISHeader `json:"header"`
}

// Connect
// KIS websocket 서버 연결
//
// 필수 Header
//   - approval_key
//   - custtype
//   - tr_type
//   - content-type
func (w *WSClient) Connect(cfg *config.Config) error {
	dialer := websocket.Dialer{
		HandshakeTimeout:  10 * time.Second,
		EnableCompression: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	header := http.Header{}
	header.Set("User-Agent", "KIS_HTS")
	conn, resp, err := dialer.Dial(cfg.KIS.WebsocketURL, header)

	if err != nil {
		if resp != nil {
			log.Println("status:", resp.Status)
			log.Println("headers:", resp.Header)
		}
		return err
	}

	w.conn = conn

	log.Println("websocket connected")

	return nil
}

// Send
// websocket 메시지 전송
func (w *WSClient) Send(data []byte) error {

	if w.conn == nil {
		log.Println("websocket not connected")
		return websocket.ErrBadHandshake
	}

	return w.conn.WriteMessage(websocket.TextMessage, data)
}

// Run
// websocket read loop
// 프로그램 종료까지 지속 수신
func (w *WSClient) Run(ctx context.Context, handler func([]byte)) {

	for {

		select {

		case <-ctx.Done():
			log.Println("websocket shutdown")
			return

		default:

			_, msg, err := w.conn.ReadMessage()
			if err != nil {
				log.Println("websocket read error:", err)
				return
			}

			// JSON 메시지 파싱
			var m KISMessage
			if err := json.Unmarshal(msg, &m); err == nil {

				// pingpong 처리
				if m.Header.TrID == "PINGPONG" {
					log.Println("recv pingpong")
					continue
				}
			}

			handler(msg)
		}
	}
}

// Close
// websocket 연결 종료
func (w *WSClient) Close() {

	if w.conn != nil {
		log.Println("websocket closed")
		w.conn.Close()
	}
}
