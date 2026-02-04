package main

import (
	"conf-browser/backend"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var (
	port   int
	noOpen bool

	//go:embed build/*
	frontend embed.FS
)

// cmd represents the base command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "conf-browser",
	Short: "Browse all conference papers in one place",
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.Handle("/", backend.MakeFileserverHandler(frontend, "build"))
		mux.HandleFunc("/backend/exit", backend.ExitHandler)
		mux.HandleFunc("/backend/papers/{conf}/{year}", backend.GetPapersHandler)

		// Open the web app in the default browser
		url := fmt.Sprintf("http://localhost:%d", port)
		if !noOpen {
			_ = open.Run(url)
		}
		fmt.Printf("Server running on %s\n", url)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), backend.AllowAllCORSMiddleware(mux)))
	},
}

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// completions is a subcommand that generates shell completions
var completions = &cobra.Command{
	Use:       "completion [bash|zsh|fish|powershell]",
	Short:     "Generate shell completions",
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	cmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	cmd.Flags().BoolVarP(&noOpen, "no-open", "n", false, "Do not open the browser automatically")
	cmd.AddCommand(completions)
}
