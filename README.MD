# :electric_plug:plg
:construction: *Pre-alpha stage* :construction:
## Architecture
```
+------------------------------+
|         Application          |
|               +-+----------+ |   +--------+
|               |A|   plg    |---->| plugin |
|               |P|  symbol  | |   +--------+
|               |I| registry |---->| plugin |
|               +-+ +------+ | |   +--------+
|                 | | gRPC | | |   +---------+
|                 | |Server|<------| plg CLI |
|                 | +------+ | |   +---------+
|                 +----------+ |
+------------------------------+
```

## Plugins
Example: *greet.go*
```go
package main

import "fmt"

type greeter struct{}

func (g greeter) Greet(name string) {
	fmt.Println("Hello " + name)
}

// Greeter will be exported as symbol "Greeter"
var Greeter greeter
```

## Metadata
### Plugin Metadata
Example: *greet.json*
```json
{
  "services": [
    { "interface": "Greeter", "symbol": "Greet" }
  ]
}
```

### Application metadata
Example: *app.json*
```json
{
  "dependencies": [
    { "interface": "Greeter", "required": true }
  ]
}
```
