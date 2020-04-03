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

type IndoDetailData struct {
	Attributes DetailData
}

type DetailData struct {
	FID        int64
	Kode_Provi int64
	Provinsi   string
	Kasus_Posi int64
	Kasus_Semb int64
	Kasus_Meni int64
}

// globalDatatCmd represents the version command
var indoDetailDatatCmd = &cobra.Command{
	Use:   "indonesia-detail-data",
	Short: "Showing indonesia corona data by province",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://api.kawalcorona.com/indonesia/provinsi/")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			var indodata []IndoDetailData

			err := json.Unmarshal(data, &indodata)
			if err != nil {
				fmt.Println("error:", err)
			}
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("Province", "Confirmed", "Deaths", "Recovered")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, dt := range indodata {
				dat := dt.Attributes
				tbl.AddRow(dat.Provinsi, dat.Kasus_Posi, dat.Kasus_Meni, dat.Kasus_Semb)
			}

			tbl.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(indoDetailDatatCmd)
}
