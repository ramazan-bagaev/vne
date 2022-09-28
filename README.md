# vne
Tool for transferring dev environment from one machine to another

## install
```go build```

## usage
```vne create -u $username -d $vne-config-location #creates user with name $username and password "pass" ```

```vne delete -u $username #deletes user with name $username```

```vne load -d $vne-config-location #load envs, bins and shell to specified config from system```

```vne unload -d $vne-config-location #set envs in system, try to load bin executables specified```

## tags
[env] - env variables

[tools] - content of bin directories specified in the $PATH

[shell] - shell location (/bin/bash; /bin/zsh ...)

[dirs] - empty dirs to create
