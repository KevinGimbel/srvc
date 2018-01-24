# `srvc`

`srvc` is a command line tool which allows you to quickly spin up a webserver with configurable routes. These routes can be configured to return HTML pages, XML, JSON, other files, custom headers, and inline-content directly from the config file.

## How to install

### Homebrew

On MacOS you can get the latest version from `brew`.

```sh
$ brew install kevingimbel/tap/srvc
```

### Binary

Alternatively you can grab the latest binary from the [releases page](https://github.com/kevingimbel/srvc/releases) and place it somewhere in your `$PATH` (e.g. `/usr/local/bin/`).

## Usage

```sh
$ srvc
$ srvc [-port 1313]
```

`-port` is optional and takes a HTTP port to serve to. The default port is 8080.

## Config

`srvc` needs a YAML configuration file in the directory it is executed in. See the sample configuration file in the [example](/example/) directory.

The same configuration file is shown below.

```yaml
# global header config, added to each route
headers:
  - key: "client"
    value: "srvc-alpha1"

# route based config
routes:
# For the route "/demo/html-page" display the "index.html" file from the "html" fodler
  /demo/html-page:
    headers:
      - key: "Content-Type"
        value: "text/html"
    file: "./html/index.html"

  # display HTML content defined inline for /demo/html-inline
  /demo/html-inline:
    headers:
      - key: "Content-Type"
        value: "text/html"
    content: |
      <h1>It works!</h1>
      <p>The content is defined inside the srvc.yaml config file

  # Display XML on /demo/xml
  /demo/xml:
    headers:
      - key: "Content-Type"
        value: "application/xml"
    content: |
      <node id="12">
        <meta charset="utf-8" />
        <link to="#sub" lang="en">
          <title>Go to sub</title>
          <text>Click here</text>
        </link>
      </node>

  # "json/from-file" returns the content of a JSON file
  /json/from-file:
    headers:
      - key: "Content-Type"
        value: "application/json"
    file: "./json/from-file/dirty.json"
```

### Headers

Header can be defined on a global level or for each route. The global headers are added to each route, in the above example each route gets a `"client": "srvc-alpha1"` header.

```yaml
headers:
    - key: "client"
      value: "srvc-alpha1"
```

### Routes

Routes make up the second global config object. Here the different routes are defined. The "key" for each nested object is the route, for example "/hello/world" which makes `srvc` respond on `localhost:8080/hello/world`.

```yaml
routes:
    /hello/world:
        headers:
            - key: "custom"
              value: "header for route /hello/world"      
            - key: "Content-Type"
                value: "text/html"
        content: |
            <h1>Hello, world!</h1>
            <p>This is inline content!</p>
```

This route will respond with a HTML page containing the following code.

```html
<h1>Hello, world!</h1>
<p>This is inline content!</p>
```

A route can be configured to respond with the content of a file, as shown below.

```html
routes:
    /hello/world/file:
        headers:
            - key: "custom"
              value: "header for route /hello/world"      
            - key: "Content-Type"
                value: "text/html"
        file: "./relative/path/to/file.html"
```

The route `/hello/world/file` now responds with the contents of the file located at `./relative/path/to/file.html`.
