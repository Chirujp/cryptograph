# cryptograh

Cryptograph is a KMS for chiru.

its role is to:

- generate Certificates and private keys for other apis
- deliver the certificates and private keys to the other apis
- generate keys for the other apis
- deliver the keys to the other apis
- generate rsas for the other apis
- deliver the rsas to the other apis

Why is it needed?

For communcate between internal apis.
We using self-signed certificates because we dont have a trusted certificate authority.

Why is it important?

It's important for encrypting the data between the apis.

## Installation

```bash
git clone https://github.com/chirujp/cryptograh.git
cd cryptograh
go get -v
go build
./cryptograh
```

