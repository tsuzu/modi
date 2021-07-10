# modi
Initialize go.mod automatically

## Installation
```
$ go install github.com/tsuzu/modi@latest
```

## Usage
Enter the directory to put go.mod in and run modi
```console
$ modi
```

If `.git` is initialized, module will be calculated from `origin`.
If not, user/orgs are retrieved from `gh` CLI, and you can select one from them.
