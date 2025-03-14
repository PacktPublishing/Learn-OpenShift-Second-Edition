package main 
 
import ( 
  "fmt" 
  "net/http" 
  "time" 
) 
 
func main() { 
  http.HandleFunc("/", HelloServer) 
  http.ListenAndServe("0.0.0.0:8080", nil) 
} 

func HelloServer(w http.ResponseWriter, r *http.Request) { 
  currentTime := time.Now().Format("2006-01-02 15:04:05") 
  fmt.Fprintf(w, "Hello Student! The request time is %s.", currentTime) 
} 
