51	  type Handler interface {
52	  	ServeHTTP(ResponseWriter, *Request)
53	  }
54
1261	type HandlerFunc func(ResponseWriter, *Request)
1262	
1263	// ServeHTTP calls f(w, r).
1264	func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
1265		f(w, r)
1266	}