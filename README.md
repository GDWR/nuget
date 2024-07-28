Nuget Client
============
[![Go Reference](https://pkg.go.dev/badge/github.com/gdwr/nuget.svg)](https://pkg.go.dev/github.com/gdwr/nuget)
[![Workflows](https://github.com/GDWR/nuget/actions/workflows/pre-commit.yml/badge.svg)](https://github.com/GDWR/nuget/actions/workflows/pre-commit.yml)

A [nuget](https://www.nuget.org/) client for golang. See [examples](./examples/) to see how it can be used.

Currently only the Catalog API is implemented, as its how I intend to use the API.


```shell
$ go get github.com/gdwr/nuget
```

Contributing
------------

Ensure you utilise the [pre-commit](https://pre-commit.com/) hooks.
```shell
pre-commit install
```


Reference
---------
* [nuget api overview](https://learn.microsoft.com/en-us/nuget/api/overview)
