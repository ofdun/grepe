## About <a name = "about"></a>
Same as grep gnu utility, but extended (TODO)

### Installing

1. Clone repo using git:
```console
git clone https://github.com/ofdun/grepe.git
```
2. Run makefile using make in main directory
```console
make install
```

## Usage <a name="usage"></a>

After installing using make you should be able to run it using grepe command.

```console
grepe pattern [input file]
```

Here is some examples:

```console
>>> echo "apple\naboba\nape" | grepe ap
apple
ape
```

```console
>>> cat example.txt
apple
aboba
ape

>>> grepe "^a[a-z]+a$" example.txt
aboba
```

Implemented flags and their usage can be got from grepe --help.
```console
>>> grepe --help
grepe utility searches any given input files, selecting lines that match one or more patterns.

Usage:
  grepe [flags]

Flags:
      --color string     Color in which matches are highlighted. According to ECMA-48 standard the next values are expected: black, red, green, yellow, blue, magenta, cyan, white, black-background, red-background, green-background, yellow-background, blue-background, magenta-background, cyan-background, white-background. The default color is green (default "green")
      --colour string    Alias for --color.
  -c, --count            Suppress normal output; instead print a count of matching lines.
  -h, --help             help for grepe
  -i, --ignore-case      Ignore case distinctions in patterns and input data.
  -v, --invert-match     Invert the sense of matching, to select non-matching lines. (-v is specified by POSIX)
  -x, --line-regexp      Select only those matches that exactly match the whole line. (-x is specified by POSIX.)
  -m, --max-count int    Stop after the first num selected lines. If num is zero, grepe stops right away without reading input. A negative num is treated as infinity and grepe does not stop; this is the default. (default -1)
      --no-ignore-case   Do not ignore case distinctions in patterns and input data. This is the default. This option is useful for passing to shell scripts that already use -i, in order to cancel its effects.
  -o, --only-matching    Print only the matched non-empty parts of matching lines, with each such part on a separate output line.
  -q, --quiet            Quiet; do not write anything to standard output. Exit with zero status if any match is found
      --silent           Quiet; do not write anything to standard output. Exit with zero status if any match is found
      --version          version for grepe
  -w, --word-regexp      Select only those lines containing matches that form whole words. This option has no effect if -x is also specified.
```

