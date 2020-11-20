#! /bin/bash
INPUT="$1"
OUTPUT="$2"
INTERFACE_NAME="$3"

echo "pragma solidity >=0.4.24 <=0.8.0;" > "$OUTPUT"
abi2solidity -i "$INPUT" -o "$OUTPUT.2"
cat "$OUTPUT.2" >> "$OUTPUT"
rm "$OUTPUT.2"
sed -i "s/GeneratedInterface/$INTERFACE_NAME/g" "$OUTPUT"