# QRLFT

QRLFT is a repository maintained by theQRL, containing tools for hashing and signing strings using Dilithium, a post-quantum cryptographic algorithm. Below is the detailed information and instructions on how to set up and use the tools provided in this repository.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Commands](#commands)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction
QRLFT offers a suite of cryptographic tools primarily focused on string hashing and signing using the Dilithium algorithm, which is designed to be resistant to quantum computer attacks.

## Features
- **String Hashing**: Hash strings using secure algorithms.
- **Dilithium Signing**: Sign strings with the Dilithium algorithm for post-quantum security.
- **Signature Verification**: Verify the validity of signatures.
- **Salt Generation**: Generate random salt of any length.
- **Public Key Retrieval**: Retrieve the public key from a given Dilithium hexseed.

## Installation
To install the QRLFT tools, ensure you have [Go](https://golang.org/dl/) installed on your machine. Then, clone the repository and build the project using the following commands:

```bash
git clone https://github.com/theQRL/qrlft.git
cd qrlft
go build
```


## Commands

```txt

NAME:
   qrlft - QRL File Tools - See docs at https://github.com/theQRL/qrlft

USAGE:
   qrlft [global options] command [command options] 

COMMANDS:
   verify     verify a dilithium signature matches the target file [eg. qrlft verify --signature=3b4e... doc.txt]
   sign       signs a file with a dilithium signature [eg. qrlft sign --hexseed=f29f58aff0b00de2844f7e20bd9eeaacc379150043beeb328335817512b29fbb7184da84a092f842b2a06d72a24a5d28 doc.txt]
   publickey  outputs the public key for a private hexseed to a file or to console [eg. qrlft publickey --hexseed=f29f58aff0b00de2844f7e20bd9eeaacc379150043beeb328335817512b29fbb7184da84a092f842b2a06d72a24a5d28 mykey.pub]
   hash       hashes a file with algorithm selected in options [eg. qrlft hash --sha256 doc.txt]
   salt       generates user-specified bytes random salt [eg. qrlft salt 16]
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```


## Usage

Once installed, you can use the following commands:


- [hash](#hash)
- [publickey](#publickey)
- [salt](#salt)
- [sign](#sign)
- [verify](#verify)



### Hash

```txt
NAME:
   qrlft hash - hashes a file with algorithm selected in options [eg. qrlft hash --sha256 doc.txt]

USAGE:
   qrlft hash [command options]

OPTIONS:
   --sha3-512    hash with SHA3-512 (default: false)
   --keccak-256  hash with Keccak-256 (default: false)
   --keccak-512  hash with Keccak-512 (default: false)
   --sha256      hash with SHA256 (default: false)
   --sha1        hash with SHA1 (default: false)
   --md5         hash with MD5 (default: false)
   --crc32       hash with CRC32 (default: false)
   --blake2s     hash with BLAKE2s (default: false)
   --quiet       just output the hash, no filename (default: false)
   --string, -s  hash a string instead of a file [eg. qrlft hash --sha256 HashThisText] (default: false)
   --help, -h    show help
   ```

- Hash a String:

```bash
./qrlft hash --sha256 --string "your string here"
```
- Hash a File:

```bash
./qrlft hash --sha256 doc.txt
```



### PublicKey

```txt
NAME:
   qrlft publickey - outputs the public key for a private hexseed to a file or to console [eg. qrlft publickey --hexseed=f29f58aff0b00de2844f7e20bd9eeaacc379150043beeb328335817512b29fbb7184da84a092f842b2a06d72a24a5d28 mykey.pub]

USAGE:
   qrlft publickey [command options]

OPTIONS:
   --hexseed SEED, --hs SEED  [Required] private key SEED
   --quiet                    just output the signature, no filename (default: false)
   --print, -p                prints the public key to the console instead of writing to a file [eg. qrlft publickey --print --hexseed=f29f58aff0b00de2844f7e20bd9eeaacc379150043beeb328335817512b29fbb7184da84a092f842b2a06d72a24a5d28] (default: false)
   --help, -h                 show help
   ```

- Retrieve Public Key

```bash
qrlft publickey --hexseed=f29f58aff0b00de2844f7e20bd9eeaacc379150043beeb328335817512b29fbb7184da84a092f842b2a06d72a24a5d28 mykey.pub
```


### Salt

```txt
NAME:
   qrlft salt - generates user-specified bytes random salt [eg. qrlft salt 16]

USAGE:
   qrlft salt [command options]

OPTIONS:
   --help, -h  show help
```

- 32 bytes salt

```bash
./qrlft salt 32
```


### Signing

```txt
NAME:
   qrlft sign - signs a file with a dilithium signature [eg. qrlft sign --hexseed=f29f58aff0b00de2844f7e20bd9eeaacc379150043beeb328335817512b29fbb7184da84a092f842b2a06d72a24a5d28 doc.txt]

USAGE:
   qrlft sign [command options]

OPTIONS:
   --hexseed SEED, --hs SEED  Signs file using the private key SEED
   --quiet                    just output the signature, no filename (default: false)
   --string, -s               hash a string instead of a file [eg. qrlft hash --sha256 HashThisText] (default: false)
   --help, -h                 show help
   ```

- Sign a String:

```bash
./qrlft sign --hexseed={SEED} --string "your string here"
```

- Sign a File:

```bash
./qrlft sign --hexseed={SEED} doc.txt
```

### Verify

```txt
NAME:
   qrlft verify - verify a dilithium signature matches the target file [eg. qrlft verify --signature=3b4e... doc.txt]

USAGE:
   qrlft verify [command options]

OPTIONS:
   --sigfile value, --sf value    Signature is a file [eg. qrlft verify --sigfile=signature.sig doc.txt]
   --signature value, -s value    Signature is included on the command line [eg. qrlft verify --signature=3b4e... doc.txt]
   --publickey value, --pk value  Specify the public key of the signer on command line [eg. qrlft verify --publickey=3b4e... doc.txt]
   --pkfile value, --pkf value    Specify the public key of the signer in a file [eg. qrlft verify --pkfile=publickey.pub doc.txt]
   --help, -h                     show help
   ```

- Verify a Signature:

```bash
./qrlft verify --sigfile sig.file --pkfile pk.file doc.txt
```

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Make your changes and commit them (git commit -m 'Add some feature').
4. Push to the branch (git push origin feature-branch).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

For more information and detailed usage instructions, please refer to the source code and comments within the repository.

Feel free to reach out if you have any questions or need further assistance. Happy coding!
