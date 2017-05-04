package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleModifyOptions struct {
	bundleID string
	Note     string
}

//NewBundleModifyCmd Modify the book into the bundle
func NewBundleModifyCmd() *cobra.Command {
	var opts bundleModifyOptions
	cmd := &cobra.Command{
		Use:   "modify bundleID",
		Short: "Modify the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			opts.bundleID = args[0]
			runModify(opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVar(&opts.Note, "note", "", "modify note to bundle")
	return cmd
}

func runModify(opts bundleModifyOptions) {
	bundle.Modify(opts.bundleID, opts.Note)
}
