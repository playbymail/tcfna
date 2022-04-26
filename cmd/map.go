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

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mdhender/tcfna/internal/model"
	"github.com/mdhender/tcfna/internal/store/csvdb"
	"github.com/mdhender/tcfna/internal/store/memory"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"path/filepath"
)

var mapGlobals struct {
	Import struct {
		Format string // csv, json
		Name   string // leave blank to avoid import
	}
	Export struct {
		Format string // html, json, png, svg
		Name   string // leave blank to avoid export
	}
}

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "map commands",
	Long:  `Commands to import, export, and create board maps.`,
	Run: func(cmd *cobra.Command, args []string) {
		//rootDir := viper.Get("files.path").(string)
		//log.Printf("[map] rootDir %q\n", rootDir)

		var board *model.MAP
		var err error

		log.Println(mapGlobals)
		if mapGlobals.Import.Name != "" {
			switch mapGlobals.Import.Format {
			case "csv":
				board, err = csvdb.Convert(mapGlobals.Import.Name)
				cobra.CheckErr(err)
			case "json":
				//board, err = jsondb.Convert(mapGlobals.Import.Name)
				//cobra.CheckErr(err)
			default:
				log.Fatalf("[map] unsupported import format %q\n", mapGlobals.Import.Format)
			}
		}
		if board == nil {
			cobra.CheckErr(fmt.Errorf("missing input file name"))
		}

		ds := memory.New(board)

		if mapGlobals.Export.Name != "" {
			switch mapGlobals.Export.Format {
			case "html":
				cobra.CheckErr(ioutil.WriteFile("svg-test.html", ds.BoardAsHTML(), 0644))
			case "json":
				if mapGlobals.Import.Name == mapGlobals.Export.Name {
					log.Fatal("[map] cowardly refusing to overwrite input file\n")
				}
				data := struct {
					Data model.HEXES `json:"data"`
				}{Data: board.Sorted}
				b, err := json.MarshalIndent(data, "", "  ")
				if err != nil {
					log.Fatalf("[map] encoding json: %+v\n", err)
				}
				cobra.CheckErr(ioutil.WriteFile(filepath.Join("..", "data", "board.json"), b, 0644))
				ds.BoardAsImage(true)
			case "png":
				ds.BoardAsImage(true)
			case "svg":
				cobra.CheckErr(ioutil.WriteFile(filepath.Join("..", "data", "board.svg"), []byte(ds.BoardAsSVG().String()), 0644))
			default:
				log.Fatalf("[map] unsupported export format %q\n", mapGlobals.Export.Format)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(mapCmd)
	mapCmd.Flags().StringVar(&mapGlobals.Import.Name, "import", "", "file name to read board map data from")
	mapCmd.Flags().StringVar(&mapGlobals.Import.Format, "import-format", "json", "file format for imported data")
	mapCmd.Flags().StringVar(&mapGlobals.Export.Name, "export", "", "file name to write board map data to")
	mapCmd.Flags().StringVar(&mapGlobals.Export.Format, "export-format", "png", "file format for exported data")
}
