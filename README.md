## Psg

Psg is a command-line tool that generates memorable passphrases and cryptographically secure passwords. It uses the [diceware](https://theworld.com/~reinhold/diceware.html) method and the Go [crypto/rand](https://pkg.go.dev/crypto/rand) package to create strong passphrases and passwords.

### Installation

To install psg you will need to install Go and then run:

```
go install github.com/lucasherediadv/psg@latest
```

### Generating passphrases and passwords

If you need to generate a passphrase:

```
psg phrase
cesarean-marrow-baking-triumph-urgency-scrambler-maximize-bloomers
```

If you need to generate a password:

```
psg word
ekwenibueadajcfkkwtt
```

If you need help with a specific command:

```
psg [command] --help
```
