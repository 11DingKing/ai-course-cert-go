package auth

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"sync"
	"time"
)

type Session struct {
	User      domain.User
	ExpiresAt time.Time
	Revoked   bool
}
type Manager struct {
	mu   sync.RWMutex
	data map[string]Session
	ttl  time.Duration
}

func New(ttl time.Duration) *Manager { return &Manager{data: map[string]Session{}, ttl: ttl} }
func (m *Manager) Create(u domain.User) (string, time.Time, error) {
	b := make([]byte, 24)
	if _, e := rand.Read(b); e != nil {
		return "", time.Time{}, e
	}
	t := time.Now().Add(m.ttl)
	m.mu.Lock()
	m.data[hex.EncodeToString(b)] = Session{u, t, false}
	m.mu.Unlock()
	return hex.EncodeToString(b), t, nil
}
func (m *Manager) Get(token string) (domain.User, bool) {
	m.mu.RLock()
	s, ok := m.data[token]
	m.mu.RUnlock()
	if !ok || s.Revoked || time.Now().After(s.ExpiresAt) {
		return domain.User{}, false
	}
	return s.User, true
}
func (m *Manager) Revoke(token string) {
	m.mu.Lock()
	if s, ok := m.data[token]; ok {
		s.Revoked = true
		m.data[token] = s
	}
	m.mu.Unlock()
}
