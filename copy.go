package chassis

import "pgxs.io/copier"

//Copy copy anyting. just proxy to copier(jinzhu/copier).
func Copy(toValue interface{}, fromValu interface{}) {
	copier.Copy(toValue, fromValu)
}
