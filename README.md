# json-config-reader
JSON Config Reader

A simple config reader which is capable of loading multiple JSON files, merging their contents and reading values from the derived strucutre by using a dot seperated key string.

# Example

```
    import "github.com/voidhaze/jsoncfg"

	config := new(Jsonconfig)
	config.Loadfiles("config.json", "config.local.json")

	host := config.Get("database.host")
```
