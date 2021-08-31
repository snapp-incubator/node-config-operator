# NodeConfig Operator


- if a node matches multiple NodeConfig, and the same label is defined in both, one randomly will be applied (last one in listing), also for `spec.taint`. It randomly switches between them!!!
- no such concept as owned labels. It only adds/modify labels. No label removal. Also only adds taints (which is safe and avoid racecondition with other operations)
