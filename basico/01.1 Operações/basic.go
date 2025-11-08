package main
// Rune é apelido para int32
import (
	"fmt"
	"math"
)

func main(){
	var val0 int = 1;
	var val1 int = 2;
	var val2 int = 3;

	// Soma
	fmt.Printf("# %v + %v	\t= %v \n",val0, val1, (val0 + val1) )
	// Subtração
	fmt.Printf("# %v - %v	\t= %v \n",val0, val1, (val0 - val1) )
	// Divisao 
	// E necessario converter em float os operandos da expressao para não perderem as casas decimais 
	fmt.Printf("# %v / %v	\t= %.2f \n",val0, val1, (float64(val0) / float64(val1)) )
	// Multiplicação
	fmt.Printf("# %v * %v	\t= %v \n",val0, val1, (val0 * val1) )
	// modulo
	fmt.Printf("# %v Mod %v\t= %v \n",val0, val1, (val0 % val1) )
	// Exponeciação
	fmt.Printf("# %v^%v	\t= %v \n",val2, val1, math.Pow( float64(val2), float64(val1) ) )
	// Raiz 
	fmt.Printf("# ²√%v	\t= %v \n", val2,  math.Sqrt(float64(val2))  )
	// Deslogamento de bits

	KB := 1 << (1 * 10)  // 1 deslocado 10 bits para esquerda = 2¹⁰ = 1024
	MB := 1 << (2 * 10)  // 1 deslocado 20 bits para esquerda = 2²⁰ = 1.048.576
	GB := 1 << (3 * 10)  // 1 deslocado 30 bits para esquerda = 2³⁰ = 1.073.741.824
	TB := 1 << (4 * 10)  // 1 deslocado 40 bits para esquerda = 2⁴⁰ = 1.099.511.627.776
		
	fmt.Println("Binarios : ")
	fmt.Printf(" KB : %b\n",KB)
	fmt.Printf(" MB : %b\n",MB)
	fmt.Printf(" GB : %b\n",GB)
	fmt.Printf(" TB : %b\n",TB)

}
