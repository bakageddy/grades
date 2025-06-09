package extractors

import (
	"errors"
	"net/url"
	"strconv"
)

type AddEntryRequest struct {
	RegisterNo uint64
	Credits uint64
	Grade string
	SubjectCode string
}

func (a *AddEntryRequest) Extract(parameter_map url.Values) error {
	if (!parameter_map.Has("register_no")) {
		return errors.New("`register_no` parameter unavailable")
	}

	if (!parameter_map.Has("credits")) {
		return errors.New("`credits` parameter unavailable")
	}

	if (!parameter_map.Has("grade")) {
		return errors.New("`grade` parameter unavailable")
	}

	if (!parameter_map.Has("subject_code")) {
		return errors.New("`subject_code` parameter unavailable")
	}

	register_no, err := strconv.ParseUint(parameter_map.Get("register_no"), 10, 64)
	if err != nil {
		return err
	}

	credits, err := strconv.ParseUint(parameter_map.Get("credits"), 10, 64)
	if err != nil {
		return err
	}

	grade := parameter_map.Get("grade")
	subject_code := parameter_map.Get("subject_code")

	a.RegisterNo = register_no
	a.Grade = grade
	a.Credits = credits
	a.SubjectCode = subject_code

	return nil
}
