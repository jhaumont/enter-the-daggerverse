# Use module from Daggerverse with Dagger Shell

## Prerequierement

You must know what is Daggerverse. If don't, please read [Use module from Daggerverse](./03-use-module-from-daggerverse.md) before continue.

## Install a module - Git info

To install a Dagger module, first init Dagger env in `hello-sh` folder:
```
# To install module, we need SDK. Bug will be fix soon
dagger init --sdk=go --source=./dagger
```

Then, install a module:
```
dagger install github.com/vbehar/daggerverse/git-info@v0.12.1
```

Test if module is loaded:
```
dagger <<< .help
✔ connect 0.2s
✔ load module 0.7s
✔ serving dependency modules 0.0s

  container-echo   Returns a container that echoes whatever string argument is provided
  grep-dir         Returns lines that match a pattern in the files of the provided Directory

  git-info      A Dagger Module to extract information about a git reference.
```

You see last line git-info is loaded.

Now get current git ref:
```
dagger <<< "git-info . | ref"
✔ connect 0.2s
✔ load module 0.7s
...
✔ .ref: String! 0.0s

HEAD
```

It works!