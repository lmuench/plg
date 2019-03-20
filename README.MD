## Architecture
```
+------------------------------+
|         Application          |
|               +-+----------+ |   +--------+
|               |A|          |---->| plugin |
|               |P|   plg    | |   +--------+
|               |I|          |---->| plugin |
|               +-+ +------+ | |   +--------+
|                 | | gRPC | | |   +--------+
|                 | |Server|<------|plg CLI |
|                 | +------+ | |   +--------+
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
    { "interface": "bar", "name": "foobar" },
    { "interface": "baz", "name": "foobaz" }
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
