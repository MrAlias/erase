# erase

When [a multi-module Go repository] has moved or deleted a module, without deprecation or retraction, the module becomes abandoned.
Abandoned modules still exist in documentation and are confusing to users, especially if they are similarly named to their replacement.

How should a project handle these abandoned modules?

## Goals

- Code that currently compiles, continues to compile.
- Upgrades to code that use the module inform the upgrader to use an alternate solution.
  - Go tooling should warn the user when they try to build with the module.
- It is clear to new users the module should not be used.
  - Go documentation should clearly state the module should not be used, or it should not contain a listing for the module.
- Support Go 1.17 and above.

## Go module directives

The Go module system has two relevant directives: [deprecation] and [retraction]

### Deprecation

> A module can be marked as deprecated in a block of comments containing the string `Deprecated`: (case-sensitive) at the beginning of a paragraph.
> [...]
> Deprecation messages are intended to inform users that the module is no longer supported and to provide migration instructions, for example, to the latest major version.
> Individual minor and patch versions cannot be deprecated; `retract` may be more appropriate for that.

> Since Go 1.17, `go list -m -u` checks for information on all deprecated modules in the build list.
> `go get` checks for deprecated modules needed to build packages named on the command line.

### Retraction

> A retract directive indicates that a version or range of versions of the module defined by go.mod should not be depended upon.
> A retract directive is useful when a version was published prematurely or a severe problem was discovered after the version was published.
> Retracted versions should remain available in version control repositories and on module proxies to ensure that builds that depend on them are not broken.

> When a module version is retracted, users will not upgrade to it automatically using `go get`, `go mod tidy`, or other commands.
> Builds that depend on retracted versions should continue to work, but users will be notified of retractions when they check for updates with `go list -m -u` or update a related module with `go get`.

## Investigation

### Starting point

![20220527_110500](https://user-images.githubusercontent.com/5543599/170767962-5506307b-6965-4c29-b98e-2b138a7e5ad0.png)

Original respository structure:

```
erase
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── mushroom
│   ├── go.mod
│   └── mushroom.go
└── README.md
```

[a multi-module Go repository]: https://github.com/open-telemetry/opentelemetry-go
[deprecation]: https://go.dev/ref/mod#go-mod-file-module-deprecation
[retraction]: https://golang.org/ref/mod#go-mod-file-retract
