# GoUtils
Library with a set of tools for easy creating software in Golang.

## How to use

### Configs

#### Init config version 1
```go
config := configs.Config{Name: "cfgName"}
config.Init(map[string]any{
"key1": "val1",
"key2": 2,
"key3": true,
}, 1)
```
#### Init config version 2.
```go
config := configs.Config{Name: "cfgName"}
config.Init(map[string]any{
"key1": "val1",
"key2": 2,
"key3": true,
"key4": false,
}, 2)
```

After changing the configuration version, the necessary parameters will be added automatically.

#### Exists key
```go
if config.Exists("key1") {
	//TODO
}
```

#### Get key
```go
val, exists := config.Get("key1")
if exists {
	todo(val.(string))
}
```

#### Set key
```go
config.Set("key2", 5)
config.Save()
```