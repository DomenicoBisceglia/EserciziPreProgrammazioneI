package main
import ("fmt"
        "os"
        "strconv"
        "bufio")

func main(){
  i, _ := strconv.Atoi(os.Args[1])
  j, _ := strconv.Atoi(os.Args[2])
  file := os.Args[3]
  b, _ := os.Open(file)
  scanner := bufio.NewScanner(b)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    s := scanner.Text()
    s = s[i-1:j]
    fmt.Println(s)
  }
}
