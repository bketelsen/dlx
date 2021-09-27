package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:                   "docs",
	Short:                 "Generates dlx's command line docs",
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Args:                  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Root().DisableAutoGenTag = true
		return doc.GenMarkdownTreeCustom(cmd.Root(), "site/content/docs/cmd", func(_ string) string {
			return ""
		}, func(s string) string {
			return "/docs/cmd/" + strings.TrimSuffix(s, ".md")
		})
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
