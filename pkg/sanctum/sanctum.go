package sanctum

import "time"

type Authenticate interface {
}

type AccessToken interface {
}

type Sanctum interface {
	CreateToken(string, []string, time.Duration) AccessToken
}
