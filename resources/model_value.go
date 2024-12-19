/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Value struct {
	Key
	Attributes ValueAttributes `json:"attributes"`
}
type ValueResponse struct {
	Data     Value    `json:"data"`
	Included Included `json:"included"`
}

type ValueListResponse struct {
	Data     []Value  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustValue - returns Value from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustValue(key Key) *Value {
	var value Value
	if c.tryFindEntry(key, &value) {
		return &value
	}
	return nil
}
