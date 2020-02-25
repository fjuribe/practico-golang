package main

import "fmt";
//Escribir un programa en el cual se ingresen cuatro números, calcular e informar la suma de los dos primeros y el producto del tercero y el cuarto.

func main(){

	var a1,a2,a3,a4,multiplica,suma int;

	fmt.Print("ingrese el valor1:");
	fmt.Scan(&a1);

	fmt.Print("ingrese el valor2:");
	fmt.Scan(&a2);

	fmt.Print("ingrese el valor3:");
	fmt.Scan(&a3);

	fmt.Print("ingrese el valor4:");
	fmt.Scan(&a4);

    suma=a1+a2;
    multiplica=a3*a4;



	fmt.Print("la suma de los valores primeros es ",suma);
	fmt.Print("\nla multiplicacion de los ultimos numeros es",multiplica);
}