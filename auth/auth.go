package auth

type AuthRequest struct {
	register_no int64
	password string
}

type LoginRequest struct {
	register_no int64
	password string
}

type LogoutRequest struct {
	register_no int64
}
