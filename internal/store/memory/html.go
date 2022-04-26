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

package memory

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

func (ds *STORE) BoardAsHTML() []byte {
	start := time.Now()

	// create the svg for the board
	b := &bytes.Buffer{}
	_, _ = fmt.Fprintln(b, `<!doctype html>`)
	_, _ = fmt.Fprintln(b, `<html lang="en">`)
	_, _ = fmt.Fprintln(b, `<head>`)
	_, _ = fmt.Fprintln(b, `<meta charset="utf-8">`)
	_, _ = fmt.Fprintln(b, `<title>SVG Test</title>`)
	//_, _ = fmt.Fprintln(b, `<style>div.scroll {background-color: #fed9ff;width: 95%;height: 95%;overflow: auto;text-align: justify;padding: 1%;}</style>`)
	_, _ = fmt.Fprintln(b, `</head>`)
	_, _ = fmt.Fprintln(b, `<body>`)
	//_, _ = fmt.Fprintln(b, `<div class="scroll">`)
	_, _ = fmt.Fprintln(b, ds.BoardAsSVG().String())
	//_, _ = fmt.Fprintln(b, `</div>`)
	_, _ = fmt.Fprintln(b, "</body>")
	_, _ = fmt.Fprintln(b, "</html>")

	elapsed := time.Now().Sub(start)
	log.Printf("[html] elapsed time %+v\n", elapsed)

	return b.Bytes()
}
