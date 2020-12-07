package _6_greedy

import (
	"fmt"
	"testing"
)

func TestScheduling(t *testing.T) {
	deadline := []int {4,2,4,3,1,4,6}
	punish := []int{70,60,50,40,30,20,10}
	total,sch := schedule(deadline,punish)
	fmt.Println("Total punish: ",total)
	fmt.Printf("Solution: %v\n",sch)
	fmt.Println("------------------------------------------------------------")


	for i := range punish {
		punish[i] = 70 - punish[i]
	}
	total,sch = schedule(deadline,punish)
	fmt.Println("Total punish: ",total)
	fmt.Printf("Solution: %v\n",sch)
}
