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

type GlobalData struct {
	Attributes Attribute
}
type Attribute struct {
	OBJECTID       string
	Country_Region string
	Confirmed      int64
	Deaths         int64
	Recovered      int64
	Active         int64
}

// globalDatatCmd represents the version command
var globalDatatCmd = &cobra.Command{
	Use:   "global-data",
	Short: "Showing global corona data",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://api.kawalcorona.com/")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var globaldata []GlobalData

			err := json.Unmarshal(data, &globaldata)
			if err != nil {
				fmt.Println("error:", err)
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("Country Region", "Confirmed", "Deaths", "Recovered", "Active")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, dt := range globaldata {
				dat := dt.Attributes
				tbl.AddRow(dat.Country_Region, dat.Confirmed, dat.Deaths, dat.Recovered, dat.Active)
			}

			tbl.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(globalDatatCmd)
}
