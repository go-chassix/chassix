package chassis

import "github.com/jinzhu/copier"

//Copy copy anyting. just proxy to copier(jinzhu/copier).
func Copy(toValue interface{}, fromValu interface{}) {
	copier.Copy(toValue, fromValu)
}
