## About <a name = "about"></a>
Same as grep gnu utility, but extended (TODO)

### Installing

1. Clone repo using git:
```console
git clone https://github.com/ofdun/grepe.git
```
2. Run makefile using make in main directory
```console
make
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

>>> grepe ap example.txt
apple
ape
```

