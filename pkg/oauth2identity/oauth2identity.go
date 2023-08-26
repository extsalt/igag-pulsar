package oauth2identity

type Oauth2Identity interface {
	GetEmail() (string, error)
	GetName() (string, error)
	GetProfilePicture() (string, error)
}
