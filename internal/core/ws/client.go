package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type safeConn struct {
	mu   sync.Mutex
	conn *websocket.Conn
}

func (s *safeConn) WriteJSON(v any) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.conn.WriteJSON(v)
}

type Handler struct {
	pool *pgxpool.Pool
}

func NewWSHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 4,
	WriteBufferSize: 1024 * 32,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
