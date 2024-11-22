<div style="text-align: center">

# go-airzone

> go-airzone is a Golang client library for Airzone Local API.

</div>

---

## Table of Contents

<!-- TOC -->
* [go-airzone](#go-airzone)
  * [Table of Contents](#table-of-contents)
  * [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
    * [Usage](#usage)
      * [Creating a Client](#creating-a-client)
  * [Contributing](#contributing)
  * [License](#license)
<!-- TOC -->

---


## Getting Started

### Prerequisites

* Go environment
* Airzone system access

### Installation

```bash
go get -u github.com/merlindorin/go-airzone
```

### Usage

To use `go-airzone` in your project, you must import the library and use its structures and methods.

#### Creating a Client

Before you can interact with the Airzone local API, you need to create a new client. The following code demonstrates how
to instantiate a client:

```go
import (
    "net/url"
    "github.com/merlidorin/go-airzone"
)

func main() {
  l, err := zap.NewDevelopment(zap.IncreaseLevel(level))
  if err != nil {
    panic(err)
  }
  
  u, err := url.Parse('http://192.168.1.235:3000')
  if err != nil {
  panic(err)
  }
  
  cl := pkg.NewClient(u, l)
}
```

## Contributing

Feel free to fork the repository, make changes, and submit pull requests. Your contributions make the open-source
community thrive.

## License

This project is released under the MIT License - see the [`LICENSE`](./LICENSE.md) file for details.
