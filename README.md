# json-config-reader
JSON Config Reader

A simple config reader which is capable of loading multiple JSON files, merging their contents and reading values from the derived strucutre by using a dot seperated key string. Configs are loaded in order, with the information in later configs overriding that of earlier configs. See fixtures for examples.

# Example

```go
    import "github.com/voidhaze/jsoncfg"

    config := new(Jsonconfig)
    config.Loadfiles("config.json", "config.local.json")

    host := config.Get("database.host")
```
