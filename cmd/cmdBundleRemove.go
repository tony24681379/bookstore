package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/bundle"
)

type bundleRemoveOptions struct {
	bundleID string
	bookID   int
}

//NewBundleRemoveCmd Remove the book into the bundle
func NewBundleRemoveCmd() *cobra.Command {
	var opts bundleRemoveOptions

	cmd := &cobra.Command{
		Use:   "remove bundleID",
		Short: "Remove the book from the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			opts.bundleID = args[0]
			runRemove(opts)
		},
	}
	flags := cmd.Flags()
	flags.IntVar(&opts.bookID, "book-id", 0, "Book ID")
	return cmd
}

func runRemove(opts bundleRemoveOptions) {
	bundle.Remove(opts.bundleID, opts.bookID)
}
