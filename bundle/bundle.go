package bundle

import (
	"fmt"
)

//Add book into bundle
func Add(bundleID string, bookID int) {
	fmt.Println("add", bundleID, bookID)
}

//Create the bundle
func Create(bundleID string, note string) {
	fmt.Println("create", bundleID, note)
}

//Delete the bundle
func Delete(bundleID string) {
	fmt.Println("delete", bundleID)
}

//Modify the bundle
func Modify(bundleID string, note string) {
	fmt.Println("Modify", bundleID, note)
}

//Remove the bundle
func Remove(bundleID string, bookID int) {
	fmt.Println("Remove", bundleID, bookID)
}
