package health

import "sync"

type HealthStatus struct {
  health bool
  ready bool
  rwLock sync.RWMutex
}

func NewHealthStatus() *HealthStatus {
  return &HealthStatus{
    health: true,
    ready: true,
    rwLock: sync.RWMutex{},
  }
}

func (h *HealthStatus) GetHealth() bool {
  h.rwLock.RLock()
  defer h.rwLock.RUnlock()
  return h.health
}

func (h *HealthStatus) SetHealth(health bool) {
  h.rwLock.Lock()
  defer h.rwLock.Unlock()
  h.health = health
}

func (h *HealthStatus) GetReady() bool {
  h.rwLock.RLock()
  defer h.rwLock.RUnlock()
  return h.ready
}

func (h *HealthStatus) SetReady(ready bool) {
  h.rwLock.Lock()
  defer h.rwLock.Unlock()
  h.ready = ready
}

