package cmd

// get the packages you will need
import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// fire up computeStats
var computeStatsCmd = &cobra.Command{
	Use:   "computeStats",
	Short: "Compute summary statistics for housing data",
	Run: func(cmd *cobra.Command, args []string) {
		computeStats()
	},
}

func init() {
	rootCmd.AddCommand(computeStatsCmd)
}

func computeStats() {
	start := time.Now() // Begin recording the processing time.

	file, _ := os.Open("housesInput.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // read the file now

	min, max, sum, count := initializeStats()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		for i, valStr := range record {
			val, _ := strconv.ParseFloat(valStr, 64)

			if val < min[i] {
				min[i] = val
			}
			if val > max[i] {
				max[i] = val
			}
			sum[i] += val
			count[i]++
		}
	}

	mean := make([]float64, len(sum))
	for i := range mean {
		mean[i] = sum[i] / float64(count[i])
	}

	writeStatsToFile(min, max, mean)

	duration := time.Since(start)                 // Close out the processing time.
	fmt.Printf("Processing time: %v\n", duration) // Drop the processing tiem into the terminal
}

func initializeStats() (min, max, sum []float64, count []int) {
	columns := 7
	min = make([]float64, columns)
	max = make([]float64, columns)
	sum = make([]float64, columns)
	count = make([]int, columns)

	// Time for a sweet for loop
	for i := range min {
		min[i] = 1e10
		max[i] = -1e10
	}
	return
}

// Record the summary stats in a file
func writeStatsToFile(min, max, mean []float64) {
	outputFile, _ := os.Create("housesOutputGo.txt")
	defer outputFile.Close()

	// Do it 100 times with a for loop
	for i := 0; i < 100; i++ {
		_, _ = fmt.Fprintf(outputFile, "Run %d\n", i+1)
		_, _ = fmt.Fprintln(outputFile, "Column\tMin\tMax\tMean")
		for j := 0; j < len(min); j++ {
			_, _ = fmt.Fprintf(outputFile, "%d\t%.2f\t%.2f\t%.2f\n", j+1, min[j], max[j], mean[j])
		}
		_, _ = fmt.Fprintln(outputFile, "")
	}
}
