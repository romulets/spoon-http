package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUriOnly(t *testing.T) {
	rawParams := []string{"http://localhost"}

	uri, flags, problem := ParseParameters(rawParams)

	assert.Equal(t, "http://localhost", uri, "url is properly returned")
	assert.Equal(t, map[string]string{}, flags, "should not have parameters")
	assert.Nil(t, problem)
}

func TestIgnoreFileName(t *testing.T) {
	rawParams := []string{"INPUT FILE NAME", "http://localhost"}

	uri, flags, problem := ParseParameters(rawParams)

	assert.Equal(t, "http://localhost", uri, "url is properly returned")
	assert.Equal(t, map[string]string{}, flags, "should not have parameters")
	assert.Nil(t, problem)
}

func TestParseParametersFlagsCorrectly(t *testing.T) {
	rawParams := []string{"INPUT FILE NAME", "-X", "POST", "-c", "json", "http://localhost"}

	uri, flags, problem := ParseParameters(rawParams)

	assert.Equal(t, "http://localhost", uri, "url is properly returned")

	assert.Equal(t, map[string]string{
		"X": "POST",
		"c": "json",
	}, flags, "Should ParseParameters flags correctly")

	assert.Nil(t, problem)
}

func TestWrongFlagFormat(t *testing.T) {
	rawParams := []string{"INPUT FILE NAME", "withoutDash", "POST", "-c", "json", "http://localhost"}

	uri, flags, problem := ParseParameters(rawParams)

	assert.Empty(t, "", uri)
	assert.Nil(t, flags)
	assert.Equal(t, errors.New("Argument 'withoutDash' is invalid (must start with '-')"), problem, "should have a problem with flag formating")
}

func TestFlagWithoutFormat(t *testing.T) {
	rawParams := []string{"INPUT FILE NAME", "-v", "-c", "json", "-p", "http://localhost"}

	uri, flags, problem := ParseParameters(rawParams)

	assert.Equal(t, "http://localhost", uri, "url is properly returned")

	assert.Equal(t, map[string]string{
		"v": "",
		"c": "json",
		"p": "",
	}, flags, "Should ParseParameters silent flags correctly")

	assert.Nil(t, problem)
}
