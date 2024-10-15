package dto

type Request struct {
	Body string `json:"InputString"`
}

type Response struct {
	Body string `json:"OutputString"`
}
