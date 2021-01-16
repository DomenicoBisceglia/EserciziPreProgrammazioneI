package main
import ("fmt"
        "bufio"
        "os"
        "strconv"
        "sort")

func moda(serie []float64)float64{
  mappa := make(map[float64]int)
  for i := 0; i < len(serie); i++{
    mappa[(serie[i])] ++
  }
  max := 0
  var pmax float64
  for i := 0; i < len(serie); i++{
    if mappa[(serie[i])] > max{
      max = mappa[(serie[i])]
      pmax = serie[i]
    }
  }
  return pmax
  }

func mediana(serie []float64) float64{
  if len(serie) % 2 == 0{
    return (serie[(len(serie)/2)- 1] + serie[(len(serie)/2)]) / 2.0
  }else{
    return serie[((len(serie) + 1) / 2) - 1]
  }
}

func main(){
  b := bufio.NewScanner(os.Stdin)
  serie := make([]float64, 0)
  for b.Scan(){
    n, _ := strconv.ParseFloat(b.Text(), 64)
    serie = append(serie, n)
  }
  sort.Float64s(serie)
  mediana := mediana(serie)
  moda := moda(serie)
  fmt.Println(serie)
  fmt.Println(mediana)
  fmt.Println(moda)
}
