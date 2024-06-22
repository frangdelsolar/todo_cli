package test

import (
	"fmt"
	"strings"

	"github.com/frangdelsolar/todo_cli/pkg/logger"
)

var t *Test

var log *logger.Logger

func init() {
    log = logger.GetLogger()
}

type Running struct {
    Name string
    Count int
}

type Test struct {
    Errors []error
    Current *Running
}

// NewTest creates a new instance of the Test struct.
//
// This function initializes a new instance of the Test struct with an empty
// Errors slice and a Running struct with empty fields. It returns a pointer to
// the newly created Test instance.
//
// Returns:
// - *Test: A pointer to the newly created Test instance.
func NewTest() *Test {
    t = &Test{
        Errors: make([]error, 0),
        Current: &Running{},
    }

    return t
}


// GetTest returns the instance of the Test struct.
//
// This function is a getter method that returns the instance of the Test struct.
// It does not take any parameters.
// It returns a pointer to the Test struct.
func GetTest() *Test {
    return t
}

// AssertEqual checks if two values are equal and logs an error if they are not.
//
// Parameters:
// - a: The first value to compare.
// - b: The second value to compare.
//
// Returns: None.
func (t *Test) AssertEqual (a interface{}, b interface{}) {
    if a != b {
        err := fmt.Errorf("expected %v, got %v", a, b)
        t.Errors = append(t.Errors, fmt.Errorf("expected %v, got %v", a, b))
        
        log.Err(err).Msgf("Assertion failed")
        return
    }

    log.Debug().Msgf("expected %v, got %v. Assertion passed", a, b)
}

// AssertErrorContains checks if the given error contains the expected string.
//
// It takes two parameters:
// - err: The error to be checked.
// - expected: The expected string to be contained in the error.
//
// If the error is nil, it creates a new error with the message "expected error, got nil"
// and appends it to the Errors slice of the Test struct.
// It also logs a failure message.
//
// If the error is not nil and does not contain the expected string,
// it creates a new error with the message "expected error '%s', got '%s'"
// and appends it to the Errors slice of the Test struct.
// It also logs a failure message.
//
// If the error is not nil and contains the expected string,
// it logs a debug message indicating that the assertion passed.
//
// Returns: None.
func (t *Test) AssertErrorContains (err error, expected string) {
    if err == nil {
        err = fmt.Errorf("expected error, got nil")
        t.Errors = append(t.Errors, err)
        log.Err(err).Msgf("Assertion failed")
        return
    }

    if !strings.Contains(err.Error(), expected) {
        err = fmt.Errorf("expected error '%s', got '%s'", expected, err.Error())
        t.Errors = append(t.Errors, err)
        log.Err(err).Msgf("Assertion failed")
        return
    }

    log.Debug().Msgf("expected '%s', got '%s'. Assertion passed", expected, err.Error())
}

// AssertErrorNil checks if the given error is nil.
//
// It takes an error parameter `err` which represents the error to be checked.
// If the error is not nil, it appends the error to the `Errors` slice of the `Test` struct and logs a failure message.
// Otherwise, it logs a debug message indicating that the assertion passed.
//
// Parameters:
// - err: The error to be checked.
//
// Returns: None.
func (t *Test) AssertErrorNil (err error) {
    if err != nil {
        t.Errors = append(t.Errors, err)
        log.Err(err).Msgf("Assertion failed")
        return
    }

    log.Debug().Msgf("expected nil, got nil. Assertion passed")
}

// AssertNotEmpty checks if the given value is not empty.
//
// It takes a parameter `a` of type `interface{}` which represents the value to be checked.
// If the value is `nil`, it appends an error to the `Errors` slice of the `Test` struct and logs a failure message.
// Otherwise, it logs a debug message indicating that the assertion passed.
func (t *Test) AssertNotEmpty (a interface{}) {
    if a == nil {
        err := fmt.Errorf("expected not empty, got empty")
        t.Errors = append(t.Errors, err)
        log.Err(err).Msgf("Assertion failed")
        return
    }

    log.Debug().Msgf("expected not empty, got not empty. Assertion passed")
}

// Run executes the test with the given name.
//
// It updates the current test name and count if the test has not been run before.
// Otherwise, it logs a fatal error message and panics.
//
// Parameters:
// - testName: The name of the test to be executed.
//
// Returns: None.
func (t *Test) Run(testName string) {

    if t.Current.Name == "" {

        t.Current.Name = testName
        t.Current.Count = len(t.Errors)
        return 
    }

    err := fmt.Errorf("trying to run while is running. You may have not closed the previous test with t.Stop()")
    log.Fatal().Interface("Test", t).Err(err).Msgf("Fix your tests! Something is terribly wrong.")
    panic(err)
}

// Stop stops the execution of the current test and logs the result.
//
// It checks if any errors occurred during the test execution and logs the result accordingly.
// If there were errors, it logs an error message. Otherwise, it logs a success message.
// Finally, it resets the current test to an empty state.
//
// No parameters.
// No return values.
func (t *Test) Stop() {
    if len(t.Errors) > t.Current.Count {
        log.Error().Msgf("Failed %s", t.Current.Name)
    } else {
        log.Info().Msgf("Passed %s", t.Current.Name)
    }

    t.Current = &Running{}
}
