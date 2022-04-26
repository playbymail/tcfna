/*******************************************************************************
 * TCFNA - Game Engine for SPI's Campaign for North Africa
 * Copyright (C) 2022. Michael D Henderson
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
 ******************************************************************************/

package model

type MAP struct {
	Hexes  map[string]*HEX `json:"hexes"`
	Sorted HEXES
}

type TERRAIN struct {
	Section string `json:"section,omitempty"`
}

type Terrain struct {
	Airfield                TERRAIN `json:"airfield,omitempty"`
	Border                  TERRAIN `json:"border,omitempty"`
	Clear                   TERRAIN `json:"clear,omitempty"`
	Coast                   TERRAIN `json:"coast,omitempty"`
	Delta                   TERRAIN `json:"delta,omitempty"`
	Desert                  TERRAIN `json:"desert,omitempty"`
	Escarpment              TERRAIN `json:"escarpment,omitempty"`
	FlyingBoatAlightingArea TERRAIN `json:"flying boat alighting area,omitempty"`
	FlyingBoatBasin         TERRAIN `json:"flying boat basin,omitempty"`
	HeavyVegetation         TERRAIN `json:"heavy vegetation,omitempty"`
	MajorCity               TERRAIN `json:"major city,omitempty"`
	MajorRiver              TERRAIN `json:"major river,omitempty"`
	MinorRiver              TERRAIN `json:"minor river,omitempty"`
	Mountain                TERRAIN `json:"mountain,omitempty"`
	OffMapAirfield          TERRAIN `json:"off map airfield,omitempty"`
	OffMapFlyingBoatBasin   TERRAIN `json:"off map flying boat basin,omitempty"`
	Oasis                   TERRAIN `json:"oasis,omitempty"`
	Port                    TERRAIN `json:"port,omitempty"`
	Railroad                TERRAIN `json:"railroad,omitempty"`
	Ridge                   TERRAIN `json:"ridge,omitempty"`
	Road                    TERRAIN `json:"road,omitempty"`
	RockGravel              TERRAIN `json:"rock/gravel,omitempty"`
	Rough                   TERRAIN `json:"rough,omitempty"`
	SaltMarsh               TERRAIN `json:"salt marsh,omitempty"`
	Sea                     TERRAIN `json:"sea,omitempty"`
	Slope                   TERRAIN `json:"slope,omitempty"`
	Swamp                   TERRAIN `json:"swamp,omitempty"`
	Track                   TERRAIN `json:"track,omitempty"`
	TrainingArea            TERRAIN `json:"training area,omitempty"`
	UnfinishedRailroad      TERRAIN `json:"unfinished railroad,omitempty"`
	UnfinishedRoad          TERRAIN `json:"unfinished road,omitempty"`
	VillageBir              TERRAIN `json:"village/bir,omitempty"`
	Wadi                    TERRAIN `json:"wadi,omitempty"`
}
