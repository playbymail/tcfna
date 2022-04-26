/*
 * TCFNA - Game Engine for SPI's Campaign for North Africa
 * Copyright (c) 2022 Michael D Henderson
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package memory

import (
	"fmt"
	"log"
	"math"
	"time"
)

func (ds *STORE) BoardAsSVG() *svg {
	start := time.Now()

	// find the min and max values for rows and columns
	_, maxRow, _, maxCol := ds.board.Sorted.Bounds()

	radius := 30.0
	offset := (math.Sqrt(3) * radius) / 2
	maxX := 40.0 + offset*float64(maxCol*2)
	maxY := 40.0 + offset*float64(maxRow)*math.Sqrt(3)

	s := &svg{}
	s.id = "s"
	s.viewBox.minX = 0
	s.viewBox.minY = 0
	s.viewBox.width = int(maxX)
	s.viewBox.height = int(maxY)

	// the board has 0,0 in the lower left.
	// svg has 0,0 in the upper left.
	// we have to change y from [0..maxY] to [maxY..0].
	for _, hex := range ds.board.Sorted {
		x := 40.0 + offset*float64(hex.Column*2)
		if hex.Row%2 == 0 {
			x += offset
		}
		y := 40.0 + offset*float64(hex.Row)*math.Sqrt(3)
		y = maxY - y

		poly := &polygon{x: x, y: y, radius: radius, label: hex.Label}
		poly.style.fill = terrainToFillColor(hex.Terrain)
		poly.style.stroke = "LightGrey"
		poly.style.stroke = "Grey"
		if poly.style.fill == poly.style.stroke {
			poly.style.stroke = "Black"
		}
		poly.style.strokeWidth = "2px"
		for _, p := range poly.hexPoints() {
			poly.points = append(poly.points, point{x: p.x, y: p.y})
		}

		s.polygons = append(s.polygons, poly)
	}

	elapsed := time.Now().Sub(start)
	log.Printf("[svg] elapsed time %+v\n", elapsed)

	return s
}

type point struct {
	x, y float64
}

func (p point) String() string {
	return fmt.Sprintf("%f,%f", p.x, p.y)
}

type polygon struct {
	x, y, radius float64
	label        string
	style        struct {
		fill        string
		stroke      string
		strokeWidth string
	}
	points []point
}

func (p polygon) hexPoints() (points []point) {
	for theta := 0.0; theta < math.Pi*2.0; theta += math.Pi / 3.0 {
		points = append(points, point{x: p.x + p.radius*math.Sin(theta), y: p.y + p.radius*math.Cos(theta)})
	}
	return points
}

func (p polygon) String() string {
	s := fmt.Sprintf(`<polygon style="fill: %s; stroke: %s; stroke-width: %s;"`, p.style.fill, p.style.stroke, p.style.strokeWidth)
	if len(p.points) != 0 {
		s += fmt.Sprintf(` points="`)
		for i, pt := range p.points {
			if i != 0 {
				s += " "
			}
			s += pt.String()
		}
		s += `"`
	}
	s += "></polygon>\n"
	s += fmt.Sprintf(`<text x="%f" y="%f" text-anchor="middle" fill="grey" font-size="12">%s</text>`, p.x, p.y, p.label)
	return s
}

type svg struct {
	id      string
	viewBox struct {
		minX, minY    int
		width, height int
	}
	polygons []*polygon
}

func (s svg) String() string {
	t := "<svg"
	if s.id != "" {
		t += fmt.Sprintf(" id=%q", s.id)
	}
	t += fmt.Sprintf(` width="%d" height="%d"`, s.viewBox.width+40, s.viewBox.height+40)
	t += fmt.Sprintf(` viewBox="%d %d %d %d"`, s.viewBox.minX, s.viewBox.minY, s.viewBox.width+40, s.viewBox.height+40)
	t += ` xmlns="http://www.w3.org/2000/svg">`
	for _, p := range s.polygons {
		t += fmt.Sprintf("\n%s", p.String())
	}
	return t + "\n</svg>"
}

// terrainToFillColor returns the background fill for a terrain type.
func terrainToFillColor(t string) string {
	switch t {
	case "Clear":
		return "hsl(53, 100%, 94%)"
	case "Delta":
		return "hsl(74, 48%, 76%)"
	case "Desert":
		return "hsl(48, 81%, 66%)"
	case "Gravel":
		return "hsl(49, 79%, 89%)"
	case "Mountain":
		return "hsl(47, 40%, 63%)"
	case "Ocean":
		return "LightBlue"
	case "Rock":
		return "hsl(49, 79%, 89%)"
	case "Rock/Gravel":
		return "hsl(49, 79%, 89%)"
	case "Rough":
		return "hsl(43, 43%, 77%)"
	case "Salt Marsh":
		return "hsl(65, 85%, 90%)"
	case "Sea":
		return "hsl(197, 78%, 85%)"
	case "Swamp":
		return "hsl(68, 78%, 93%)"
	case "Vegetation":
		return "hsl(85, 56%, 71%)"
	default:
		return "hsl(39, 100%, 50%)"
	}
}
