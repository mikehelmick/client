## kn funk languages

Manage fun(k) installed language SDKs

### Synopsis

Manage fun(k) installed language SDKs

```
kn funk languages [flags]
```

### Examples

```

 # List all known languages
 kn funk languages

 # Install a specific language SDK
 fn funk languages --install go

 # Update a specific language SDK
 fn funk languages --update nodejs

 # Uninstall a specific language SDK
 fn funk languages --uninstall java

```

### Options

```
  -h, --help   help for languages
```

### Options inherited from parent commands

```
      --config string       kn config file (default is $HOME/.kn/config.yaml)
      --kubeconfig string   kubectl config file (default is $HOME/.kube/config)
      --log-http            log http traffic
```

### SEE ALSO

* [kn funk](kn_funk.md)	 - Functions command group

