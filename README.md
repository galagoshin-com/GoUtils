# GoUtils
Library with a set of tools for easy creating software in Golang.

## Install

`go get -t github.com/Galagoshin/GoUtils`

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

### Crypto

#### HashDir
```go
lastDirHash, err := crypto.HashDir("path/to/dir", "prefix", crypto.Hash1) //returns string
```

#### MD5
```go
md5 := crypto.Md5([]byte("example")) //returns string
```

#### SHA
```go
sha1 := crypto.Sha1([]byte("example")) //returns string
sha256 := crypto.Sha256([]byte("example")) //returns string
sha512 := crypto.Sha512([]byte("example")) //returns string
```

### Encoding

#### Base64

```go
b64 := encoding.EncodeBase64([]byte("example")) //returns Base64 (string)
dcd := encoding.DecodeBase64(b64) //returns []byte
```

### Events

#### Register event handler

```go
events.RegisterEvent(events.Event{
    Name: events.EventName("AnyEvent"), 
    Execute: OnEvent,
})

func OnEvent(args ...any) {
    first_arg := args[0].(bool)
    if first_arg {
        fmt.Println("Event executed!")
    }
}
```

#### Calling event

```go
arg := true
events.CallAllEvents(events.EventName("AnyEvent"), arg)
```

### Files

#### Example usage

```go
file := files.File{Path: "path/to/file"}
if file.Exists() {
    err := config.file.Create()
    if err != nil {
        panic(err)
    }
    
    err = file.WriteString("data")
    if err != nil {
        panic(err)
    }

    err = file.Close()
    if err != nil {
    	panic(err)
    }

    err := file.Open(os.O_RDWR)
    if err != nil {
        panic(err)
    }
    content := config.file.ReadString() //returns content from file

    err = file.Close()
    if err != nil {
        panic(err)
    }
}
```

Use `Read()` and `Write()` for `[]byte` manipulations.

#### Check is directory

```go
if file.IsDir() {
	//TODO
}
```

#### Check is file

```go
if file.IsFile() {
	//TODO
}
```

#### Create directory

```go
dir := files.Directory{Path: "path/to/dir"}
err := dir.Create()
```
or
```
err := dir.CreateAll()
```

#### Remove directory
```go
dir := files.Directory{Path: "path/to/dir"}
err := dir.Remove()
```

### Json

#### Json string to `map[string]any`
```go
mp, err := json.Decode(json.Json("{\"key\": \"value\"}"))
if err == nil {
	todo(mp["key"].(string)) 
}
```

#### `map[string]any` to json string
```go
mp = map[string]any{
	"key": "value",
}
jsn, err := json.Encode(mp)
if err == nil {
	todo(string(jsn))
}
```

### Random

#### Example usage
```go
random.SetSeed(1) //Automatically sets to time.Now().UnixNano()
rand_int := random.RangeInt(0, 10) //returns random int from 0 to 10
```

### Requests

#### Example usage
```go
request := requests.Request{
    Method: requests.POST,
    Url: requests.URL("http://example.com"),
    Data: url.Values{
        "key1": {"val"},
        "key2": {"val"},
    },
    Headers: http.Header{
        "key1": {"val"},
        "key2": {"val"},
    },
    Cookies: []*http.Cookie{
        {Name: "name1", Value: "value"},
        {Name: "name2", Value: "value"},
    },
    Timeout: time.Second,
}
response, err := request.Send()
if err == nil {
    fmt.Println(response.Text())
}
```

### Schedulers

#### RepeatingTask
```go
var task = scheduler.RepeatingTask{
	Duration:   time.Second,
	OnComplete: TaskExecutor,
}

func TaskExecutor(args ...any) {
	destroy := args[0].(bool)
	if destroy {
		task.Destroy()
	}else{
		fmt.Println("Task executed!")
	}
}

task.Run(false)
```

#### OneTimeTask
```go
var task = scheduler.OneTimeTask{
	Duration:   time.Second,
	OnComplete: TaskExecutor,
}

func TaskExecutor(...any) {
    fmt.Println("Task executed!")
}

task.Run()
```

#### UnixTimeTask
```go
task := scheduler.UnixTimeTask{}
task.SetKey("name", time.Second * 10)
for i := 0; i < 20; i++ {
    time_left, exists := task.GetTimeLeft("name")
    if exists {
        if time_left <= 0 {
            fmt.Println("Task done for 'name'!")
            task.RemoveKey("name")
        }else{
            fmt.Println(time_left, " left.")
        }
    }
    time.Sleep(time.Second)
}
```

### Strings

#### Check is numeric
```go
if strings.IsNumeric("123") {
	//TODO
}
```

#### Calculate distance
```go
strings.Distance("hello", "helli") //returns 1
strings.Distance("hello", "helii") //returns 2
```

### Time

#### Measure execution
```go
logger.Print(fmt.Sprintf("Function done (%f s.)", time.MeasureExecution(func() {
    //TODO
})))
```