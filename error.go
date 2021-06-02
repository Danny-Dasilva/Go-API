package main

import (
    "errors"
    "fmt"
)

type WrappedError struct {
    StatusCode int
    Context string
    Err     error
}

func (r *WrappedError) Temporary() int {
    return r.StatusCode
}
func (w *WrappedError) Error() string {
    return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info string) *WrappedError {
    return &WrappedError{
        StatusCode: 503,
        Context: info,
        Err:     err,
    }
}

func main() {
    err := errors.New("boom!")
    err = Wrap(err, "main")

    fmt.Println(err)
    re, ok := err.(*WrappedError)
    _=ok
    fmt.Println(re.Temporary())
}
