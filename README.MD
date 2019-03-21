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
Example: *foobar.go*
```
package main

func foobar() {
  // ...
}

func foobaz() {
  // ...
}
```

## Metadata
### Plugin Metadata
Example: *foobar.json*
```
{
  "services": [
    { "interface": "bar", "symbol": "foobar" },
    { "interface": "baz", "symbol": "foobaz" }
  ]
}
```

### Application metadata
Example: *foo.json*
```
{
  "dependencies": [
    { "interface": "bar", "required": true },
    { "interface": "baz", "required": true }
  ]
}
```
