package sanctum

import "time"

type AccessToken interface {
}

type Sanctum interface {
	CreateToken(string, []string, time.Duration) AccessToken
}
