package routes



import (
  "time"
)

type request struct{
	URL   string      `json:"url"`
	CustomShort string   `json:"short"`
	Expiry    time.Duration  `json:"expiry"`
}

type resposne struct{
	URL             string    `json:"url"`
	CustomShort     string      `json:"short"`
	Expiry          time.Duration    `json:"short"`
	XRateRemaining   int              `json:"rate_limit"`
	XRateLimitRest   time.Duration     `json:"rate_limit_duration"`
}  