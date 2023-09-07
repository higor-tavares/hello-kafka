package utils

type Headers map[string]string

func RespondWith(
	w http.ResponseWriter,
	status int,
	headers Headers) {
	for k,v := range headers {
		w.Header().Set(k,v)
	}
	w.WriteHeader(status)
}