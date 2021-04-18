package xapi

type Session struct {
	client *Client
}

type SessionRef string

func (s Session) LoginWithPassword(
	uname, pwd, version, originator string,
) (SessionRef, error) {
	var ref SessionRef
	err := s.client.rpc.Call(
		&ref,
		"session.login_with_password",
		uname,
		pwd,
		version,
		originator,
	)
	return ref, err
}
