package dlx

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/bketelsen/dlx/state"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

type CmdDocs struct {
	Global *state.Global
}

func (c *CmdDocs) Command() *cobra.Command {
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
			return doc.GenMarkdownTreeCustom(cmd.Root(), "site/content/docs/cmd", func(filename string) string {
				now := time.Now().Format(time.RFC3339)
				name := filepath.Base(filename)
				base := strings.TrimSuffix(name, path.Ext(name))
				return fmt.Sprintf(fmTemplate, strings.Replace(base, "_", " ", -1), strings.Replace(base, "_", " ", -1), strings.Replace(base, "_", " ", -1), now, now)
			}, func(s string) string {
				return "/docs/cmd/" + strings.TrimSuffix(s, ".md")
			})
		},
	}
	return docsCmd
}

const fmTemplate = `---
title: %s
description: %s
lead: %s
date: %s
lastmod: %s
draft: false
images: []
menu:
  docs:
    parent: "cli"
weight: 100
toc: true
---
`
