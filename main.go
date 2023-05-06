package main
import(
	"net/http"
	"os"
	"log"
	"fmt"
)
func getDir()string{
	path,_ := os.Getwd()
	return path
}
func welcome(){
	welcome := `  __ _   ___          ___   ___  _ __ __   __  ___  _ __ 
 / _' | / _ \  _____ / __| / _ \| '__|\ \ / / / _ \| '__|
| (_| || (_) ||_____|\__ \|  __/| |    \ V / |  __/| |   
 \__, | \___/        |___/ \___||_|     \_/   \___||_|   
 |___/ `
 	fmt.Println(welcome)
	}
func getPort() string {
	ct := os.Args
	var target string
	if len(ct) == 1 {
		target = "8000"
		} else if len(ct) != 4{
			fmt.Println("ERROR: BAD PORT NUMBER")
			os.Exit(1)
		}else{
			target = ct[1]
			}
	port := ":" + target
	return port
}
func handler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir(getDir())).ServeHTTP(w, r)
	log.Print("[GET] ",http.StatusOK)
}
func main() {
	welcome()
	port := getPort()
	http.HandleFunc("/", handler)
	log.Print("[START] Server is running on http://localhost:" + port + " ...")
	log.Fatal(http.ListenAndServe(port, nil))
}

