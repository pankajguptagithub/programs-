package main

import(
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"time"
)

var ph = []string{"A","B","C","D","E"}

const hunger = 3
const think = time.Second/10
const eat = time.Second/10

var fmt = log.New(os.Stdout,"",0)
var done = make(chan bool)

type fork byte

func philosopher(phName string, dominantHand chan fork,otherHand chan fork,done chan bool){
	fmt.Println(phName,"seated")
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration){
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h:= hunger; h>0;h--{
		fmt.Println(phName,"hungry")
		<-dominantHand
		<-otherHand
		fmt.Println(phName,"eating")
		rSleep(eat)
		dominantHand <- 'f'
		otherHand <- 'f'
		fmt.Println(phName,"thinking")
		rSleep(think)
		fmt.Println("-------------------")
	}
	fmt.Println(phName,"satisfied")
	done <- true
	fmt.Println(phName,"left the table")
}
func main(){
	fmt.Println("At the beginning,table empty")
	place0:= make(chan fork,1)
	place0 <- 'f'
	placeLeft := place0
	for i:=1;i<len(ph);i++{
		placeRight := make(chan fork,1)
		placeRight <- 'f'
		go philosopher(ph[i],placeLeft,placeRight,done)
		placeLeft = placeRight
	}
	go philosopher(ph[0],place0,placeLeft,done)
	for range ph {
		<-done
	}
	fmt.Println("All done, table empty")
}
