## Passgen

Passgen is a command-line tool that generates secure and memorable passphrases and cryptographically secure random passwords. It uses the [diceware](https://theworld.com/~reinhold/diceware.html) method and the Go [crypto/rand](https://pkg.go.dev/crypto/rand) package to create strong and reliable passwords.

## Usage

### Installation

To install passgen you will need to install Go and then run:

```
go install github.com/lucasherediadv/passgen@latest
```

### Generating passphrases and passwords

If you need to generate a passphrase:

```
passgen phrase
crop-gab-skimmed-petty-basket-lyrics-blooper-neuter
```

If you need to generate a password:

```
passgen word
@sFhjQ7p7CmbGTOMw!o^
```

If you need help with a specific command:

```
passgen [command] --help
```

For more information about the use and further customization of the generated passphrases and passwords, check out the full documentation [here](https://github.com/lucasherediadv/passgen/blob/main/doc/passgen.md).
