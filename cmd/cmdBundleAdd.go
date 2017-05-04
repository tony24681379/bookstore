package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type bundleAddOptions struct {
	bundleID int
	bookID   int
}

//NewBundleAddCmd add the book into the bundle
func NewBundleAddCmd() *cobra.Command {
	var opts bundleAddOptions

	cmd := &cobra.Command{
		Use:   "add [OPTIONS]",
		Short: "Add the book into the bundle",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Add the book into the bundle")
		},
	}
	flags := cmd.Flags()
	flags.IntVar(&opts.bundleID, "bundle-id", -1, "Bundle ID")
	flags.IntVar(&opts.bookID, "book id", -1, "Book ID")
	return cmd
}
