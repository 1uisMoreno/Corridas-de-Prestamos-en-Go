package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type datos struct { //Estructura de los datos que van a recibir
	prodcuto string
	cantidad float64
	plazo    float64
	tipo     string
}

func (d datos) mostrar() { //Funcion para saber que funcion llamar para dependiendo los datos ingresados
	var tasa float64
	prodcuto := strings.ToUpper(d.prodcuto) //El valor escrito del producto se convierte a mayusuclas para evitar que el unas este
	FijoDecre := strings.ToUpper(d.tipo)
	switch prodcuto { //se hizo un switch del tipo de prooducto, este con 3 casos: Efectivo, Bien Mueble y Bien Inmueble
	/*
		Dentro de cada case estaran las condicionales que deben cumplir los datos para que se pueda genear la corrida de manera satisfacoria
	*/
	case "EFECTIVO":
		if d.plazo < 6 { //Condicionales de los meses
			fmt.Println("El plazo debe ser mayor a 6 meses")
		} else if d.plazo > 12 {
			fmt.Println("El plazo debe ser menor a 12 meses")
		} else {
			tasa = .02
			if FijoDecre == "FIJO" { //Condicional del tipo de prestamo
				Fijo(d.cantidad, d.plazo, tasa)
			} else if FijoDecre == "DECRECIENTE" {
				Decreciente(d.cantidad, d.plazo, tasa)
			} else {
				fmt.Println("Los datos ingresados no son correctos")
			}
		}
	case "BIENESMUEBLES":
		if d.plazo < 6 { //Condicionales de los meses
			fmt.Println("El plazo debe ser mayor a 6 meses")
		} else if d.plazo > 24 {
			fmt.Println("El plazo debe ser menor a 24 meses")
		} else {
			tasa = .025
			if FijoDecre == "FIJO" { //Condicional del tipo de prestamo
				Fijo(d.cantidad, d.plazo, tasa)
			} else if FijoDecre == "DECRECIENTE" {
				Decreciente(d.cantidad, d.plazo, tasa)
			} else {
				fmt.Println("Los datos ingresados no son correctos")
			}
		}
	case "BIENESINMUEBLES":
		if d.plazo < 12 { //Condicionales de los meses
			fmt.Println("El plazo debe ser mayor a 12 meses")
		} else if d.plazo > 60 {
			fmt.Println("El plazo debe ser menor a 60 meses")
		} else {
			tasa = .05
			if FijoDecre == "FIJO" { //Condicional del tipo de prestamo
				Fijo(d.cantidad, d.plazo, tasa)
			} else if FijoDecre == "DECRECIENTE" {
				Decreciente2(d.cantidad, d.plazo, tasa)
			} else {
				fmt.Println("Los datos ingresados no son correctos")
			}
		}
	default:
		fmt.Println("Los datos ingresados no son correctos")
	}
}
func Decreciente(cantidad float64, plazo float64, tasa float64) { //Funcion para las cuotas decrecientes
	//Esta funcion se hizo diferente por que se le añade la funcionalidad de que d
	saldoInicial := cantidad
	cuotaMensual := (cantidad*tasa)/1 - math.Pow((1+tasa), plazo) //Formula para
	amortización := cantidad / float64(plazo)
	interes := cantidad * tasa
	saldoFinal := saldoInicial
	var sumaAmortizacion, sumaIntereses, sumaCuotaMensual float64 //Se crean las varables para poner el total de la sumas
	fmt.Println("████████████████████████████████████████████████████████████████████████████")
	fmt.Println("█ No. Periodos | Amortización |  Interes  | Cuota Mensual |  Saldo final   █")
	fmt.Println("█ -------------|--------------|-----------|---------------|--------------- █")
	for i := 0; i <= int(plazo); i++ { //Bucle para la corrida de los datos, se ejecutará las cantidad de meses del prestamo
		if i == 0 {
			fmt.Printf("█      %3d     |       0      |     0     |        0      | %12.2f   █\n", i, saldoFinal)
		} else {
			fmt.Printf("█      %3d     | %10.2f   |%10.2f | %10.2f    | %12.2f   █\n", i, amortización, interes, cuotaMensual, saldoFinal)
		}
		saldoInicial = saldoFinal
		interes = saldoFinal * tasa
		cuotaMensual = amortización + interes
		saldoFinal = saldoInicial - amortización
		sumaAmortizacion = sumaAmortizacion + amortización
		sumaIntereses = sumaIntereses + interes
		sumaCuotaMensual = sumaCuotaMensual + cuotaMensual
	}
	fmt.Println("█ -------------|--------------|-----------|---------------|--------------- █")
	fmt.Printf("█      %s   | %10.2f   |%10.2f | %10.2f    |                █\n", "Total", sumaAmortizacion, sumaIntereses, sumaCuotaMensual)
}
func Decreciente2(cantidad float64, plazo float64, tasa float64) { //Funcion para las cuotas decrecientes de Bienes Inmuebles
	/*
		Esta funcion es diferente la decreciente debido a que le incluye la funcionalidad de que cada año el interes se reduzca un 1%.
		Para eso se crea un varaible de cuota mensual para cada interes.
		Una cuota mensual cuando el interes sea del 4%, otra para cuando sea  3%, etc.
	*/
	saldoInicial := cantidad
	cuotaMensual := (cantidad*tasa)/1 - math.Pow((1+tasa), plazo) //Formula para cuota mensaul
	cuotaMensual2, cuotaMensual3, cuotaMensual4, cuotaMensual5 := cuotaMensual, cuotaMensual, cuotaMensual, cuotaMensual
	amortización := cantidad / float64(plazo)
	interes := cantidad * tasa
	infoInteres := tasa
	saldoFinal := saldoInicial
	interes2, interes3, interes4, interes5 := saldoFinal*(tasa-0.01), saldoFinal*(tasa-0.02), saldoFinal*(tasa-0.03), saldoFinal*(tasa-0.04)
	primerSaldoFinal := saldoInicial
	infoInteres2, infoInteres3, infoInteres4, infoInteres5 := tasa-0.01, tasa-0.02, tasa-0.03, tasa-0.04
	var sumaAmortizacion, sumaIntereses, sumaCuotaMensual float64
	fmt.Println("█████████████████████████████████████████████████████████████████████████████████████████")
	fmt.Println("█ No. Periodos | Amortización |  Interes  | Cuota Mensual |  Saldo final | % de interes █")
	fmt.Println("█ -------------|--------------|-----------|---------------|--------------|------------- █")
	for i := 0; i <= int(plazo); i++ { //Bucle para la corrida de los datos, se ejecutará las cantidad de meses del prestamo
		if i == 0 {
			fmt.Printf("█      %3d     |       0      |     0     |        0      | %12.2f |      %.0f  %%    █\n", i, primerSaldoFinal, infoInteres*100)
		} else if i <= 12 { //Condicion para que cuando sera el primer año, se mantenga el interes del 5%
			fmt.Printf("█      %3d     | %10.2f   |%10.2f | %10.2f    | %12.2f |      %.0f  %%    █\n", i, amortización, interes, cuotaMensual, saldoFinal, infoInteres*100)
		} else if i > 12 && i <= 24 { //Condicion para que cuando sera el segundo año, ya se esten utilizando los datos con 1% menos de interes
			fmt.Printf("█      %3d     | %10.2f   |%10.2f | %10.2f    | %12.2f |      %.0f  %%    █\n", i, amortización, interes2, cuotaMensual2, saldoFinal, infoInteres2*100)
		} else if i > 24 && i <= 36 { //Condicion para que cuando sera el tecer año, ya se esten utilizando los datos con 2% menos de interes
			fmt.Printf("█      %3d     | %10.2f   |%10.2f | %10.2f    | %12.2f |      %.0f  %%    █\n", i, amortización, interes3, cuotaMensual3, saldoFinal, infoInteres3*100)
		} else if i > 36 && i <= 48 { //Condicion para que cuando sera el cuarto año, ya se esten utilizando los datos con 3% menos de interes
			fmt.Printf("█      %3d     | %10.2f   |%10.2f | %10.2f    | %12.2f |      %.0f  %%    █\n", i, amortización, interes4, cuotaMensual4, saldoFinal, infoInteres4*100)
		} else if i > 48 && i <= 60 { //Condicion para que cuando sera el quinto año, ya se esten utilizando los datos con 4% menos de interes
			fmt.Printf("█      %3d     | %10.2f   |%10.2f | %10.2f    | %12.2f |      %.0f  %%    █\n", i, amortización, interes5, cuotaMensual5, saldoFinal, infoInteres5*100)
		}
		saldoInicial = saldoFinal
		interes = saldoFinal * tasa
		interes2 = saldoFinal * (tasa - 0.01)
		interes3 = saldoFinal * (tasa - 0.02)
		interes4 = saldoFinal * (tasa - 0.03)
		interes5 = saldoFinal * (tasa - 0.04)
		cuotaMensual = amortización + interes
		cuotaMensual2 = amortización + interes2
		cuotaMensual3 = amortización + interes3
		cuotaMensual4 = amortización + interes4
		cuotaMensual5 = amortización + interes5
		saldoFinal = saldoInicial - amortización
		sumaAmortizacion = sumaAmortizacion + amortización
		sumaIntereses = sumaIntereses + interes
		sumaCuotaMensual = sumaCuotaMensual + cuotaMensual
	}
	fmt.Println("█ -------------|--------------|-----------|---------------|--------------|------------- █")
	fmt.Printf("█      %s   | %10.2f   |%10.2f | %10.2f    |              |              █\n", "Total", sumaAmortizacion, sumaIntereses, sumaCuotaMensual)
}
func Fijo(cantidad float64, plazo float64, tasa float64) { //Funcion para las cuotas fijas
	saldoInicial := cantidad
	cuotaMensual := cantidad * ((tasa * math.Pow((1+tasa), plazo)) / (math.Pow((1+tasa), plazo) - 1)) //Formula para obtener la cuota mensual fija
	amortización := cuotaMensual
	interes := cantidad * tasa
	saldoFinal := saldoInicial
	var sumaAmortizacion, sumaIntereses, sumaCuotaMensual float64 //Se crean las varables para poner el total de la sumas
	fmt.Println("████████████████████████████████████████████████████████████████████████████")
	fmt.Println("█ No. Periodos | Amortización |  Interes  | Cuota Mensual |  Saldo final   █")
	fmt.Println("█ -------------|--------------|-----------|---------------|--------------- █")
	for i := 0; i <= int(plazo); i++ { //Bucle para la corrida de los datos, se ejecutará las cantidad de meses del prestamo
		if i == 0 {
			fmt.Printf("█      %3d     |       0      |     0     |        0      | %12.2f   █\n", i, saldoFinal) //En el periodo cero lo unico que imprime es el saldo final que es con el se empiza la corrida
		} else {
			fmt.Printf("█      %3d     | %10.2f   |%10.2f | %10.2f    | %12.2f   █\n", i, amortización, interes, cuotaMensual, saldoFinal)

		}
		saldoInicial = saldoFinal             //El saldo inicial siempre se inicia con el saldo final del periodo anterior
		interes = saldoFinal * tasa           //El interes siempre es por la deuda faltante
		amortización = cuotaMensual - interes //En el saldo fijo la amortizacion o pago capital es la cuota mensaul menos los intereses
		saldoFinal = saldoInicial - amortización
		sumaAmortizacion = sumaAmortizacion + amortización
		sumaIntereses = sumaIntereses + interes
		sumaCuotaMensual = sumaCuotaMensual + cuotaMensual
	}
	fmt.Println("█ -------------|--------------|-----------|---------------|--------------- █")
	fmt.Printf("█      %s   | %10.2f   |%10.2f | %10.2f    |                █\n", "Total", sumaAmortizacion, sumaIntereses, sumaCuotaMensual)
}
func main() {
	intCant, _ := strconv.ParseFloat(os.Args[2], 64) //Como los datos que reciben son string, se convierten a flotantes.
	intPlazo, _ := strconv.ParseFloat(os.Args[3], 64)
	datoS := datos{ //Se reciben los datos del strcut mediante la terminal. Se toma desde el valor 1 y no desde el 0 por que ese es la direccion en donde se guarda.
		prodcuto: os.Args[1], //Estos datos no ocupan convertirse, ya recibe el parametro string.
		cantidad: intCant,
		plazo:    intPlazo,
		tipo:     os.Args[4],
	}
	datoS.mostrar()
}
