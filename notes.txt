
## fleet & fleetctl

> Think of it as a configuration tool for your cluster.

List machines.

```sh
fleetctl list-machines
```

List units.

```sh
fleetctl list-units
fleetctl list-unit-files
```


Can use fleetctl ssh to ssh into a running service.

```sh
fleetctl ssh nginx # ssh into running nginx service
```

Info about the environment: COREOS_PRIVATE_IPV4, COREOS_PUBLIC_IPV4.

```sh
fleetctl ssh nginx cat /etc/environment
```

Displays last 10 logs of the service and follows up on current ones.

```sh
fleetctl journal -f hello.servic
```
## etcd & etcdctl

> Globally distributed key/value
