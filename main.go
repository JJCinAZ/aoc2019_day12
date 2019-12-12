package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Vector struct {
	v [3]int  // x, y, z
	minV [3]int
	maxV [3]int
}

type Moon struct {
	Position Vector
	Velocity Vector
}

const (
	data = `
<x=-5, y=6, z=-11>
<x=-8, y=-4, z=-2>
<x=1, y=16, z=4>
<x=11, y=11, z=-4>`
	test1 = `
<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
)

func main() {
	part1()
}

func part1() {
	moons, _ := GetMoons(strings.NewReader(test1))
	for steps := 1; steps <= 10; steps++ {
		CalcVelocities(moons)
		MoveMoons(moons)
	}
	PrintMoons(moons)
	PrintEnergy(moons)
	PrintRanges(moons)
}

func part2() {
	moons, _ := GetMoons(strings.NewReader(test1))
	for steps := 1; steps <= 1000; steps++ {
		CalcVelocities(moons)
		MoveMoons(moons)
	}
	PrintMoons(moons)
	PrintEnergy(moons)
}

func PrintEnergy(moons []Moon) {
	e := 0
	for l := 0; l < len(moons); l++ {
		p := abs(moons[l].Position.v[0]) + abs(moons[l].Position.v[1]) + abs(moons[l].Position.v[2])
		k := abs(moons[l].Velocity.v[0]) + abs(moons[l].Velocity.v[1]) + abs(moons[l].Velocity.v[2])
		e += p * k
	}
	fmt.Println(e)
}

func PrintMoons(moons []Moon) {
	for l := 0; l < len(moons); l++ {
		fmt.Printf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>\n",
			moons[l].Position.v[0],moons[l].Position.v[1],moons[l].Position.v[2],
			moons[l].Velocity.v[0],moons[l].Velocity.v[1],moons[l].Velocity.v[2])
	}
}

func PrintRanges(moons []Moon) {
	for l := 0; l < len(moons); l++ {
		fmt.Printf("pos=<x=%3d⋯%3d, y=%3d⋯%3d, z=%3d⋯%3d>, vel=<x=%3d⋯%3d, y=%3d⋯%3d, z=%3d⋯%3d>\n",
			moons[l].Position.minV[0], moons[l].Position.maxV[0],
			moons[l].Position.minV[1], moons[l].Position.maxV[1],
			moons[l].Position.minV[2], moons[l].Position.maxV[2],
			moons[l].Velocity.minV[0], moons[l].Velocity.maxV[0],
			moons[l].Velocity.minV[1], moons[l].Velocity.maxV[1],
			moons[l].Velocity.minV[2], moons[l].Velocity.maxV[2])
	}
}

func CalcVelocities(moons []Moon) {
	for l := 0; l < len(moons) - 1; l++ {
		for r := l + 1; r < len(moons); r++ {
			calcVelocity(&moons[l], &moons[r])
		}
	}
}

func calcVelocity(a, b *Moon) {
	for i := 0; i < len(a.Position.v); i++ {
		if a.Position.v[i] < b.Position.v[i] {
			a.Velocity.v[i]++
			b.Velocity.v[i]--
		} else if b.Position.v[i] < a.Position.v[i] {
			b.Velocity.v[i]++
			a.Velocity.v[i]--
		}
	}
}

func MoveMoons(moons []Moon) {
	for l := 0; l < len(moons); l++ {
		a := &moons[l]
		for i := 0; i < 3; i++ {
			a.Position.v[i] += a.Velocity.v[i]
			if a.Position.v[i] < a.Position.minV[i] {
				a.Position.minV[i] = a.Position.v[i]
			} else if a.Position.v[i] > a.Position.maxV[i] {
				a.Position.maxV[i] = a.Position.v[i]
			}
			if a.Velocity.v[i] < a.Velocity.minV[i] {
				a.Velocity.minV[i] = a.Velocity.v[i]
			} else if a.Velocity.v[i] > a.Velocity.maxV[i] {
				a.Velocity.maxV[i] = a.Velocity.v[i]
			}
		}
	}
}

func GetMoons(r io.Reader) ([]Moon, error) {
	re := regexp.MustCompile(`<([xyz])=([-+]?\d+),\s?([xyz])=([-+]?\d+),\s?([xyz])=([-+]?\d+)>`)
	scanner := bufio.NewScanner(r)
	moons := make([]Moon, 0, 4)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		a := re.FindAllStringSubmatch(line, -1)
		for _, m := range a {
			var moon Moon
			for i := 1; i < len(m); i += 2 {
				switch m[i] {
				case "x": moon.Position.v[0], _ = strconv.Atoi(m[i+1])
				case "y": moon.Position.v[1], _ = strconv.Atoi(m[i+1])
				case "z": moon.Position.v[2], _ = strconv.Atoi(m[i+1])
				}
			}
			moons = append(moons, moon)
		}
	}
	return moons, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}