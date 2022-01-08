package ws

type AuthParams struct {
	EMail    string
	PassWord string
}

type AuthHeader struct {
	AccessToken string
	Client      string
	UID         string
}
