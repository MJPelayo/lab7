package main

import "errors"

// 🎯 validate student
func validateStudent(s Student) error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Programme == "" {
		return errors.New("programme is required")
	}
	if s.Year <= 0 {
		return errors.New("year must be greater than 0")
	}
	return nil
}

// 🎯 validate course
func validateCourse(c Course) error {
	if c.Code == "" {
		return errors.New("code is required")
	}
	if c.Title == "" {
		return errors.New("title is required")
	}
	if c.Department == "" {
		return errors.New("department is required")
	}
	if c.Instructor == "" {
		return errors.New("instructor is required")
	}
	if c.Credits <= 0 {
		return errors.New("credits must be > 0")
	}
	if c.Capacity <= 0 {
		return errors.New("capacity must be > 0")
	}
	return nil
}
