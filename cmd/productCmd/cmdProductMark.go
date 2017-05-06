package productCmd

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/product"
)

type productMarkOptions struct {
	id   string
	good bool
	long bool
	risk bool
	due  time.Time
}

//NewProductMarkCmd Mark the book into the product
func NewProductMarkCmd(db *gorm.DB) *cobra.Command {
	var opts productMarkOptions

	cmd := &cobra.Command{
		Use:   "mark",
		Short: "Mark the book which the product",
		Run: func(cmd *cobra.Command, args []string) {
			opts.id = args[0]
			runMark(db, opts)
		},
	}
	flags := cmd.Flags()
	flags.BoolVar(&opts.good, "good", false, "Mark the the as good sales")
	flags.BoolVar(&opts.long, "long-term", false, "Mark the book as long term")
	flags.BoolVar(&opts.risk, "risk", false, "Mark the the as risk")
	return cmd
}

func runMark(db *gorm.DB, opts productMarkOptions) {
	product.Mark(db, opts.id, opts.good, opts.long, opts.risk, opts.due)
}
