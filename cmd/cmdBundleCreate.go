package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleCreateOptions struct {
	BundleName string
	Note       string
}

//NewBundleCreateCmd Create the book into the bundle
func NewBundleCreateCmd() *cobra.Command {
	var opts bundleCreateOptions
	cmd := &cobra.Command{
		Use:   "create bundleName",
		Short: "Create the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			opts.BundleName = args[0]
			runCreate(opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.Note, "note", "", "Add note to bundle")
	return cmd
}

func runCreate(opts bundleCreateOptions) {
	bundle.Create(opts.BundleName, opts.Note)
}
