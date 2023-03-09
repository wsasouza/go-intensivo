package main

import (
	"fmt"
	"time"
)

// **Load Balancer**
func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

// thread 1
func main() {
	ch := make(chan int)
	qtsWorkers := 3

	// inicializa os workers
	for i := 0; i < qtsWorkers; i++ {
		go worker(i, ch)
	}

	// joga a carga para os workers
	for i := 0; i < 100000; i++ {
		ch <- i
	}
}

// Exemplo 1
// **Multithreading em GO**
// go task("Tarefa 1") // Nova thread T2
// go task("Tarefa 2") // Nova thread T3
// task("Tarefa 3") // T1

// func task(name string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(name, ":", i)
// 		time.Sleep(time.Second)
// 	}
// }

// Exemplo 2
// **Canais de comunicação entre threads**
// canal := make(chan string) // canal de comunicação entre as threads

// go func() {
// 	canal <- "Olá mundo"
// 	canal <- "Olá mundo 2"
// }()

// msg := <-canal
// msg2 := <-canal
// fmt.Println("T1:", msg)
// fmt.Println("T2:", msg2)

// Exemplo 3
// **Multithreading em Range**
// 	canal := make(chan int) // canal de comunicação entre as threads
// 	go publish(canal)

// 	go reader(canal)
// 	time.Sleep(time.Second * 5)
// }

// func reader(canal chan int) {
// 	for value := range canal {
// 		fmt.Println("Enviado:", value)
// 	}
// }

// func publish(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }
