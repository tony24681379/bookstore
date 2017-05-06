package recommendCmd

import (
	"math"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/recommend"
)

type recommendCreateOptions struct {
	recommendName string
	bookstoreID   uint
	capacity      uint
	dayLow        uint
	dayHigh       uint
	weekLow       uint
	weekHigh      uint
	monthLow      uint
	monthHigh     uint
	stackLow      uint
	stackHigh     uint
	stockLow      uint
	stockHigh     uint
}

//NewRecommendCreateCmd Create the book into the recommend
func NewRecommendCreateCmd(db *gorm.DB) *cobra.Command {
	var opts recommendCreateOptions
	cmd := &cobra.Command{
		Use:   "create recommendName bookstoreID",
		Short: "Create the recommended area, if bookstoreID = 0, it's set by head office",
		Run: func(cmd *cobra.Command, args []string) {
			opts.recommendName = args[0]
			if b, err := strconv.Atoi(args[1]); err == nil {
				opts.bookstoreID = uint(b)
				runCreate(db, opts)
			}
		},
	}
	flags := cmd.Flags()
	flags.UintVar(&opts.capacity, "capacity", 10, "Add capacity to recommend area")
	flags.UintVar(&opts.dayLow, "day-low", 0, "Add day-low to recommend area")
	flags.UintVar(&opts.dayHigh, "day-high", math.MaxUint16, "Add day-high to recommend area")
	flags.UintVar(&opts.weekLow, "week-low", 0, "Add week-low to recommend area")
	flags.UintVar(&opts.weekHigh, "week-high", math.MaxUint16, "Add week-high to recommend area")
	flags.UintVar(&opts.monthLow, "month-low", 0, "Add month-low to recommend area")
	flags.UintVar(&opts.monthHigh, "month-high", math.MaxUint16, "Add month-high to recommend area")
	flags.UintVar(&opts.stackLow, "stack-low", 0, "Add stack-low to recommend area")
	flags.UintVar(&opts.stackHigh, "stack-high", math.MaxUint16, "Add stack-high to recommend area")
	flags.UintVar(&opts.stockLow, "stock-low", 0, "Add stock-low to recommend area")
	flags.UintVar(&opts.stockHigh, "stock-high", math.MaxUint16, "Add stock-high to recommend area")
	return cmd
}

func runCreate(db *gorm.DB, opts recommendCreateOptions) {
	recommend.Create(db, opts.recommendName, opts.bookstoreID, opts.capacity,
		opts.dayLow, opts.dayHigh,
		opts.weekLow, opts.weekHigh,
		opts.monthLow, opts.monthHigh,
		opts.stackLow, opts.stackHigh,
		opts.stockLow, opts.stockHigh)
}
