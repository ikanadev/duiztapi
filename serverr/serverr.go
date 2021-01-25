package serverr

// ServErr represents a server error
type ServErr struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

// Error to implement golang error interface
func (s ServErr) Error() string {
	return s.Message
}

// New creates a new ServErr
func New(code int, message string) ServErr {
	return ServErr{code, message}
}
