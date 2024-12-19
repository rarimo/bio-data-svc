/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AddValue struct {
	Key
	Attributes AddValueAttributes `json:"attributes"`
}
type AddValueRequest struct {
	Data     AddValue `json:"data"`
	Included Included `json:"included"`
}

type AddValueListRequest struct {
	Data     []AddValue `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustAddValue - returns AddValue from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAddValue(key Key) *AddValue {
	var addValue AddValue
	if c.tryFindEntry(key, &addValue) {
		return &addValue
	}
	return nil
}
