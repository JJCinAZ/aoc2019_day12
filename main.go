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
}

type Moon struct {
	Position Vector
	Velocity Vector
}

const (
	test1 = `
		<x=-1, y=0, z=2>
		<x=2, y=-10, z=-7>
		<x=4, y=-8, z=8>
		<x=3, y=5, z=-1>`
)

func main() {
	moons, _ := GetMoons(strings.NewReader(test1))
	CalcVelocities(moons)
	PrintMoons(moons)
	MoveMoons(moons)
}

func PrintMoons(moons []Moon) {
	for l := 0; l < len(moons) - 1; l++ {
		fmt.Printf("pos=<x=%)
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
	for l := 0; l < len(moons) - 1; l++ {
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