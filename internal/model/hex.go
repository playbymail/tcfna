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

package model

// HEXES is a sortable slice of HEX values
type HEXES []*HEX

// Bounds returns the min and max values for rows and columns
func (h HEXES) Bounds() (minRow, maxRow, minCol, maxCol int) {
	minRow, maxRow, minCol, maxCol = 0, 0, 0, 0
	for _, hex := range h {
		if hex.Row < minRow {
			minRow = hex.Row
		}
		if maxRow < hex.Row {
			maxRow = hex.Row
		}
		if hex.Column < minCol {
			minCol = hex.Column
		}
		if maxCol < hex.Column {
			maxCol = hex.Column
		}
	}
	return minRow, maxRow, minCol, maxCol
}

// Len implements sort.Interface
func (h HEXES) Len() int {
	return len(h)
}

// Less implements sort.Interface
func (h HEXES) Less(i, j int) bool {
	if h[i].Row < h[j].Row {
		return true
	} else if h[i].Row == h[j].Row {
		return h[i].Column < h[j].Column
	}
	return false
}

// Swap implements sort.Interface
func (h HEXES) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// HEX is the data for a hex on the map board
type HEX struct {
	Id         string `json:"id"`
	Section    string `json:"section"`
	Row        int    `json:"row"`
	Column     int    `json:"column"`
	Label      string `json:"label,omitempty"`
	Name       string `json:"Name,omitempty"`
	Terrain    string `json:"terrain,omitempty"`
	Habitation string `json:"habitation,omitempty"`
	Misc       string `json:"misc,omitempty"`
	Sides      struct {
		NE HEXSIDE `json:"ne,omitempty"`
		E  HEXSIDE `json:"e,omitempty"`
		SE HEXSIDE `json:"se,omitempty"`
		SW HEXSIDE `json:"sw,omitempty"`
		W  HEXSIDE `json:"w,omitempty"`
		NW HEXSIDE `json:"nw,omitempty"`
	} `json:"sides,omitempty"`
}

type HEXSIDE struct {
	Elevation string `json:"elevation,omitempty"`
	Trans     string `json:"trans,omitempty"`
	Water     string `json:"water,omitempty"`
}
