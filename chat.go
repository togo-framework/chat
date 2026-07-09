// Package chat is the Chat BASE plugin for togo: Floating chat widget + a responder registry. The base ships a DB keyword responder; a chat-ai driver swaps in real model answers via the same contract.
//
// Driver plugins self-register a Descriptor via chat.Register in their init();
// an app installs a driver by blank-importing it. This base owns the registry and
// exposes it at GET /api/responders.
package chat

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"

	"github.com/togo-framework/togo"
)

// Descriptor is one registered Chat item contributed by a driver.
type Descriptor struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	DefaultZone string `json:"defaultZone,omitempty"`
	Asset       string `json:"asset,omitempty"` // URL of the driver's built web asset
}

var (
	mu       sync.RWMutex
	registry = map[string]Descriptor{}
)

// Register installs a descriptor. Drivers call this from their init().
func Register(d Descriptor) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := registry[d.Key]; exists {
		return
	}
	registry[d.Key] = d
}

// All returns every registered descriptor, sorted by key.
func All() []Descriptor {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]Descriptor, 0, len(registry))
	for _, d := range registry {
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Key < out[j].Key })
	return out
}

func init() {
	togo.RegisterProviderFunc("chat", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/responders", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{"responders": All()})
		})
		return nil
	})
}
