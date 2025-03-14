Create a user: 

`oc create user <username>`

Create a group: 

`oc adm groups new <group-name>`

Create a group and add a user there: 

`oc adm groups new <group-name> <username>`

Remove user from a group: 

`oc adm groups remove-users <group-name> <username> `

Synchronize Group with the LDAP IdP provider info: 

`oc adm groups sync --sync-config=/path/to/ldap-sync-config.yaml --confirm`

Create an SA: 

`oc create serviceaccount <sa-name>  [-n <namespace>]`

Create a Role: 

`oc create role <role-name> --verb=<list> --resource=<list>  [-n <namespace>]`

Create a ClusterRole: 

`oc create clusterrole <role-name> --verb=<list> --resource=<list>`

Create RoleBinding for User: 

`oc create rolebinding admin --clusterrole= <role-name> --user=<user-name> --group=<group-name>`

Create RoleBinding for Service Account: 

`oc create rolebinding admin-binding --role=admin --serviceaccount=<namespace>:<sa-name>`

Get Rolebingings and Users, Groups, and Service Accounts for them: 

`oc get rolebindings -owide  [-n <namespace>]`

Check who can do specific actions, for example: 

`oc adm policy who-can <verb> <resource> [-n <namespace>]`
