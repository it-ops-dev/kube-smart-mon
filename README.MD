# Description

This is a modified S.M.A.R.T disk monitoring for Prometheus clusters. Feel free to poke around and create an issue for improvements, or changes.

# Requirements

* Kubernetes cluster
* Prometheus configured to scrape annotations ```prometheus.io/{path,scrape,port}```
* Grafana for visualization

# Installation

It is as simple as: 
```
kubectl apply -f https://raw.githubusercontent.com/it-ops-dev/kube-smart-mon/main/daemonset.yaml
```
And you can download the grafana dashboard from here: https://github.com/it-ops-dev/kube-smart-mon/tree/main/grafana

# Credits

Credits for micha37-martins writing a very nice SMART parser! Thank you!
Inspired by https://github.com/micha37-martins/S.M.A.R.T-disk-monitoring-for-Prometheus

