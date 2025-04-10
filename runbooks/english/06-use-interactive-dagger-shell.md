# Introducing Dagger Shell

The Mar 26, 2025, Dagger.io's team [announce](dagghttps://dagger.io/blog/a-shell-for-the-container-age-introducing-dagger-shell) an amazing feature: Dagger Interractive Shell.

This feature is very usefull to test Dagger.io and allow Dagger to be used by non-developper people.

Dagger shell is not a shell like Bash, Zsh...It is more like [jq](https://jqlang.org) syntax.

Let's go started.

## Launch interactive shell

To start Dagger shell, just type:
```
❯ dagger
Dagger interactive shell. Type ".help" for more information. Press Ctrl+D to exit.
⋈
─ esc nav mode · > run prompt ────────────────────────────────────────────────────
```

You can exit Dagger shell by using `Ctrl+C`, `Ctrl+D` or `exit`.

# Get help

In Dagger shell, you can use `.help` command to display help:
```
✔ .help 0.0s
  cache-volume   Constructs a cache volume for a given cache key.                                                                                                                                                                           
  container      Creates a scratch container.                                                                                                                                                                                               
  directory      Creates an empty directory.
...
```

You can have information on module or command. See [Built-in help](https://docs.dagger.io/features/shell/#built-in-help) section of official documentation.

# Run first command

Now, you can enter any Dagger command.

Enter this command:
```
container | from alpine | with-exec whoami
```

You tell to Dagger to run a container based on alpine Docker image and run commande `whoami`. But wait, you have this output:
```
✔ container | from alpine | with-exec whoami 0.0s
Container@xxh3:83a64a1566b1bdba
```

You don't see the result of `whoami` command!

It's normal, you must ask to Dagger shell to get output of command and display it.

Another example with `echo 'hello'` and print stdout:
``` 
container | from alpine | with-exec -- sh -c "echo hello"
``` 

Now, you have result of `echo` command:
```
✔ container | from alpine | with-exec -- sh -c "echo hello" | stdout 0.0s
hello
```

By adding `| stdout`, Dagger shell will print stdout of container on stdout of terminal.

# Run Dagger shell in... shell 😮

You can run script your Dagger shell from your unix shell, like `cat` command:
```
dagger <<EOF
container \
| from alpine \
| with-exec -- sh -c "echo hello" \
| stdout
EOF
```

The result:
```
✔ connect 0.2s
✔ loading type definitions 0.2s

✔ container: Container! 0.0s
$ .from(address: "alpine"): Container! 0.6s CACHED
$ .withExec(args: ["sh", "-c", "echo hello"]): Container! 0.0s CACHED
✔ .stdout: String! 0.0s

hello
```

Or:
```
echo 'container | from alpine | with-exec -- sh -c "echo hello" | stdout' | dagger
```

Or by storing command in file:
```
cat > hello.dsh <<EOF
container \
| from alpine \
| with-exec -- sh -c "echo hello" \
| stdout
EOF

dagger hello.dsh
```

# Mount a folder or file in container

Now, we want create a file in current dir from container:
```
dagger <<EOF
container \
| from alpine \
| with-exec -- sh -c "echo titi > /data/test" \
| with-directory /data/ .
EOF
```

You got this error message:
```
Error: input: container.from.withExec.withDirectory.id process "sh -c echo titi > /data/test" did not complete successfully: exit code: 1
```

because, order of parameters is important. Swap lines `with-directory` and `with-exec`:
```
dagger <<EOF
container \
| from alpine \
| with-directory /data/ . \
| with-exec -- sh -c "echo titi > /data/test"
EOF
```

If you got to the `/workspaces/enter-the-daggerverse/hello-sh` folder, file `test` is not here.
You must `export` it:
```
dagger <<EOF
container \
| from alpine \
| with-directory /data/ . \
| with-exec -- sh -c "echo titi > /data/test" \
| export ./test
EOF
```
