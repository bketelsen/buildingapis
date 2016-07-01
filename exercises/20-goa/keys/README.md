This directory contains a RSA key pair encoded in PEM format used to
encrypt and sign JWT tokens. Note that the private key is not protected
by a password to keep things simple.

These were generated with:

```
openssl genrsa -out jwt.key 2048
openssl rsa -in jwt.key -outform PEM -pubout -out jwt.key.pub
```
