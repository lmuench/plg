## Architecture
```
+------------------------------+
|         Application          |
|               +-+----------+ |   +--------+
|               |A|   plg    |---->| plugin |
|               |P| service  | |   +--------+
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
    { "interface": "bar", "function": "foobar" },
    { "interface": "baz", "function": "foobaz" }
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
