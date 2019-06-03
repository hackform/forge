package cmd

import (
	"github.com/hackform/forge/gen"
	"github.com/spf13/cobra"
)

var (
	genPrefix   string
	genSuffix   string
	genNoIgnore bool
	genDryRun   bool
	genVerbose  bool
)

// genCmd represents the model command
var genCmd = &cobra.Command{
	Use:   "gen [path | file glob ...]",
	Short: "Executes command line directives in source files",
	Long: `Executes command line directives in source files

Directives appear in the form of:

	<prefix>command args[<suffix>|'\n'|EOF]

forge gen directives end on the first new line or suffix

Arguments may also be replaced by environment variables:

file.txt:
	+forge:gen echo $TEST ${TEST} ${BOGUS:-default val} "$TEST ${TEST}" '$TEST ${TEST}' +gen:end

output:
	$ TEST=test forge gen
	test test default val test test $TEST ${TEST}
`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.Execute(genPrefix, genSuffix, genNoIgnore, genDryRun, genVerbose, args)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")
	genCmd.PersistentFlags().BoolVarP(&genNoIgnore, "noignore", "i", false, "do not use .gitignore")
	genCmd.PersistentFlags().StringVarP(&genPrefix, "prefix", "p", "+forge:gen", "set prefix for forge directive")
	genCmd.PersistentFlags().StringVarP(&genSuffix, "suffix", "s", "+gen:end", "set suffix for forge directive")
	genCmd.PersistentFlags().BoolVarP(&genDryRun, "dryrun", "n", false, "do not exec directives but print what would be executed")
	genCmd.PersistentFlags().BoolVarP(&genVerbose, "verbose", "v", false, "increase the verbosity of the output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
