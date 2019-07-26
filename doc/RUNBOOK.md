# Sine Service Runbook

This document describes how to monitor and fix issues with the sine-service.

When a run finishes without resolution, escalate to Incident Management- regardless of whether you can solve the issue or not.

## Locations

 * [Pipeline](https://circleci.com/gh/jspc/sine-service)
 * [Environment Repo](https://github.com/jspc/ori-env)
 * [Environment Pipeline](https://circleci.com/gh/jspc/ori-env)
 * [Chronograf](https://monitoring.ori.jspc.pw)  - Note: you need to be a member of `j-and-j-global` or `ori-edge` to access this

## Resources

 * [cluster](https://cloud.digitalocean.com/kubernetes/clusters/58e79fde-053c-4067-a644-ce4c09581696?i=401acb)
 * [config bucket](https://cloud.digitalocean.com/spaces/config-ori-jspc-pw?i=401acb)
 * [helm charts](https://cloud.digitalocean.com/spaces/charts-jspc-pw?i=401acb)

## Commands:

 * List deployments: `helm --tiller-namespace ori list`
 * Delete deployment: `helm --tiller-namespace ori delete --purge sine-service`
 * List pods: `kubectl --namespace ori get po -l app=sine-service`
 * List ingress: `kubectl --namespace ori get ing -l app=sine-service`
 * View logs: `kubectl --namespace ori logs -l app=sine-service`
 * Find redis service address: `kubectl --namespace ori get svc redis-redis-ha`

## Runs

### Sine Service refuses to start

1. When `sine-service` logs show: `panic: dial tcp [::1]:6379: connect: connection refused`
 1. This address is the default redis host
 1. Check whether the helm chart contains a value against `redis.address`, correct and deploy if not
 1. Goto next run
1. When `sine-service` cannot connect to its configured redis
 1. Check helm deployments for redis, and re-run the environment pipeline if it does not exist
 1. If it exists, but is in a failed state, delete it and re-run the pipeline
 1. If it is in a successful state, check the redis service address matches what the chart is set to.
 1. If you get this far without resolution, escalate
