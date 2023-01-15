package jsoncfg

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadfiles(t *testing.T) {

	// basic load and run, this also tests that master data merging is functioning correctly
	config := new(Jsonconfig)
	config.Loadfiles("fixtures/config.invalid.json", "fixtures/config.also_invalid.json", "fixtures/config.json", "fixtures/config.local.json")
	localhost := config.Get("database.host")

	assert.Equal(t, "127.0.0.1", localhost, "the expected host was fetched from the merged config")

	// load singular
	config2 := new(Jsonconfig)
	config2.Loadfiles("fixtures/config.json")
	localhost2 := config2.Get("database.host")

	assert.Equal(t, "mysql", localhost2, "the expected host was fetched from the config")

	// load invalid file
	config3 := new(Jsonconfig)
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	config3.Loadfiles("fixtures/config.invalid.json")

	assert.Contains(t, buf.String(), "Error decoding json:", "invalid JSON file was found")
	t.Log(buf.String())

}

func TestLoadfile(t *testing.T) {

	// load valid
	config := new(Jsonconfig)

	raw := []byte(`{
		"environment": "production",
		"database": {
		  "host": "mysql",
		  "port": 3306,
		  "username": "divido",
		  "password": "divido"
		},
		"cache": {
		  "redis": {
			"host": "redis",
			"port": 6379
		  }
		}
	  }`)

	config.Loadfile(raw)

	localhost := config.Get("database.host")

	assert.Equal(t, "mysql", localhost, "the expected host was fetched from the config")

	// load invalid
	config2 := new(Jsonconfig)
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	raw2 := []byte(`WE LIKE TRAINS`)
	config2.Loadfile(raw2)

	assert.Contains(t, buf.String(), "Error decoding json:", "invalid JSON file was found")

	// check master data has been updated correctly
	config3 := new(Jsonconfig)

	raw3 := []byte(`{
		"foo": "bar",
		"object" : { "foo" : "bar"}
	}`)

	config3.Loadfile(raw3)

	assert.Equal(t, "bar", config3.Get("foo"), "the expected string data is fetched from the config")
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, config3.Get("object"), "the expected object data is fetched from the config")
}

func TestGet(t *testing.T) {

	config3 := new(Jsonconfig)

	raw3 := []byte(`{
		"string" : "hello world",
		"number" : 123,
		"object" : { "foo" : "bar"},
		"true" : true,
		"false" : false,
		"array" : [1,2,3],
		"null" : null
	}`)

	config3.Loadfile(raw3)

	assert.Equal(t, "hello world", config3.Get("string"), "the expected string data fetched from the config")
	assert.Equal(t, "bar", config3.Get("object.foo"), "the expected object data fetched from the config")
	assert.Equal(t, true, config3.Get("true"), "the expected true data fetched from the config")
	assert.Equal(t, false, config3.Get("false"), "the expected false data fetched from the config")
	assert.Equal(t, float64(3), config3.Get("array.3"), "the expected array data fetched from the config")
	assert.Equal(t, nil, config3.Get("null"), "the expected false data fetched from the config")
}
