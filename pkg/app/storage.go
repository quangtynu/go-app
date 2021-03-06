package app

// BrowserStorage is the interface that describes a web browser storage.
type BrowserStorage interface {
	// Set sets the item to the given key. The item must be json convertible in
	// json.
	Set(k string, i interface{}) error

	// Get gets the item associated to the given key and store it in the given
	// value.
	// It returns an error if v is not a pointer.
	Get(k string, v interface{}) error

	// Del deletes the item associated with the given key.
	Del(k string)

	// Clear deletes all items.
	Clear()
}
