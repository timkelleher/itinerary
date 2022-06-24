package backend

var methods = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"HEAD",
	"PATCH",
	"CONNECT",
	"OPTIONS",
	"TRACE",
}

func Methods() []string {
	return methods
}
