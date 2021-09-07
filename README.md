# NodeConfig Operator

An operator to manage node labels, annoations and taints.


To have the node configuration as IaC, there are two operators:

1. https://github.com/barpilot/node-labeler-operator: this one does support node taints and labels, but can only select nodes based on labels, not node name pattern (regex). Also it is not an actively maintained project and has no activity since 2018.

2. https://github.com/openshift-kni/node-label-operator: this one does support node pattern, but does not support node taints and annotations. Also this operator is deprectated and not maintained anymore. see [this issue](https://github.com/openshift-kni/node-label-operator/issues/13#issuecomment-910010585).

This operator supports features of both aforementioned operators.

## Behavior of the operator

- If a node matches multiple NodeConfig, and the same label ro `spec.taint` is defined in both, one randomly will be applied (last one in listing).
- No such concept as owned labels. It only adds/modify labels, no label removal. Also only adds/modifies taints based on keys, no removal. (it will be supported in future)

## Instructions

### Development

* `make generate` update the generated code for that resource type.
* `make manifests` Generating CRD manifests.
* `make test` Run tests.

### Build

Export your image name:

```
export IMG=ghcr.io/your-repo-path/image-name:latest
```

* `make build` builds golang app locally.
* `make docker-build` build docker image locally.
* `make docker-push` push container image to registry.

### Run, Deploy
* `make run` run app locally
* `make deploy` deploy to k8s.

### Clean up

* `make undeploy` delete resouces in k8s.

## Configuration


Flags:

```
  -health-probe-bind-address string
        The address the probe endpoint binds to. (default ":8081")
  -kubeconfig string
        Paths to a kubeconfig. Only required if out-of-cluster.
  -leader-elect
        Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.
  -metrics-bind-address string
        The address the metric endpoint binds to. (default ":8080")
  -zap-devel
        Development Mode defaults(encoder=consoleEncoder,logLevel=Debug,stackTraceLevel=Warn). Production Mode defaults(encoder=jsonEncoder,logLevel=Info,stackTraceLevel=Error) (default true)
  -zap-encoder value
        Zap log encoding (one of 'json' or 'console')
  -zap-log-level value
        Zap Level to configure the verbosity of logging. Can be one of 'debug', 'info', 'error', or any integer value > 0 which corresponds to custom debug levels of increasing verbosity
  -zap-stacktrace-level value
        Zap Level at and above which stacktraces are captured (one of 'info', 'error', 'panic').
```

For sample NodeConfig objects, see [config/samples](config/samples) directory.


## Roadmap

- [ ] OwnedLabels
- [ ] Selecting Nodes based on label

## Metrics

| Metric                                              | Notes
|-----------------------------------------------------|------------------------------------
| thanosfederateproxy_scrape_duration_seconds_count   | Total number of scrape requests with response code
| thanosfederateproxy_scrape_duration_seconds_sum     | Duration of scrape requests with response code
| thanosfederateproxy_scrape_duration_seconds_bucket  | Count of scrape requests per bucket (for calculating percentile)

## Security

### Reporting security vulnerabilities

If you find a security vulnerability or any security related issues, please DO NOT file a public issue, instead send your report privately to cloud@snapp.cab. Security reports are greatly appreciated and we will publicly thank you for it.

## License

Apache-2.0 License, see [LICENSE](LICENSE).
