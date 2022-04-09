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

```bash
for name in 'alice' 'bob' 'carol' 'dave' 'eve'; do
    ssh-keygen -m pem -N '' -C $name -f $name
done
```

### install `dodo`

```text
go install github.com/dodo-says/dodo/cmd/dodo@master
```

Thanks to [cobra](https://github.com/spf13/cobra), `dodo` provides completion for `bash`, `zsh`, `fish` and `powershell`, using completion would improve the experience, see:

- `dodo completion bash --help`
- `dodo completion zsh --help`
- etc...

### Create a committee

Committee is a group of participants who share the same secret.

```text
dodo committee add --description "dodo's first committee" dodo
```

### Add the members into committee

```bash
for name in 'alice' 'bob' 'carol' 'dave' 'eve'; do
    dodo committee-member add --committee-name dodo --public-key $name.pub $name
done
```

### ANYBODY (not only the committee members) could create a secret

Someone (like me), create a new record with message "STRRL is a lazy guy". And it requires at least `4` (by `--threshold`)approval from the committee to decrypt this message.

```text
dodo record add --committee-name dodo --message "STRRL is a lazy guy" --threshold 4
```

### ANYBODY (not only the committee members) could create a proposal to decrypt the secret

> replace the ids when you execute the following commands

```text
❯ dodo record list --committee-name dodo 
ID      Description     Committee       Threshold
af84039a-6d08-49d7-8428-076b91639082            dodo    4


❯ dodo decrypt-proposal create --record-id 22735228-f8fb-4691-b728-863bf5694210 --reason "I think this message is dangerous" 

❯ dodo decrypt-proposal list                                                                                                
ProposalID      RecordID        Reason
49eef4c6-c62e-43c2-9747-f3eb79e98c5f    af84039a-6d08-49d7-8428-076b91639082    I think this message is dangerous

❯ dodo decrypt-proposal inspect --proposal-id 49eef4c6-c62e-43c2-9747-f3eb79e98c5f 
Proposal ID: 49eef4c6-c62e-43c2-9747-f3eb79e98c5f
Proposal Reason: I think this message is dangerous
Record ID: af84039a-6d08-49d7-8428-076b91639082
Record Description: 
Committee: dodo
Approve Committee Members: 

❯ dodo record decrypt --record-id af84039a-6d08-49d7-8428-076b91639082 
No proposal has enough approvals, you should concat with other committee members for more approvals
Available proposals:
Proposal ID: 49eef4c6-c62e-43c2-9747-f3eb79e98c5f, Reason: I think this message is dangerous, threshould: 4, approved members: 
```

### Only committee members could approve the proposal

```bash
# notice that alice doesn't approve the proposal
export PROPOSAL_ID=49eef4c6-c62e-43c2-9747-f3eb79e98c5f
for name in 'bob' 'carol' 'dave' 'eve'; do
    dodo decrypt-proposal get-encrypted-slice --proposal-id $PROPOSAL_ID --member-name $name | age -d -i ./$name | dodo decrypt-proposal approve --proposal-id $PROPOSAL_ID
done
```

### After arrive the threshold, the record can be decrypted

```text
❯ dodo decrypt-proposal inspect --proposal-id 49eef4c6-c62e-43c2-9747-f3eb79e98c5f 
Proposal ID: 49eef4c6-c62e-43c2-9747-f3eb79e98c5f
Proposal Reason: I think this message is dangerous
Record ID: af84039a-6d08-49d7-8428-076b91639082
Record Description: 
Committee: dodo
Approve Committee Members: bob, carol, dave, eve

❯ dodo record decrypt --record-id af84039a-6d08-49d7-8428-076b91639082 
Decrypted record: STRRL is a lazy guy
```

### purge the playground and dodo's storage

```text
cd ../
rm -rf try-dodo
rm -rf ~/.dodo
```
