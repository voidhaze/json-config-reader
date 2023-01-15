package jsoncfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadfiles(t *testing.T) {

	// basic load and run
	var config Jsonconfig

	config.Loadfiles("fixtures/config.json", "fixtures/config.local.json")

	localhost := config.Get("database.host")

	assert.Equal(t, "127.0.0.1", localhost, "the expected host was fetched from the merged config")

	// load dir
	// load list
	// load singular
	// load invalid file

	// chech master data merging is functioning correctly
}

func TestLoadfile(t *testing.T) {

	// load valid
	// load invalid
	// load relative
	// load absolute

	// check master data has been updated correctly
}

func TestGet(t *testing.T) {
	// get string
	// get array
	// get object
	// get true
	// get false
	// get non existant / error
}
