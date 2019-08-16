package api

type NFSApi interface {
	Login(username, password string) (cookies string)
}
