
package main
import ("fmt"
        "os"
        "strconv"
        "bufio")

func cancella (lista []string) []string{
  for i := 0; i < len(lista); i++{
    n, err := strconv.Atoi(lista[i])
    if err == nil {
      if i+n > len(lista){
        lista = lista[:i]
      }else{
        lista = append(lista[:i], lista[i + n + 1:]...)
      }

    }
  }
  return lista
}

func main(){


  if len(os.Args) == 2{
    file := os.Args[1]
    b, err := os.Open(file)
    if err != nil{
      fmt.Println("File non accessibile")
    }else{
      scanner := bufio.NewScanner(b)
      scanner.Split(bufio.ScanWords)

      lista := make([]string, 0)
      for scanner.Scan(){
        s := scanner.Text()
        lista = append(lista, s)
      }
      fmt.Println(lista)
      lista = cancella(lista)
      fmt.Println(lista)
    }
  }else{
    fmt.Println("Fornire 1 nome di file")
  }

}
