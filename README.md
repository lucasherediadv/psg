## Passgen

Passgen is a passphrase and password generator that uses the diceware method to generate secure and memorable passphrases, and use the Go crypto package to generate cryptographically secure random passwords.

If you need to generate a passphrase, run `passgen phrase`:

```sh
$ passgen phrase
crop-gab-skimmed-petty-basket-lyrics-blooper-neuter
```

If you need to generate a password, run `passgen word`:

```sh
$ passgen word
@sFhjQ7p7CmbGTOMw!o^
```

## Usage

### Installing

To install you will need to install Go and then run:

```
go install github.com/lucasherediadv/passgen@latest
```

### Generating passphrases and passwords

