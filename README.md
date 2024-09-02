# Xk

The X replacement on Kubernetes

The idea is to join into a common room for all operators and exchange
messages through out a self-hosted centralized server (i.e. InspIRCd)

In the long term the feed can have events, logs and Kubernetes actions
posted automatically in the channel by the server/controller.

### Pre-requisites

* Run a IRC Server on K8s

### How to use the plugin

todo: Use port-forward locally to reach the server by service label. 
todo: Use `rivo/tview` for UI 

`kubectl x --nick <user CN> --server localhost:6667`

