package cmd
import (
	"fmt"
	"os"

	"hgerst/crkt/player"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to crkt the command line sound player, try typing 'crkt play'")
	},
  }

var play = &cobra.Command{
	Use: "play",
	Short: "this plays the music",
	Long: "heres the long description of it playing music",
	Run: func(cmd *cobra.Command, args []string) {
		player.PlayRain()
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
  }