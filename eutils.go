package eutils

import "errors"

func Err2Str(e error) string {
    if e != nil {
        return e.Error()
    }
    return ""
}

func Str2Err(s string) error {
    if s == "" {
        return nil
    }
    return errors.New(s)
}
