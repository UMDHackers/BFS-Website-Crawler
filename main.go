package main
import(
 "fmt"
 // "io/ioutil"
 "net/http"
 "html/template"
)
type Results struct {
	url string
}
type Index struct {

}
func indexHandler(w http.ResponseWriter, r *http.Request) {
  i := Index{}
  tem, _ := template.ParseFiles("html/index.html")
  tem.Execute(w, i)
}
func resultsHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("this is the message: ", r.Method)
  if r.Method == "GET" {
    http.Redirect(w, r, "/", 301)
  } else {
    r.ParseForm()
    fmt.Println("this is the url ", r.Form["url"])
  }
}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
      fmt.Fprintf(w, "404! Ya dun goofed!")
    }
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/results", resultsHandler)
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
  http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.ListenAndServe(":8080", nil)

}
