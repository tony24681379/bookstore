package productCmd

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/product"
)

type productMarkOptions struct {
	id   string
	good string
	long string
	risk string
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
	flags.StringVar(&opts.good, "good", "", "Mark the the as good sales")
	flags.StringVar(&opts.long, "long-term", "", "Mark the book as long term")
	flags.StringVar(&opts.risk, "risk", "", "Mark the book as risk")
	return cmd
}

func runMark(db *gorm.DB, opts productMarkOptions) {
	product.Mark(db, opts.id, opts.good, opts.long, opts.risk, opts.due)
}
