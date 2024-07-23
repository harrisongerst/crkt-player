package cmd
import (
	"fmt"
	"os"

	"hgerst/crkt/client"
	"hgerst/crkt/views"
	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
)

var rootCmd = &cobra.Command{
	Use:   "crkt",
	Short: "crkt music",
	Long: "crkt music and sound player",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to crkt the command line sound player, try typing 'crkt play'")
	},
  }

var play = &cobra.Command{
	Use: "play",
	Short: "this plays the music",
	Long: "heres the long description of it playing music",
	Run: func(cmd *cobra.Command, args []string) {
		client.PlayFile(args[0])
	},
}
var start = &cobra.Command{
	Use: "start",
	Short: "this plays the music",
	Long: "heres the long description of it playing music",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(views.SelectInitialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
} 
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }

  func init() {
	rootCmd.AddCommand()
	rootCmd.AddCommand(play)
	rootCmd.AddCommand(start)
  }