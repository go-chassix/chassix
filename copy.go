package chassix

import "github.com/jinzhu/copier"

//Copy copy anything. just proxy to copier(jinzhu/copier).
func Copy(target interface{}, from interface{}) {
	copier.Copy(target, from)
}
