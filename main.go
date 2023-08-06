package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	//"strings"
)

type Particle struct {
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
	r_c float64
	Id  int `json:"id"`
}

func NewParticle(x float64, y float64, r_c float64, id int) *Particle {
	return &Particle{x, y, r_c, id}
}

func (p *Particle) IsNear(o *Particle) bool {
	// non pac-man space euclidean distance check
	dist := math.Sqrt(math.Pow((p.X-o.X), 2) + math.Pow((p.Y-o.Y), 2))
	return dist <= p.r_c
}

type SpatialHash struct {
	CellSize float64
	Cells    map[SpatialHashKey][]*Particle
}

type SpatialHashKey struct {
	x, y int
}

func NewSpatialHash(cellSize float64) *SpatialHash {
	cells := make(map[SpatialHashKey][]*Particle)
	return &SpatialHash{cellSize, cells}
}

func (h *SpatialHash) KeyForParticle(p *Particle) SpatialHashKey {
	// floor it
	x := int(p.X / h.CellSize)
	y := int(p.Y / h.CellSize)
	return SpatialHashKey{x, y}
}

func (h *SpatialHash) Add(p *Particle) {
	key := h.KeyForParticle(p)
	h.Cells[key] = append(h.Cells[key], p)
}

func (h *SpatialHash) FindNeighbors(p *Particle) []Particle {
	var res []Particle
	key := h.KeyForParticle(p)
	// check adjecent cells
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			k := SpatialHashKey{key.x + dx, key.y + dy}
			// h.Cells[k] all particles in neighborspace

			for i := 0; i < len(h.Cells[k]); i++ {
				posP := h.Cells[k][i]
				if posP.Id != p.Id && p.IsNear(posP) {
					res = append(res, *posP)
				}
			}
		}
	}
	return res
}

func PrintableListOfNearParticles(o []Particle) []string {
	// var out []string
	out := make([]string, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = fmt.Sprintf("%d is near", o[i].Id)
	}

	return out
}

func randomF(min, max float64) float64 {
	return (min + rand.Float64()*(max-min))
}

func main() {
	// TODO: read from params file instead
	if len(os.Args) != 5 {
		fmt.Printf("\nExpected 4 parameters: N, L, M, r_c\n")
		os.Exit(1)
	}

	N, _ := strconv.Atoi(os.Args[1])             // number of particles
	L, _ := strconv.ParseFloat(os.Args[2], 32)   // size of cell
	M, _ := strconv.Atoi(os.Args[3])             // number of cells (MxM)
	r_c, _ := strconv.ParseFloat(os.Args[4], 32) // interaction radius
	fmt.Printf("%d %f %d %f\n", N, L, M, r_c)    // sanity check

	size := L * float64(M)
	sh := NewSpatialHash(L)
	var list []*Particle // keep a list of pointers for ease of use...
	for i := 0; i < N; i++ {
		// TODO: read from initial condifitions file instead of random placement
		p := NewParticle(randomF(0, size), randomF(0, size), r_c, i)
		list = append(list, p)
		sh.Add(p)
	}

	// TODO: print state in sh
	fmt.Println("particles:")
	listB, _ := json.Marshal(list)
	fmt.Println(string(listB))

	// TODO: check time/performance
	// TODO: print to a file output
	out := make([][]Particle, N)
	for i := 0; i < N; i++ {
		p := list[i]
		near := sh.FindNeighbors(p)
		// fmt.Printf("Particle %d at (%f, %f) w/ r:%f\nNeighbors:\n", p.Id, p.X, p.Y, p.r_c)
		// fmt.Println(strings.Join(PrintableListOfParticles(near), "\n"))
		out = append(out, near)
	}

	// TODO: use the list of particles and do a naive near search, for comparison

	fmt.Println("output:")
	outB, _ := json.Marshal(out)
	fmt.Println(string(outB))
}
