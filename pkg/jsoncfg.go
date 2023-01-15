package jsoncfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ieee0824/go-deepmerge"
)

var master_data, local_data interface{}


// Given a list of file names, read them in order and parse the files into a master data store
func Loadfiles(files ...string) {

	for _, file := range files {
		file_raw, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return

		}

		fmt.Println("Reading : " + file)
		fmt.Println(string(file_raw))

		load_file(file_raw)

	}
}

// Read and load a individual JSON configuration files, given the file name
// decodes the JSON into a internal data struture and merges the new data structure
// with the existing master data structure

func Loadfile(file_data []byte) {
	err := json.Unmarshal(file_data, &local_data)
	if err != nil {
		fmt.Println("Error decoding json:", err)
		return

	}

	if master_data != nil {
		merged_data, err := deepmerge.Merge(master_data, local_data)
		if err != nil {
			fmt.Println("Error merging data:", err)
			return
		}
		master_data = merged_data
	} else {
		master_data = local_data
	}

	fmt.Println(master_data)
}

// Reads a value from the master data structure, given a dot seperated key string
// For example "database.host"
// Supports Arrays and etc...
func Get(get_str string) {
	fragments := strings.Split(get_str, ".")
	datum := master_data
	for _, fragment := range fragments {
		datum = datum.(map[string]interface{})[fragment]
		//check the datum to see if it's a X object or nil
		// throw errors etc...
		fmt.Println(datum)
	}

}
