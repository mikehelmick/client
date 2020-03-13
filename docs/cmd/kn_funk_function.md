## kn funk function

Creates, lists, and deploys functions.

### Synopsis

Creates, lists, and deploys functions.

```
kn funk function [flags]
```

### Examples

```

 # List all functions
 kn funk function list

 # Create a function
 kn funk create [funcName] -t [CloudEvent Type]

```

### Options

```
  -h, --help   help for function
```

### Options inherited from parent commands

```
      --config string       kn config file (default is ~/.config/kn/config.yaml)
      --kubeconfig string   kubectl config file (default is ~/.kube/config)
      --log-http            log http traffic
```

### SEE ALSO

* [kn funk](kn_funk.md)	 - Functions command group
* [kn funk function create](kn_funk_function_create.md)	 - Create a fun(k) function

