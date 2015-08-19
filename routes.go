package main

// START OMIT
http.HandleFunc("/hello", handleHello)
http.HandleFunc("/goodbye", handleGoodbye)
http.HandleFunc("/", handleIndex)
// END OMIT
