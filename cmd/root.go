package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/shanmugh/package-sorter/pkg/sorter"
	"github.com/spf13/cobra"
)

type config struct {
	inputFile          string
	maxVolume          int
	maxMass            int
	maxSingleDimension int
}

func NewCommand() (*cobra.Command, error) {
	cfg := config{}

	rootCmd := &cobra.Command{
		Use:   "package-sorter",
		Short: "Sorts packages into different categories based on their dimensions and weight",
		Long: `This application sorts packages into different categories based on their dimensions and weight.
	It supports the following categories:
	- STANDARD: Packages that are within the standard dimensions and weight limits
	- SPECIAL: Packages that are either too bulky or too heavy
	- REJECTED: Packages that are both too bulky and too heavy`,

		RunE: cfg.runE,
	}

	rootCmd.Flags().StringVarP(&cfg.inputFile, "input", "i", "", "Path to the input file")
	rootCmd.Flags().IntVarP(&cfg.maxVolume, "max-volume", "v", 1_000_000, "Maximum volume of the package")
	rootCmd.Flags().IntVarP(&cfg.maxMass, "max-mass", "m", 20, "Maximum mass of the package")
	rootCmd.Flags().IntVarP(&cfg.maxSingleDimension, "max-single-dimension", "d", 150, "Maximum single dimension of the package")

	rootCmd.MarkFlagRequired("input")

	return rootCmd, nil
}

type Input []sorter.Package

type Output struct {
	Standard []sorter.Package `json:"standard"`
	Special  []sorter.Package `json:"special"`
	Rejected []sorter.Package `json:"rejected"`
}

func (cfg *config) runE(cmd *cobra.Command, _ []string) error {
	content, err := os.ReadFile(cfg.inputFile)
	if err != nil {
		return errors.Wrapf(err, "read input file %s", cfg.inputFile)
	}

	input := Input{}
	if err := json.Unmarshal(content, &input); err != nil {
		return errors.Wrapf(err, "unmarshal input file %s", cfg.inputFile)
	}

	s := sorter.NewSorter(cfg.maxSingleDimension, cfg.maxVolume, cfg.maxMass)

	output := Output{}

	for _, p := range input {
		stack := s.Sort(p.Width, p.Height, p.Length, p.Mass)
		switch stack {
		case sorter.Standard.String():
			output.Standard = append(output.Standard, p)
		case sorter.Special.String():
			output.Special = append(output.Special, p)
		case sorter.Rejected.String():
			output.Rejected = append(output.Rejected, p)
		}
	}

	prettyJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(prettyJSON))

	return nil
}
