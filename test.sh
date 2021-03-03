#!/bin/sh
set -eu

go build

foobar="foo bar"

encoded=$(echo -n $foobar | ./urle | sed '/^$/d')

if [ "$encoded" != "foo+bar" ] ; then
  echo "encoding failed"
  echo "expected: (foo+bar)"
  echo "got: ($encoded)"
  exit 1
fi


decoded=$(echo $encoded | ./urle -d | sed '/^$/d')

if [ "$decoded" != "$foobar" ] ; then
  echo "decoding failed"
  echo "expected: ($foobar)"
  echo "got: ($decoded)"
  exit 1
fi

echo "passed"


