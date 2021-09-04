# kubectl-whoami

Show who am I via parse info from kubeconfig.

## install

Download from github release page.

## usage

summary:

```
$ kubectl-whoami 
User(subject.common_name): kubernetes-admin
Organization(subject.organization): system:masters
Groups(subject.names):
 - system:masters
 - kubernetes-admin
NotBefore(not_before): 2021-03-18 07:26:17 +0800 CST
NotAfter(not_after): 2022-03-18 07:26:24 +0800 CST
```

raw:

```
$ kubectl-whoami --raw
[
{
 "subject": {
  "common_name": "kubernetes-admin",
  "organization": "system:masters",
  "names": [
   "system:masters",
   "kubernetes-admin"
  ]
 },
 "issuer": {
  "common_name": "kubernetes",
  "names": [
   "kubernetes"
  ]
 },
 "serial_number": "4331093511755289885",
 "not_before": "2021-03-17T23:26:17Z",
 "not_after": "2022-03-17T23:26:24Z",
 "sigalg": "SHA256WithRSA",
 "authority_key_id": "",
 "subject_key_id": "",
 "pem": "-----BEGIN CERTIFICATE-----\nMIIC8jCCAdqgAwIBAgIIPBs...H8l3aWY=\n-----END CERTIFICATE-----\n"
}
]
```
