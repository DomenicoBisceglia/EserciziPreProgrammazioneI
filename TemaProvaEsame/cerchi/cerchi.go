package main
import "fmt"
import "os"
import "bufio"
import "math"

type Cerchio struct{
  nome string
  x, y, raggio float64
}
func newCircle(descr string) Cerchio{
  var cerchio Cerchio

  fmt.Sscanf(descr, "%s %f %f %f", &cerchio.nome, &cerchio.x, &cerchio.y, &cerchio.raggio)

  return cerchio
}

func distanzaPunti(x1, x2, y1, y2 float64) float64{

  return math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2))

}

func tocca(cerc1, cerc2 Cerchio) bool{
  /*restituisce true se i cerchi sono tangenti (internamente o esternamente)*/
  if (cerc1.raggio + cerc2.raggio) == distanzaPunti(cerc1.x, cerc2.x, cerc1.y, cerc2.y){
    return true
  }else{
    return false
  }
}

func maggiore(cerc1, cerc2 Cerchio) bool{
  if cerc1.raggio > cerc2.raggio{
    return true
  }else{
    return false
  }
}

func main(){
  var cerchio2 Cerchio
  cerchio2.nome = ""
  cerchio2.x = 0
  cerchio2.y = 0
  cerchio2.raggio = 0
  b := bufio.NewScanner(os.Stdin)
  var To string
  var Mag string
  for b.Scan(){
    descrizione := b.Text()
    cerchio1 := newCircle(descrizione)
    if tocca(cerchio1, cerchio2){
      To = "tangente"
    }else{
      To = "non tangente"
    }
    if maggiore(cerchio1, cerchio2){
      Mag = "maggiore"
    }else{
      Mag = "minore o uguale"
    }
    fmt.Println(cerchio1, To, ",", Mag)

  }
}
