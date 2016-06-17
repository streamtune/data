package data

// Pageable is the struct holding the page requests
type Pageable struct {
	Page int
	Size int
	Sort *Sort
}

// NewPageable creates a new Pageable instance for provided page and size
func NewPageable(page, size int) *Pageable {
	return NewSortedPageable(page, size, nil)
}

// NewSortedPageable creats a new Pageable instance for provided page, size and Sort object
func NewSortedPageable(page, size int, sort *Sort) *Pageable {
	return &Pageable{page, size, sort}
}

// Offset returns the offset from start
func (p *Pageable) Offset() int {
	return p.Page * p.Size
}

// HasPrevious check if the Pageable has previous page or not
func (p *Pageable) HasPrevious() bool {
	return p.Page > 0
}

// PreviousOrFirst will return a new Pageable for the previous page or the first page if no previous page exists
func (p *Pageable) PreviousOrFirst() *Pageable {
	if p.HasPrevious() {
		return p.Previous()
	}
	return p.First()
}

// Next creates a new Pageable for the next result page
func (p *Pageable) Next() *Pageable {
	return &Pageable{p.Page + 1, p.Size, p.Sort}
}

// Previous creates a new Pageable for the previous result page
func (p *Pageable) Previous() *Pageable {
	if p.Page == 0 {
		return p
	}
	return &Pageable{p.Page - 1, p.Size, p.Sort}
}

// First creates a new Pageable for the first result page
func (p *Pageable) First() *Pageable {
	return &Pageable{0, p.Size, p.Sort}
}
