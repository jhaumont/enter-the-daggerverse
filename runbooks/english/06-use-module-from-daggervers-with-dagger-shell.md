# Use module from Daggerverse with Dagger Shell

## Install a module - Git info

To install a Dagger module, first init Dagger env in `hello-sh` folder:
```
# To install module, we need SDK. It's a bug and willl be fixed soon  
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

The last line shows that `git-info` is loaded.  

Now use this module, to get current git ref:
```
dagger <<< "git-info . | ref"
✔ connect 0.2s
✔ load module 0.7s
...
✔ .ref: String! 0.0s

HEAD
```

It works 🎉

## Dagger module

We have play with basics of Dagger. To go further, you can [Create, publish a module in Daggerverse and use it](07-create-publish-module.md) with Go.