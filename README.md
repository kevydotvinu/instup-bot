### Instup Bot
A bot for custom task and notification.

### Commands
#### Manifest Graph
For the below command, it responses with a link that shows the OpenShift release image [manifest graph][1]. The graph is ordered by the number and component of the manifest file.
```bash
!manifestgraph 4.10.10
```
### Dependency
#### Curl-paste
The [curl-paste][2] should be running independently for the Instup Bot to work. The HTTP POST request will send the form data along and respond with a link to the paste. The HTTP GET will retrieve the paste with the given ID as plain-text.

[1]: https://github.com/openshift/enhancements/blob/master/dev-guide/cluster-version-operator/user/reconciliation.md#manifest-graph
[2]: https://github.com/kevydotvinu/curl-paste
