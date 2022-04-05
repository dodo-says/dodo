# dodo

This PoC is not quite same with normal shamir's secret sharing, it introduces using asymmetric cryptography to restore
the secret.

## TODO

- [ ] command line tool for PoC
- [ ] client-server structure
- [ ] Provide this service as public API

## Basic Usage

`dodo committee add dodo-says`

`dodo committee member add --committee dodo-says --public-key-file=./alice-public.pem alice`

`dodo committee member add --committee dodo-says --public-key-file=./bob-public.pem bob`

`dodo committee member add --committee dodo-says --public-key-file=./charlie-public.pem charlie`

`dodo record add --committee dodo-says "STRRL is the laziest dodo"`

`dodo committee proposal create --committee dodo-says --record-id=<uuid>`

`dodo committee proposal get --committee dodo-says --member-name=alice --record-id=<uuid> slice | <openssl decrypt>`

`dodo committee proposal approve --committee dodo-says --member-name=alice --record-id=<uuid> <content-of-decrypted-slice>`

> same as charlie

`dodo record read --committee dodo-says --record-id=<uuid>`