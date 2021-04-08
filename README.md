# distory
`distory` saves your command histories on each directory.

## install


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

