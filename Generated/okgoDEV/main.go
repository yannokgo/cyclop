
package main
import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func main() {
		http.Handle("/", http.FileServer(http.Dir("./static")))
        port := os.Getenv("PORT")
        if port == "" {
                        port = "8080"
                        log.Printf("Defaulting to port %s", port)
		}
		fmt.Printf("Listening on port %s", port)
        log.Printf("Listening on port %s", port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                        log.Fatal(err)
        }
}
