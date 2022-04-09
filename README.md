# dodo

This PoC is not quite same with normal shamir's secret sharing, it introduces using asymmetric cryptography to restore
the secret.

## TODO

- [x] command line tool for PoC
- [ ] client-server structure
- [ ] Provide this service as public API

## Basic Usage / Demo

`dodo` would storage its data under `~/.dodo` in json files, you could inspect them.

### create a playground

```text
mkdir try-dodo
cd try-dodo
```

### prepare key pair for several members

```text
ssh-keygen -m pem -N '' -C alice -f alice
ssh-keygen -m pem -N '' -C bob -f bob
ssh-keygen -m pem -N '' -C carol -f carol
ssh-keygen -m pem -N '' -C dave -f dave
ssh-keygen -m pem -N '' -C eve -f eve
```

### install `dodo`

```text
go install github.com/dodo-says/dodo/cmd/dodo@master
```

### Create a committee

Committee is a group of participants who share the same secret.

```text
dodo committee add --description "dodo's first committee" dodo
```

### Add the members into committee

```text
dodo committee-member add --committee-name dodo --public-key alice.pub alice
dodo committee-member add --committee-name dodo --public-key bob.pub bob
dodo committee-member add --committee-name dodo --public-key carol.pub carol
dodo committee-member add --committee-name dodo --public-key dave.pub dave
dodo committee-member add --committee-name dodo --public-key eve.pub eve
```

### ANYBODY (not only the committee members) could create a secret

Someone (like me), create a new record with message "STRRL is a lazy guy". And it requires at least `5` (by `--threshold`)approval from the committee to decrypt this message.

```text
dodo record add --committee-name dodo --message "STRRL is a lazy guy" --threshold 5
```

### ANYBODY (not only the committee members) could create a proposal to decrypt the secret

> replace the ids when you execute the following commands

```text
❯ dodo record list --committee-name dodo 
ID      Description     Committee       Threshold
22735228-f8fb-4691-b728-863bf5694210            dodo    5

❯ dodo decrypt-proposal create --record-id 22735228-f8fb-4691-b728-863bf5694210 --reason "I think this message is dangerous" 

❯ dodo decrypt-proposal list                                                      
ProposalID      RecordID        Reason
b4b5f61d-43b9-41d3-bfab-f53497724fef    22735228-f8fb-4691-b728-863bf5694210    I think this message is dangerous

❯ dodo decrypt-proposal inspect --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef                                          
Proposal ID: b4b5f61d-43b9-41d3-bfab-f53497724fef
Proposal Reason: I think this message is dangerous
Record ID: 22735228-f8fb-4691-b728-863bf5694210
Record Description: 
Committee: dodo
Approve Committee Members: 

❯ dodo record decrypt --record-id 22735228-f8fb-4691-b728-863bf5694210 
No proposal has enough approvals, you should concat with other committee members for more approvals
Available proposals:
Proposal ID: b4b5f61d-43b9-41d3-bfab-f53497724fef, Reason: I think this message is dangerous, threshould: 5, approved members: 

```

### Only committee members could approve the proposal

```text
dodo decrypt-proposal get-encrypted-slice --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef --member-name alice | age -d -i ./alice | dodo decrypt-proposal approve --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef
dodo decrypt-proposal get-encrypted-slice --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef --member-name bob | age -d -i ./bob | dodo decrypt-proposal approve --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef
dodo decrypt-proposal get-encrypted-slice --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef --member-name carol | age -d -i ./carol | dodo decrypt-proposal approve --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef
dodo decrypt-proposal get-encrypted-slice --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef --member-name dave | age -d -i ./dave | dodo decrypt-proposal approve --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef
dodo decrypt-proposal get-encrypted-slice --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef --member-name eve | age -d -i ./eve | dodo decrypt-proposal approve --proposal-id b4b5f61d-43b9-41d3-bfab-f53497724fef
```

### After arrive the threshold, the record can be decrypted

```text
❯ dodo record decrypt --record-id 22735228-f8fb-4691-b728-863bf5694210
Decrypted record: STRRL is a lazy guy
```

### purge the dodo's storage

```text
cd ../
rm -rf try-dodo
rm -rf ~/.dodo
```
