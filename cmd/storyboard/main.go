package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"training-storyboard.local/workbench/internal/workbench"
)

type consoleOutput struct {
	GeneratedProject workbench.CourseProject `json:"generated_project"`
	Export           workbench.ExportRecord  `json:"export"`
	UserCenter       workbench.UserCenter    `json:"user_center"`
}

func main() {
	goal := flag.String("goal", "门店值班经理能够在高峰期快速分流客诉，并完成事实记录、补救授权和班后复盘", "course goal")
	format := flag.String("format", "pptx", "export format: pdf or pptx")
	flag.Parse()

	service := workbench.NewService()
	project, err := service.CreateProject(*goal)
	if err != nil {
		exitWithError(err)
	}
	record, err := service.CreateExport(project.ID, *format)
	if err != nil {
		exitWithError(err)
	}

	output := consoleOutput{
		GeneratedProject: project,
		Export:           record,
		UserCenter:       service.UserCenter(),
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(output); err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, "storyboard:", err)
	os.Exit(1)
}
