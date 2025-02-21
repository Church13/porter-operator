---
title: Install the Porter Operator
description: Get up and running with the Porter Operator
---

If you aren't already familiar with Porter, we recommend that you [install the Porter v1 prerelease] first and then once you are comfortable, learn how to automate Porter with the operator.

The Porter Operator is installed with ... Porter!
First, use explain to see what credentials and parameters you can use when installing and configuring the operator.

```
$ porter explain -r ghcr.io/getporter/porter-operator:canary
Name: porter-operator
Description: The Porter Operator for Kubernetes. Execute bundles on a Kubernetes cluster.
Version: 1.0.0-alpha.1
Porter Version: v1.0.0-alpha.5

Credentials:
Name         Description                                                          Required   Applies To
kubeconfig   Kubeconfig file for cluster where the operator should be installed   true       All Actions

Parameters:
Name                         Description                                                                                                                                                                                                                                                 Type     Default        Required   Applies To
installationServiceAccount   Name of the service account to run installation with. If set, you are responsible for creating this service account and giving it required permissions.                                                                                                     string                  false      configureNamespace
namespace                    Setup Porter in this namespace                                                                                                                                                                                                                              string   <nil>          true       configureNamespace
porterConfig                 Porter config file, in yaml, same as ~/.porter/config.yaml                                                                                                                                                                                                  file                    false      configureNamespace
porterRepository             Docker image repository of the Porter agent. Defaults to ghcr.io/getporter/porter.                                                                                                                                                                          string                  false      configureNamespace
porterVersion                Version of the Porter agent, e.g. latest, canary, v0.33.0. Defaults to latest.                                                                                                                                                                              string                  false      configureNamespace
pullPolicy                   Specifies how the Porter agent image should be pulled. Does not affect how bundles are pulled. Defaults to PullAlways for latest and canary, and PullIfNotPresent otherwise.                                                                                string                  false      configureNamespace
serviceAccount               Name of the service account to run the Porter agent. If set, you are responsible for creating this service account and binding it to the porter-agent ClusterRole. Defaults to the porter-agent account created by the configureNamespace custom action.   string   porter-agent   false      configureNamespace
volumeSize                   Size of the volume shared between Porter and the bundles it executes. Defaults to 64Mi.                                                                                                                                                                     string                  false      configureNamespace

Actions:
Name                  Description                                                                                                                                        Modifies Installation   Stateless
configureNamespace   Add necessary rbac, service account and configuration to use Porter Operator in a namespace. Creates the namespace if it does not already exist.   false                   false
removeData           Remove Porter Operator data, such as namespaces used with configureNamespace, configuration, jobs, etc. These are not removed during uninstall.   false                   false

This bundle uses the following tools: exec, helm3, kubernetes.
```

Generate a credential set for the bundle, the only required credential for the operator is a kubeconfig for the cluster that the operator is to be installed in.
```
porter credentials generate porterops -r ghcr.io/getporter/porter-operator:canary
```

Install the operator into the porter-operator-system namespace:
```
porter install porterops -c porterops -r ghcr.io/getporter/porter-operator:canary
```

Create a namespace with the appropriate RBAC and configuration. This namespace is where you will create installation CRDs and the operator will create corresponding Jobs to execute the porter CLI.

```
porter invoke porterops --action configureNamespace --param namespace=TODO -c porterops
```

**Notes**
* The operator installs a mongodb server in its namespace (with no password set for root). This is only
  suitable for testing the operator.
* A PorterConfig resource named default is created in the specified namespace configuring Porter to use
  the kubernetes.secrets and mongodb plugin.

# Configuration

The bundle accepts a parameter, porterConfig, that should be a YAML-formatted [Porter configuration file](/configuration/).

Here is an example of the default configuration used when none is specified:

```yaml
# Resolve secrets using secrets on the cluster in the current namespace.
defaultSecretsPlugin: "kubernetes.secrets"

# Use the mongodb server that was deployed with the operator
defaultStorage: "in-cluster-mongodb"
storage:
  - name: "in-cluster-mongodb"
    plugin: "mongodb"
    config:
      url: "mongodb://mongodb.porter-operator-system.svc.cluster.local"
```

You can use a different file when installing the operator with the \--param flag:

```
porter install porterops --param porterConfig=./myconfig.yaml ...
```

The bundle also has parameters defined that control how the [Porter Agent] is configured and run.

| Parameter  | Description  |
|---|---|
| namespace  | Setup Porter in this namespace  |
| porterRepository  | Docker image repository of the Porter agent.<br/><br/>Defaults to ghcr.io/getporter/porter.  |
| porterVersion  | Version of the Porter agent, e.g. latest, canary, v0.33.0.<br/><br/>Defaults to latest.  |
| pullPolicy  | Specifies how the Porter agent image should be pulled. Does not affect how bundles are pulled.<br/><br/>Defaults to PullAlways for latest and canary, and PullIfNotPresent otherwise.  |
| serviceAccount  | Name of the service account to run the Porter agent.<br/>If set, you are responsible for creating this service account and binding it to the porter-agent ClusterRole.<br/><br/>Defaults to the porter-agent account created by the configureNamespace custom action.  |
| installationServiceAccount  | Name of the service account to run installation with.<br/>If set, you are responsible for creating this service account and giving it required permissions.  |
| volumeSize  | Size of the volume shared between Porter and the bundles it executes.<br/><br/>Defaults to 64Mi.  |


# Inspect the installation

You can use the porter CLI to query and interact with installations created by the operator.
Follow the instructions in [Connect to the in-cluster mongo database][connect] to point porter at the Mongodb server that was installed with the operator.

[install the Porter v1 prerelease]: /install/#v1-prerelease
[Porter Agent]: /operator/file-formats/#agent-config
