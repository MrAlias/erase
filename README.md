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

Released as `v0.0.2`

### Move `mushroom` -> `mushrooms`

New repository structure in release `v0.0.3`:

```
erase
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── mushrooms
│   ├── go.mod
│   └── mushroom.go
└── README.md
```

Now the Go docs page lists two similar modules.

![20220527_123433](https://user-images.githubusercontent.com/5543599/170778527-485be92f-825d-4807-9334-20d1072cb8bc.png)

Old code that depends on `github.com/MrAlias/erase/mushroom` still compiles, but it is not obvious to new users which module to use without deeper investigation about each of their versions.

The old abandoned module, `mushroom`, needs to be cleaned up.

### Deprecate `mushrooms`

Communicate to the user that the `mushroom` module is no longer supported and to provide migration instructions to start using `mushrooms` by deprecating the package.

Ultimately this means ...

- update `mushroom/go.mod` to include a deprecation notice
- release an incremented version of `mushroom` to inclued the deprecation notice

The only issue is, the `main` branch no longer contains `mushroom/go.mod`.
To accomplish this a branch needs to be created from the commit prior to its removal.

```sh
git branch erase-mushroom bfd719a~1
```

Now update `mushroom/go.mod`.

```sh
sed -i.old '1s;^;// Deprecated: use github.com/MrAlias/erase/mushrooms instead.\n;' mushroom/go.mod
git commit -m "Deprecate github.com/MrAlias/erase/mushroom" mushroom/go.mod
git push origin erase-mushroom
git tag -s -a -m "Release github.com/MrAlias/erase/mushroom v0.0.3" mushroom/v0.0.3
git push origin mushroom/v0.0.3
```

There is now a new `mushroom/v0.0.3` tag in the repository!

![20220527_132940](https://user-images.githubusercontent.com/5543599/170785250-7d204e72-9ffb-463e-9b6b-9b9e689c43be.png)

And similarly the package documentation lists `github.com/MrAlias/erase/mushroom` as deprecated.

![20220527_133110](https://user-images.githubusercontent.com/5543599/170785481-91812348-54ba-49e5-ab19-6806323bc1c9.png)

The last thing to do for this is to protect the `erase-mushroom` branch.
It should not be allowed to be deleted, otherwise the source of the `v0.0.3` version of module will not available.

Create an empty branch protection rule in GitHub.

![20220527_135002](https://user-images.githubusercontent.com/5543599/170787801-5f05bee5-dfe9-44b9-b144-b30aa792d842.png)

By default, this protection rule will prevent the `erase-mushroom` branch from being deleted.

[a multi-module Go repository]: https://github.com/open-telemetry/opentelemetry-go
[deprecation]: https://go.dev/ref/mod#go-mod-file-module-deprecation
[retraction]: https://golang.org/ref/mod#go-mod-file-retract
