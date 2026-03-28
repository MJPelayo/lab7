package main

import "errors"

func validateStudent(s Student) error {
	if s.Name == "" {
		return errors.New("name required")
	}
	if s.Programme == "" {
		return errors.New("programme required")
	}
	if s.Year <= 0 {
		return errors.New("invalid year")
	}
	return nil
}

func validateCourse(c Course) error {
	if c.Code == "" || c.Title == "" || c.Department == "" || c.Instructor == "" {
		return errors.New("missing fields")
	}
	if c.Credits <= 0 || c.Capacity <= 0 {
		return errors.New("invalid values")
	}
	return nil
}
