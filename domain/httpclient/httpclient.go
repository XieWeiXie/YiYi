package httpclient

type HttpClient interface {
	Get(url string) ([]byte, error)
	Post(url string) ([]byte, error)
}
