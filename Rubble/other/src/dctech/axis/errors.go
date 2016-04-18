/*
Copyright 2013-2015 by Milo Christiansen

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented; you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
*/

package axis

type ErrTyp int
const (
	// The path does not point to a valid item.
	ErrNotFound ErrTyp = iota
	
	// The requested action could not be carried out because the item is read-only.
	ErrReadOnly
	
	// The action cannot be done the the item (for example trying to Write a FileSystem).
	ErrBadAction
	
	// An error from an external library, with an attached AXIS path.
	ErrRaw
)

// Create a new AXIS Error with the given type.
func NewError(typ ErrTyp, path *Path) error {
	return &Error{
		Path: path.String(),
		Typ: typ,
	}
}

// WrapError wraps an error from an external library.
func WrapError(err error, path *Path) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*Error); ok {
		return err
	}
	
	return &Error{
		Path: path.String(),
		Typ: ErrRaw,
		Err: err,
	}
}

type Error struct {
	Path string
	Typ ErrTyp
	
	// If Typ == ErrRaw
	Err error
}

func (err *Error) Error() string {
	switch err.Typ {
	case ErrNotFound:
		return "File/Dir Not Found: " + err.Path
	case ErrReadOnly:
		return "File/Dir Read-only: " + err.Path
	case ErrBadAction:
		return "The DataSource does not allow that action: " + err.Path
	case ErrRaw:
		return err.Error() + " AXIS path: " + err.Path
	default:
		return "Invalid Error Code: " + err.Path
	}
}
