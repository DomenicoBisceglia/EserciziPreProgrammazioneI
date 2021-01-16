package main
import ("fmt"
        "os")
/*
Scrivere una funzione ricorsiva che, dati un byte c ed una stringa s,
restituisce la posizione della prima occorrenza di c in s, e -1 se c non occorre in s.
E un main che legge da linea di comando i dati e testi la funzione
*/
func main(){

  c:= os.Args[1]
  s:= os.Args[2]
  b := c[0]
  ris := primaOccorrenza(0, b, s)
  if ris == -1{
    fmt.Println("Non c'è nessuna occorrenza di", c, "in", s)
  }else{
    fmt.Println(c, "si trova nella posizione", ris, "di", s)
  }

}

func primaOccorrenza(pos int, c byte, s string)int {
  if pos < len(s){
    if c == s[pos]{
      return pos
    }else{
      return primaOccorrenza(pos + 1, c, s)
    }

  }else{
    return -1
  }

}
