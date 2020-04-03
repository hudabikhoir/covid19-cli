package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type IndoData struct {
	Name      string
	Positif   string
	Sembuh    string
	Meninggal string
}

// globalDatatCmd represents the version command
var indoDatatCmd = &cobra.Command{
	Use:   "indonesia-data",
	Short: "Showing indonesia corona data",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://api.kawalcorona.com/indonesia/")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var indodata []IndoData

			err := json.Unmarshal(data, &indodata)
			if err != nil {
				fmt.Println("error:", err)
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("", "Confirmed", "Deaths", "Recovered")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, dat := range indodata {
				tbl.AddRow(dat.Name, dat.Positif, dat.Meninggal, dat.Sembuh)
			}

			tbl.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(indoDatatCmd)
}
