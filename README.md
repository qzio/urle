# urle

encode and decode urlencoded strings

```
# encode
$ echo -n 'foo bar' | ./urle

foo+bar

# decode
echo 'foo+bar' | ./urle -d

foo bar

```

