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
| controller_runtime_active_workers | Number of currently used workers per controller
| controller_runtime_max_concurrent_reconciles | Maximum number of concurrent reconciles per controller
| controller_runtime_reconcile_errors_total | Total number of reconciliation errors per controller
| controller_runtime_reconcile_time_seconds | Length of time per reconciliation per controller
| controller_runtime_reconcile_total | Total number of reconciliations per controller
| rest_client_request_latency_seconds | Request latency in seconds. Broken down by verb and URL.
| rest_client_requests_total | Number of HTTP requests, partitioned by status code, method, and host.
| workqueue_adds_total | Total number of adds handled by workqueue
| workqueue_depth | Current depth of workqueue
| workqueue_longest_running_processor_seconds | How many seconds has the longest running processor for workqueue been running.
| workqueue_queue_duration_seconds | How long in seconds an item stays in workqueue before being requested
| workqueue_retries_total | Total number of retries handled by workqueue
| workqueue_unfinished_work_seconds | How many seconds of work has been done that is in progress and hasn't been observed by work_duration. Large values indicate stuck threads. One can deduce the number of stuck threads by observing the rate at which this increases.
| workqueue_work_duration_seconds | How long in seconds processing an item from workqueue takes.


## Security

### Reporting security vulnerabilities

If you find a security vulnerability or any security related issues, please DO NOT file a public issue, instead send your report privately to cloud@snapp.cab. Security reports are greatly appreciated and we will publicly thank you for it.

## License

Apache-2.0 License, see [LICENSE](LICENSE).
