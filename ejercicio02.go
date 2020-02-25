package main

import "fmt";

func main(){
	var horasTrabajo int;
	var costoHoras float32;
	var sueldo float32;

	fmt.Print("Ingrese las horas trabajadas:");
	fmt.Scan(&horasTrabajo);

	fmt.Print("Ingrese el pago por horas:");
	fmt.Scan(&costoHoras);

	sueldo=float32(horasTrabajo)*costoHoras;

	fmt.Print("El sueldo total del operario es ",sueldo);



}