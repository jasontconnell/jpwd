# JPWD

A small and simple tool to generate a password instead of googling "password generator" and picking a random one

# USAGE

## Install

With Go [installed](https://go.dev), run 

```
go install github.com/jasontconnell/jpwd/cmd/jpwd
```

## Run

```
jpwd -length 20 -uppercase -lowercase -symbols="!@#$%^&*()_+-=;:,./<>?"
```

## Options

- length: the length of the password to generate
- uppercase: use uppercase characters in generation
- lowercase: use lowercase characters in generation
- numbers: use numbers in generation
- symbols: provide a list of symbols that can be used

# Alternative Uses

Instead of using all of the sets of uppercase, lowercase, numbers, symbols etc, you can provide `false` values for each of the boolean flags and provide your own set of symbols (false is default for boolean flags)

```
>>> jpwd -length 20 -symbols="abcABC1234!@#$"
>>> $A#aa#4!B#a@#$@c
```

I will let you determine the usefulness of that.
