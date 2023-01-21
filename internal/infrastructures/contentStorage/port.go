package contentStorage

type PastebinContentStorage interface {
	Save(content []byte) (string, error)
	GetPastebinContentStoragePath() string
}
