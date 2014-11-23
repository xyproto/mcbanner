package mcbanner

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGA(t *testing.T) {
	pngbytes, err := ioutil.ReadFile("web/public/img/c1.png")
	if err != nil {
		t.Errorf("%s\n", "Could not read: web/public/img/c1.png")
	}
	FindBest(Likeness, pngbytes)
	//t.Errorf("%s\n", s)
	//const in, out = 4, 2
	//if x := Sqrt(in); x != out {
	//	t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	//}
}

func TestFitness(t *testing.T) {
	sol := NewSolution()
	fmt.Println(sol)
	fmt.Println("fitness:", sol.fitness())
}
