/*******************************************************************************
 * TCFNA
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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var testFlag bool
var verboseFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tcfna",
	Short: "TCFNA engine",
	Long: `tcfna is a game engine for TFCNA. This application creates
new games, executes orders, and generates reports for each player.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// bind viper and cobra here since this hook runs early and always
		return bindConfig(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("env: %-30s == %q\n", "HOME", homeFolder)
		log.Printf("env: %-30s == %q\n", "TCFNA_CONFIG", viper.ConfigFileUsed())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tcfna.yaml)")
	rootCmd.PersistentFlags().BoolVar(&testFlag, "test", false, "test mode")
	rootCmd.PersistentFlags().BoolVar(&verboseFlag, "verbose", false, "verbose mode")

	// Cobra also supports local flags, which will only run when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
