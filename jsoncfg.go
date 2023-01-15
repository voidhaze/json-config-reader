package jsoncfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/ieee0824/go-deepmerge"
)

type Jsonconfig struct {
	Masterdata interface{}
}

type Configforge interface {
	Get(string)
	Loadfile(string)
	Loadfiles(...string)
}

// Given a list of file names, read them in order and parse the files into a master data store
func (cfg *Jsonconfig) Loadfiles(files ...string) {

	for _, file := range files {
		file_raw, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println("Error reading file:", err)
			return

		}

		//fmt.Println("Reading : " + file)
		//fmt.Println(string(file_raw))

		cfg.Loadfile(file_raw)

	}
}

// Read and load a individual JSON configuration files, given the file name
// decodes the JSON into a internal data struture and merges the new data structure
// with the existing master data structure

func (cfg *Jsonconfig) Loadfile(file_data []byte) {
	var localdata interface{}
	err := json.Unmarshal(file_data, &localdata)
	if err != nil {
		log.Println("Error decoding json:", err)
		return

	}

	if cfg.Masterdata != nil {

		// merge the data sets so that the data in localdata always overrides masterdata
		merged_data, err := deepmerge.Merge(localdata, cfg.Masterdata)
		if err != nil {
			log.Println("Error merging data:", err)
			return
		}

		cfg.Masterdata = merged_data
	} else {

		cfg.Masterdata = localdata
	}

	//fmt.Println(cfg.Masterdata)
}

// Reads a value from the master data structure, given a dot seperated key string
// For example "database.host" attempts to fetch the value stored in the master data
// structure. Supports Arrays and etc...
func (cfg *Jsonconfig) Get(searchstr string) interface{} {

	fragments := strings.Split(searchstr, ".")

	datum := cfg.Masterdata

	for _, fragment := range fragments {

		// type switch to support arrays
		switch datum.(type) {
		case []interface{}:
			index, err := strconv.Atoi(fragment)
			if err != nil {
				log.Println("Error converting array index to int:", err)
				continue
			}
			datum = datum.([]interface{})[index-1]
		default:
			datum = datum.(map[string]interface{})[fragment]
		}

		//fmt.Println(datum)

	}

	return datum

}
