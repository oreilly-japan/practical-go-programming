package ok

// valid-json-marshal
type user struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	// "-" とすることでエンコードの対象から除いておく
	X func() `json:"-"`
}

// valid-json-marshal
