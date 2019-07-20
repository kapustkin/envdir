package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnviroment1(t *testing.T) {
	res, err := getEnviroment("../test/1")
	assert.Nil(t, err)
	assert.Equal(t, res, []string{"A_ENV=123", "B_VAR=another_val"})
}

func TestGetEnviroment2(t *testing.T) {
	res, err := getEnviroment("../test/2")
	assert.NotNil(t, err)
	assert.Equal(t, res, []string(nil))
}

func TestGetEnvParametr1(t *testing.T) {
	res, err := getEnvParametr("../test/1/A_ENV.txt")
	assert.Nil(t, err)
	assert.Equal(t, res, "A_ENV=123")
}

func TestGetEnvParametr2(t *testing.T) {
	res, err := getEnvParametr("../test/1/NO_EXIST.txt")
	assert.NotNil(t, err)
	assert.Equal(t, res, "")
}

func TestFileNameWithoutExtension(t *testing.T) {
	res := fileNameWithoutExtension("../test/1/A_ENV.txt")
	assert.Equal(t, res, "A_ENV")
}

func TestGetInputValues1(t *testing.T) {
	env, path, err := getInputValues([]string{"1233", "321"})
	assert.Equal(t, env, "1233")
	assert.Equal(t, path, "321")
	assert.Nil(t, err)
}

func TestGetInputValues2(t *testing.T) {
	env, path, err := getInputValues([]string{})
	assert.Equal(t, env, "")
	assert.Equal(t, path, "")
	assert.NotNil(t, err)
}
