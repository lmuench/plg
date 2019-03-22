# :electric_plug:plg

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
```
package plugins

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
```
{
  "services": [
    { "interface": "Greeter", "symbol": "Greet" }
  ]
}
```

### Application metadata
Example: *app.json*
```
{
  "dependencies": [
    { "interface": "Greeter", "required": true }
  ]
}
```
