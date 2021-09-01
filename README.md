# NodeConfig Operator

An operator to manage node labels, annoations and taints.

Currently we use different node names in OKD4 for each node type, and set different roles and taints based on node roles. However the issue is that we should manually manage these labels and taints, and do not have IaC for them. So if sb changes the role of an edge node or its taints, we don't have any monitoring/reconciler to make sure the pods can still schedule on it or on. For spare nodes, we usually forgot to unschedule or remove the taints from and we end up having many nodes being unscheduled.

To have the node configuration as IaC, there are two operators:

1. https://github.com/barpilot/node-labeler-operator: this one does support node taints and labels, but can only select nodes based on labels, not node name pattern (regex) which we require (we want the label to be set by opertor not manually). The other problem with this operator is that it is deprecated and has no activity since 2018.

2. https://github.com/openshift-kni/node-label-operator: this one does support node pattern, but does not support node taints and annotations. Also it does not work properly and is full of bugs. The desing of this operator is unnecessarily complex (owned labels), and adding taint support probably needs extra CRD. Also this not an actively maintained project and contribution might take some time to be merged. see [this issue](https://github.com/openshift-kni/node-label-operator/issues/13#issuecomment-910010585). `Hi @m-yosefpor, I'm sorry, but we are not maintaining this operator. Our usecase changed before we actually started using it. I will add a note to the readme.`


So it was better and easier for us to develop ourself operator.

## Behavior of operator

- if a node matches multiple NodeConfig, and the same label is defined in both, one randomly will be applied (last one in listing), also for `spec.taint`. It randomly switches between them!!!
- no such concept as owned labels. It only adds/modify labels, no label removal. Also only adds/modifies taints based on keys, no removal. (which is safe and avoid racecondition with other operations)
