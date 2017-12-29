package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	particleLines := strings.Split(string(input), "\n")

	particles := []Particle{}

	minAcceleration := int64(math.MaxInt64)
	minVelocity := int64(math.MaxInt64)
	minPosition := int64(math.MaxInt64)
	var minIndex int

	for i, pLine := range particleLines {
		parts := strings.Split(pLine, ">, ")

		particle := Particle{}

		for _, p := range parts {
			v := parseVector(p)

			switch {
			case strings.Contains(p, "p"):
				particle.position = v
			case strings.Contains(p, "v"):
				particle.velocity = v
			case strings.Contains(p, "a"):
				particle.acceleration = v
			}
		}

		particles = append(particles, particle)

		acceleration := particle.acceleration.Magnitude()
		velocity := particle.velocity.Magnitude()
		position := particle.position.Magnitude()

		var smaller bool

		if acceleration == minAcceleration {
			if velocity == minVelocity {
				if position < minPosition {
					smaller = true
				}
			}

			if velocity < minVelocity {
				smaller = true
			}
		}

		if acceleration < minAcceleration {
			smaller = true
		}

		if smaller {
			minAcceleration = acceleration
			minVelocity = velocity
			minPosition = position
			minIndex = i
		}
	}

	log.Println("Part one answer is", minIndex, "(", minAcceleration, minVelocity, minPosition, ")")

	// Run the simulation for a bit and see if we get any collisions
	for t := 0; t < 1000; t++ {
		collisionLocs := make(map[string]struct{})
		locations := make(map[string][]int)

		// Find locations for all particles
		for index, particle := range particles {
			location := particle.Location(int64(t))

			if _, ok := locations[location.String()]; ok {
				collisionLocs[location.String()] = struct{}{}
			}

			locations[location.String()] = append(locations[location.String()], index)
		}

		collisions := make(map[int]struct{})
		for v := range collisionLocs {
			for _, index := range locations[v] {
				collisions[index] = struct{}{}
			}
		}

		newParticles := []Particle{}
		for index, particle := range particles {
			if _, ok := collisions[index]; !ok {
				newParticles = append(newParticles, particle)
			}
		}

		particles = newParticles
	}

	log.Println("Part two answer is", len(particles))
}

func (p *Particle) String() string {
	return fmt.Sprintf("p=%s, v=%s, a=%s", p.position.String(), p.velocity.String(), p.acceleration.String())
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%d,%d,%d>", v.x, v.y, v.z)
}

func (p *Particle) Tick() {
	p.velocity.x += p.acceleration.x
	p.velocity.y += p.velocity.y
	p.velocity.z += p.velocity.z

	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.position.z += p.velocity.z
}

func (p *Particle) Location(time int64) Vector {
	x := p.position.x + (time * p.velocity.x) + (time * (time + 1) * p.acceleration.x / 2)
	y := p.position.y + (time * p.velocity.y) + (time * (time + 1) * p.acceleration.y / 2)
	z := p.position.z + (time * p.velocity.z) + (time * (time + 1) * p.acceleration.z / 2)

	return Vector{
		x: x,
		y: y,
		z: z,
	}
}

func (v *Vector) Equals(x Vector) bool {
	return v.x == x.x && v.y == x.y && v.z == x.z
}

func parseVector(input string) Vector {
	parts := strings.Split(input, "=<")

	cartesianComponents := strings.Split(strings.TrimSpace(strings.TrimSuffix(parts[1], ">")), ",")

	x, _ := strconv.Atoi(cartesianComponents[0])
	y, _ := strconv.Atoi(cartesianComponents[1])
	z, _ := strconv.Atoi(cartesianComponents[2])

	return Vector{
		x: int64(x),
		y: int64(y),
		z: int64(z),
	}
}

type Particle struct {
	position     Vector
	velocity     Vector
	acceleration Vector
}

type Vector struct {
	x int64
	y int64
	z int64
}

func (v *Vector) Magnitude() int64 {
	return int64(math.Pow(float64(v.x), 2) + math.Pow(float64(v.y), 2) + math.Pow(float64(v.z), 2))
}
