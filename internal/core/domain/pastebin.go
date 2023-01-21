package domain

type Pastebin struct {
	Base
	Shortlink                 string `json:"shortlink,omitempty"`
	ExpirationLengthInMinutes uint   `json:"expirationLengthInMinutes,omitempty"`
	Path                      string `json:"path,omitempty"`
}
