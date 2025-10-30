# Update Dependencies

## Steps to Update Dependencies

1. `go get -u`
2. Remove all the dependencies packages that has `// indirect` from the modules
3. `go mod tidy`

## Discovering available updates

List all of the modules that are dependencies of your current module, along with the latest version available for each:
```zsh
go list -m -u all
```

Display the latest version available for a specific module:

```zsh
go list -m -u example.com/theirmodule
```

**Example:**

```zsh
go list -m -u cloud.google.com/go/firestore
cloud.google.com/go/firestore v1.2.0 [v1.6.1]
```

## Getting a specific dependency version

To get a specific numbered version, append the module path with an `@` sign followed by the `version` you want:

```zsh
go get example.com/theirmodule@v1.3.4
```

To get the latest version, append the module path with @latest:

```zsh
go get example.com/theirmodule@latest
```

## Synchronizing your code’s dependencies

```zsh
go mod tidy
```
