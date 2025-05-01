package threat

import "time"

// internal/threat/engine.go
type Threat struct {
  ID          string    `json:"id"`          // UUID  
  Severity    int       `json:"severity"`    // 1-10  
  Type        string    `json:"type"`        // "TANK", "IED"  
  Coordinates []float64 `json:"coordinates"` // [lat, lon]  
  Timestamp   time.Time `json:"timestamp"`  
}