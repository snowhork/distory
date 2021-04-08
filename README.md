# distory
`distory` saves your command histories on each directory.

## install
```
# requirements: go version >= 1.16.0
go install github.com/snowhork/distory@v0.0.1
```

## how to use
First, you need to set hook `distory -a <COMMAND>` for your using shell. (bash,zsh, etc..)
Then, you can see directory's command history by `distory`.

## set hook
### bash

Install https://github.com/rcaloras/bash-preexec for command hook
```sh
curl https://raw.githubusercontent.com/rcaloras/bash-preexec/master/bash-preexec.sh -o ~/.bash-preexec.sh
```

Edit your `~/.bash_profile`
```sh
source ~/.bash-preexec.sh
preexec() { distory -a "$1"; }
```

