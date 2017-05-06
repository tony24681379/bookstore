package recommendCmd

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/tony24681379/bookstore/recommend"
)

type recommendDeleteOptions struct {
	recommendID uint
	bookstoreID uint
}

//NewRecommendDeleteCmd Delete the book into the recommend
func NewRecommendDeleteCmd(db *gorm.DB) *cobra.Command {
	var opts recommendDeleteOptions
	cmd := &cobra.Command{
		Use:   "delete recommendID bookstoreID",
		Short: "Delete the recommended area",
		Run: func(cmd *cobra.Command, args []string) {
			if r, err := strconv.Atoi(args[0]); err == nil {
				if b, err := strconv.Atoi(args[1]); err == nil {
					opts.recommendID = uint(r)
					opts.bookstoreID = uint(b)
					runDelete(db, opts)
				}
			}
		},
	}
	return cmd
}

func runDelete(db *gorm.DB, opts recommendDeleteOptions) {
	recommend.Delete(db, opts.recommendID, opts.bookstoreID)
}
