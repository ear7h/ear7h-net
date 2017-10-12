package home

var addr string

func Put(s string) {
	addr = s
}

func Get() string {
	return addr
}
