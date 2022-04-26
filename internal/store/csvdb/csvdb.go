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

package csvdb

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/mdhender/tcfna/internal/model"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// record is the structure from Michael Miller's CNA-Hex-Database-20150117.csv file.
type record struct {
	hexID         string
	mapHex        string
	hexColumn     string
	hexRow        string
	hexRRCCC      string
	name          string
	terrain       string
	habitation    string
	misc          string
	hsElevationNE string
	hsElevationE  string
	hsElevationSE string
	hsElevationSW string
	hsElevationW  string
	hsElevationNW string
	pad           string
	hsTransNE     string
	hsTransE      string
	hsTransSE     string
	hsTransSW     string
	hsTransW      string
	hsTransNW     string
	hsWaterNE     string
	hsWaterE      string
	hsWaterSE     string
	hsWaterSW     string
	hsWaterW      string
	hsWaterNW     string
}

func Convert(name string) (*model.MAP, error) {
	// load the CSV file from disk
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// create a standard reader.
	// the file isn't quote escaped, so we shouldn't need additional parameters.
	records, err := csv.NewReader(bytes.NewReader(b)).ReadAll()
	if err != nil {
		return nil, err
	}

	// log the raw number of records as a first sanity check
	log.Printf("[csvdb] read %d records\n", len(records))
	if len(records) == 0 {
		return nil, fmt.Errorf("input file is empty")
	} else if len(records) < 2 {
		return nil, fmt.Errorf("input file has no header")
	} else if len(records) < 3 {
		return nil, fmt.Errorf("input file has no sub-header")
	}

	// these are the header and sub-header columns we expect.
	// we pulled this list from the 2015/01/17 version of Michael Miller's file.
	headerColumns := []string{"HexID", "Map Hex", "Hex Column", "Hex Row", "Hex RRCCC", "Name", "Terrain", "Habitation", "Misc", "hsElevation NE", "hsElevation E", "hsElevation SE", "hsElevation SW", "hsElevation W", "hsElevation NW", "", "hsTrans NE", "hsTrans E", "hsTrans SE", "hsTrans SW", "hsTrans W", "hsTrans NW", "hsWater NE", "hsWater E", "hsWater SE", "hsWater SW", "hsWater W", "hsWater NW"}
	subHeaderColumns := []string{"", "", "", "", "", "", "", "", "", "1-DnEsc", "2-DnSlp", "3-Ridge", "4-UpEsc", "5-UpSlp", "", "", "1-Track", "2-Road", "3-RR", "4-UnfRd", "5-UnfRR", "6-Rd&RR", "1-Wadi", "2-SeaHS", "3-River", "4-Nile", "5-Border", "", ""}

	var badInput bool

	for col, colName := range headerColumns {
		if !(col < len(records[0])) {
			log.Printf("[csvdb] %d: header: missing column %3d %q\n", len(records[0]), col+1, colName)
			badInput = true
		} else if records[0][col] != colName {
			log.Printf("[csvdb] %d: header: missing column %3d %q / %q\n", len(records[0]), col+1, colName, records[0][col])
			badInput = true
		}
	}
	for col, colName := range subHeaderColumns {
		if !(col < len(records[1])) {
			log.Printf("[csvdb] %d: sub-header: missing column %3d %q\n", len(records[1]), col+1, colName)
			badInput = true
		} else if records[1][col] != colName {
			log.Printf("[csvdb] %d: sub-header: missing column %3d %q / %q\n", len(records[0]), col+1, colName, records[0][col])
			badInput = true
		}
	}
	// header and sub-header meet our expectations.
	// let us now validate the length of the data rows.
	expectedColumns := len(headerColumns) + 1 // because the input seems to have an extra (but empty) column
	for row, record := range records {
		if len(record) < expectedColumns {
			log.Printf("[csvdb] %5d: data: missing   columns %3d/%3d\n", row+1, len(record), expectedColumns)
			badInput = true
		} else if len(record) > expectedColumns {
			log.Printf("[csvdb] %5d: data: additonal columns %3d/%3d\n", row+1, len(record), expectedColumns)
			badInput = true
		}
	}

	if badInput {
		return nil, fmt.Errorf("input not as expected")
	}

	// the file has the header and data rows that seem to be the right length.
	// now go through and convert every data record
	m := &model.MAP{Hexes: make(map[string]*model.HEX)}
	for _, rec := range records[2:] { // the 2: lets us skip the headers
		// copy every column into a map record so that we can use names instead
		// of column numbers during the conversion.
		r := recordToRow(rec)

		// hexId seems to have the format BRRCC, where
		//  B   is the board number (board A is 1, board E is 5)
		//  RR  is the two digit row number on that board
		//  CC  is the two digit column number on that board
		// Both RR and CC are left-padded with zeroes.
		hexId, _ := strconv.Atoi(r.hexID)
		section := []string{"*", "A", "B", "C", "D", "E"}[hexId/10_000]
		sectionRow, sectionCol := (hexId/100)%100, hexId%100

		// if section == "C" && sectionRow == 47 && sectionCol == 8 { // (hexId == 34807 || hexId == 34708)
		//   log.Printf("B %s RR %2d CC %2d input %+v\n", section, sectionRow, sectionCol, r)
		// } else if section == "C" && sectionRow == 48 && sectionCol == 7 { // (hexId == 34807 || hexId == 34708)
		//   log.Printf("B %s RR %2d CC %2d mapHex %q\n", section, sectionRow, sectionCol, r.mapHex)
		// }

		switch section { // boards A-D are 33 columns wide
		case "A":
			// no adjustment needed
		case "B":
			sectionCol += 33
		case "C":
			sectionCol += 2 * 33
		case "D":
			sectionCol += 3 * 33
		case "E":
			sectionCol += 4 * 33
		}

		hex := &model.HEX{
			Id:         fmt.Sprintf("%02d%03d", sectionRow, sectionCol),
			Label:      fmt.Sprintf("%s%02d%02d", section, (hexId/100)%100, hexId%100),
			Section:    section,
			Row:        sectionRow,
			Column:     sectionCol,
			Name:       r.name,
			Habitation: r.habitation,
			Terrain:    r.terrain,
		}
		hex.Sides.NE.Elevation = r.hsElevationNE
		hex.Sides.NE.Trans = r.hsTransNE
		hex.Sides.NE.Water = r.hsWaterNE
		hex.Sides.E.Elevation = r.hsElevationE
		hex.Sides.E.Trans = r.hsTransE
		hex.Sides.E.Water = r.hsWaterE
		hex.Sides.SE.Elevation = r.hsElevationSE
		hex.Sides.SE.Trans = r.hsTransSE
		hex.Sides.SE.Water = r.hsWaterSE
		hex.Sides.SW.Elevation = r.hsElevationSW
		hex.Sides.SW.Trans = r.hsTransSW
		hex.Sides.SW.Water = r.hsWaterSW
		hex.Sides.W.Elevation = r.hsElevationW
		hex.Sides.W.Trans = r.hsTransW
		hex.Sides.W.Water = r.hsWaterW
		hex.Sides.NW.Elevation = r.hsElevationNW
		hex.Sides.NW.Trans = r.hsTransNW
		hex.Sides.NW.Water = r.hsWaterNW

		m.Hexes[hex.Id] = hex
		m.Sorted = append(m.Sorted, hex)
	}

	sort.Sort(m.Sorted)

	return m, nil
}

// recordToRow copies the raw input into a struct.
// The fields are trimmed of leading and trailing whitespace.
func recordToRow(row []string) *record {
	return &record{
		hexID:         strings.TrimSpace(row[0]),
		mapHex:        strings.TrimSpace(row[1]),
		hexColumn:     strings.TrimSpace(row[2]),
		hexRow:        strings.TrimSpace(row[3]),
		hexRRCCC:      strings.TrimSpace(row[4]),
		name:          strings.TrimSpace(row[5]),
		terrain:       strings.TrimSpace(row[6]),
		habitation:    strings.TrimSpace(row[7]),
		misc:          strings.TrimSpace(row[8]),
		hsElevationNE: strings.TrimSpace(row[9]),
		hsElevationE:  strings.TrimSpace(row[10]),
		hsElevationSE: strings.TrimSpace(row[11]),
		hsElevationSW: strings.TrimSpace(row[12]),
		hsElevationW:  strings.TrimSpace(row[13]),
		hsElevationNW: strings.TrimSpace(row[14]),
		pad:           strings.TrimSpace(row[15]),
		hsTransNE:     strings.TrimSpace(row[16]),
		hsTransE:      strings.TrimSpace(row[17]),
		hsTransSE:     strings.TrimSpace(row[18]),
		hsTransSW:     strings.TrimSpace(row[19]),
		hsTransW:      strings.TrimSpace(row[20]),
		hsTransNW:     strings.TrimSpace(row[21]),
		hsWaterNE:     strings.TrimSpace(row[22]),
		hsWaterE:      strings.TrimSpace(row[23]),
		hsWaterSE:     strings.TrimSpace(row[24]),
		hsWaterSW:     strings.TrimSpace(row[25]),
		hsWaterW:      strings.TrimSpace(row[26]),
		hsWaterNW:     strings.TrimSpace(row[27]),
	}
}
