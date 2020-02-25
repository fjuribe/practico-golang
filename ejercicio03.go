package main

import "fmt";
//Realizar la carga del lado de un cuadrado, mostrar por pantalla el perímetro del mismo (El perímetro de un cuadrado se calcula multiplicando el valor del lado por cuatro)

func main(){
	var lado int;
	fmt.Print("ingrese el lado de un cuadrado:");
	fmt.Scan(&lado);

	fmt.Print("La carga del cuadrado es ",lado);
	fmt.Print("\nEl perimetro es ",(lado*4));



}