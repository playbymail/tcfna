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
	"github.com/fogleman/gg"
	"log"
	"math"
	"path/filepath"
	"time"
)

func (ds *STORE) BoardAsImage(save bool) {
	start := time.Now()

	// default background fill based on terrain
	clear := hslToRgb(53, 1.0, 0.94)       // hsl(53, 100%, 94%)
	delta := hslToRgb(74, 0.48, 0.76)      // hsl(74, 48%, 76%)
	desert := hslToRgb(48, 0.81, 0.66)     // hsl(48, 81%, 66%)
	gravel := hslToRgb(49, 0.79, 0.89)     // hsl(49, 79%, 89%)
	mountain := hslToRgb(47, 0.40, 0.63)   // hsl(47, 40%, 63%)
	ocean := hslToRgb(193, 0.67, 0.28)     // LightBlue
	rock := hslToRgb(49, 0.79, 0.89)       // hsl(49, 79%, 89%)
	rockGravel := hslToRgb(49, 0.79, 0.89) // hsl(49, 79%, 89%)
	rough := hslToRgb(43, 0.43, 0.77)      // hsl(43, 43%, 77%)
	saltMarsh := hslToRgb(65, 0.85, 0.90)  // hsl(65, 85%, 90%)
	sea := hslToRgb(198, 0.78, 0.86)       // hsl(197, 78%, 85%)
	swamp := hslToRgb(68, 0.78, 0.93)      // hsl(68, 78%, 93%)
	unknown := hslToRgb(39, 1.0, 0.5)      // hsl(39, 100%, 50%)
	vegetation := hslToRgb(85, 0.56, 0.71) // hsl(85, 56%, 71%)

	// find the min and max values for rows and columns
	_, maxRow, _, maxCol := ds.board.Sorted.Bounds()

	border, radius, rotation := 40.0, 30.0, math.Pi/2
	offset := (math.Sqrt(3) * radius) / 2
	maxX := 2*border + offset*float64(maxCol*2)
	maxY := 3*border + offset*float64(maxRow)*math.Sqrt(3)

	dc := gg.NewContext(int(maxX+radius), int(maxY+radius))

	a := 1.0 // default alpha to opaque

	// default background color to something grey-ish.
	// this will become the border for the map.
	dc.SetRGBA(0.8, 0.8, 0.8, a)
	dc.Clear() // clears and fills entire image with current color

	// convert hexes to a png image
	for _, hex := range ds.board.Sorted {
		x := offset * float64(hex.Column*2)
		if hex.Row%2 == 0 {
			x += offset
		}
		y := border + offset*float64(hex.Row)*math.Sqrt(3)
		y = maxY - y // flips the coordinates

		// draws the "path" of the hex
		dc.DrawRegularPolygon(6, x, y, radius, rotation)

		// use the terrain to determine the fill for the hex
		var fillColor HSL
		switch hex.Terrain {
		case "Clear":
			fillColor = clear
		case "Delta":
			fillColor = delta
		case "Desert":
			fillColor = desert
		case "Gravel":
			fillColor = gravel
		case "Mountain":
			fillColor = mountain
		case "Ocean":
			fillColor = ocean
		case "Rock":
			fillColor = rock
		case "Rock/Gravel":
			fillColor = rockGravel
		case "Rough":
			fillColor = rough
		case "Salt Marsh":
			fillColor = saltMarsh
		case "Sea":
			fillColor = sea
		case "Swamp":
			fillColor = swamp
		case "Vegetation":
			fillColor = vegetation
		default:
			fillColor = unknown
		}
		dc.SetRGBA(fillColor.red, fillColor.green, fillColor.blue, a)
		dc.FillPreserve()

		dc.SetRGBA(0, 0, 0, a)
		dc.DrawStringAnchored(hex.Label, x, y-radius/3, 0.5, 0.5)
		if hex.Sides.NE.Elevation != "" {
			dc.DrawStringAnchored(hex.Sides.NE.Elevation, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.NE.Trans != "" {
			dc.DrawStringAnchored(hex.Sides.NE.Trans, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.NE.Water != "" {
			dc.DrawStringAnchored(hex.Sides.NE.Water, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.E.Elevation != "" {
			dc.DrawStringAnchored(hex.Sides.E.Elevation, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.E.Trans != "" {
			dc.DrawStringAnchored(hex.Sides.E.Trans, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.E.Water != "" {
			dc.DrawStringAnchored(hex.Sides.E.Water, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.SE.Elevation != "" {
			dc.DrawStringAnchored(hex.Sides.SE.Elevation, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.SE.Trans != "" {
			dc.DrawStringAnchored(hex.Sides.SE.Trans, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.SE.Water != "" {
			dc.DrawStringAnchored(hex.Sides.SE.Water, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.SW.Elevation != "" {
			dc.DrawStringAnchored(hex.Sides.SW.Elevation, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.SW.Trans != "" {
			dc.DrawStringAnchored(hex.Sides.SW.Trans, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.SW.Water != "" {
			dc.DrawStringAnchored(hex.Sides.SW.Water, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.SW.Elevation != "" {
			dc.DrawStringAnchored(hex.Sides.W.Elevation, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.W.Trans != "" {
			dc.DrawStringAnchored(hex.Sides.W.Trans, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.W.Water != "" {
			dc.DrawStringAnchored(hex.Sides.W.Water, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.NW.Elevation != "" {
			dc.DrawStringAnchored(hex.Sides.NW.Elevation, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.NW.Trans != "" {
			dc.DrawStringAnchored(hex.Sides.NW.Trans, x, y+radius/3, 0.5, 0.5)
		} else if hex.Sides.NW.Water != "" {
			dc.DrawStringAnchored(hex.Sides.NW.Water, x, y+radius/3, 0.5, 0.5)
		}

		// draw a black line around the hex
		dc.SetRGBA(0, 0, 0, a)
		dc.SetLineWidth(4.0)
		dc.Stroke() // colors the line path and clears the path
	}

	elapsed := time.Now().Sub(start)
	log.Printf("[png] elapsed time %+v\n", elapsed)

	if save {
		_ = dc.SavePNG(filepath.Join("..", "data", "board.png"))

		elapsed = time.Now().Sub(start)
		log.Printf("[png] elapsed time %+v\n", elapsed)
	}
}

type HSL struct {
	red, green, blue float64
}

func hslToRgb(hue, saturation, luminance float64) HSL {
	// hue must be between 0 and 360
	for hue < 0 {
		hue += 360
	}
	for hue > 360 {
		hue -= 360
	}
	// saturation must be between 0 and 1
	for saturation < 0 {
		saturation += 1
	}
	for saturation > 1 {
		saturation -= 1
	}
	// luminance must be between 0 and 1
	for luminance < 0 {
		luminance += 1
	}
	for luminance > 1 {
		luminance -= 1
	}

	if saturation < 0.001 {
		// gray shade
		return HSL{luminance, luminance, luminance}
	}

	var t1, t2 float64
	if luminance < 0.5 {
		t1 = luminance * (1.0 + saturation)
	} else {
		t1 = luminance + saturation - (luminance * saturation)
	}
	t2 = 2*luminance - t1

	// convert 0..360 degrees to 0..1
	hue = hue / 360.0

	// default temp values for red, green, blue based on the hue.
	// then force them to the range of 0..1
	tR, tG, tB := hue+0.333, hue, hue-0.33
	for tR < 0 {
		tR += 1
	}
	for tR > 1 {
		tR -= 1
	}
	for tG < 0 {
		tG += 1
	}
	for tG > 1 {
		tG -= 1
	}
	for tB < 0 {
		tB += 1
	}
	for tB > 1 {
		tB -= 1
	}

	var red float64
	if tR*6.0 < 1.0 {
		red = t2 + (t1-t2)*6.0*tR
	} else if tR*2.0 < 1.0 {
		red = t1
	} else if tR*3.0 < 2.0 {
		red = t2 + (t1-t2)*(0.666-tR)*6.0
	} else {
		red = t2
	}

	var green float64
	if tG*6.0 < 1.0 {
		green = t2 + (t1-t2)*6.0*tG
	} else if tG*2.0 < 1.0 {
		green = t1
	} else if tG*3.0 < 2.0 {
		green = t2 + (t1-t2)*(0.666-tG)*6.0
	} else {
		green = t2
	}

	var blue float64
	if tB*6.0 < 1.0 {
		blue = t2 + (t1-t2)*6.0*tB
	} else if tB*2.0 < 1.0 {
		blue = t1
	} else if tB*3.0 < 2.0 {
		blue = t2 + (t1-t2)*(0.666-tB)*6.0
	} else {
		blue = t2
	}

	return HSL{red, green, blue}
}
