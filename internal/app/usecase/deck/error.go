package deck

import "errors"

var (
    ErrUnexpected            = errors.New("unexpected internal error")
    ErrDeckNotFound          = errors.New("deck not found")
    ErrInvalidDeckCardFilter = errors.New("selected filter doesn't exist at card decks")
)

