package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	args := os.Args
	if len(args) != 3 {
		log.Fatal("Usage: customer_generator <experimentToken> <numberOfCustomers>")
	}
	numberOfCustomers, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal("Incorrect nuber of customers")
	}
	for i := 0; i < numberOfCustomers; i++ {
		wg.Add(1)
		go request(args[1], strconv.Itoa(i), &wg)
	}
	wg.Wait()
}

func request(experimentToken, i string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get("http://localhost:8080/assignment/" + experimentToken + "/test" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + i)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Request " + i + " send")
}
