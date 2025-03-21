package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func quizTimer() {
	var segundos int
	if *timerFlag {
		fmt.Println("Cuantos segundos quieres para responder las preguntas?")
		fmt.Scan(&segundos)
	} else {
		segundos = 5
	}
	fmt.Println("Tienes ", segundos, " segundos para responder las preguntas")
	fmt.Println("Presiona enter para comenzar")
	fmt.Scanln()
	timer1 := time.NewTimer(time.Duration(segundos) * time.Second)
	go func() {
		<-timer1.C
		fmt.Println("Se acabó el tiempo")
		os.Exit(0)
	}()

}

var timerFlag = flag.Bool("timer", false, "Añadir flag para activar el timer")

func main() {
	flag.Parse()

	quizTimer()

	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}
	totalPreguntas := len(records)
	fmt.Println("Total de preguntas: ", totalPreguntas)

	preguntas := []string{}
	respuestasCorrectas := []string{}

	totalRespuestasCorrectas := 0

	inputRespuestas := make([]string, totalPreguntas)

	for _, eachrecord := range records {
		preguntas = append(preguntas, eachrecord[0])
		respuestasCorrectas = append(respuestasCorrectas, eachrecord[1])
	}

	for index, pregunta := range preguntas {
		fmt.Println("Pregunta número", index+1, ": ", pregunta)
		fmt.Scan(&inputRespuestas[index])
		if inputRespuestas[index] == respuestasCorrectas[index] {
			fmt.Println("Correcto")
			totalRespuestasCorrectas++
		} else {
			fmt.Println("Incorrecto")
			fmt.Println("La respuesta correcta es: ", respuestasCorrectas[index])
		}
	}

	fmt.Println("Total de respuestas correctas: ", totalRespuestasCorrectas)

	nota := (totalRespuestasCorrectas * 100) / totalPreguntas
	fmt.Println("Nota: ", nota, " sobre 100")

}
