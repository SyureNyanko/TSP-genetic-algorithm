package main

import (
	"fmt"
	"math/rand"
)

const (
	POINTSUM = 10 
	RANDOMSEED = 10
	POINTRANGE = 100
)

type Point struct {
	X int
	Y int
}

type Queue struct {
	line []Point
}



func generateRandomPoint() Point{
	return Point{rand.Intn(POINTRANGE), rand.Intn(POINTRANGE)}
}

func main(){
	var q Queue
	rand.Seed(RANDOMSEED)
	for i := 0 ; i < POINTSUM ; i++ {
		p := generateRandomPoint()
		fmt.Println(p.X, p.Y)
		q.line = append(q.line, p)
	}
	
}