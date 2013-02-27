Plugin Functions
================

Implemented Functions
---------------------

### Position

### Map

Standard mapping.

tags.protocol=map("proto")

### Map Until

### Regular Expressions

Searches for a field in argument one of the plugin function and checks the value of that field for
the supplied regex. The tag is filled with the match of that regex.
In case of no match a empty string is returned.
```
tags.destination_address=regex("dst", "\d+\.\d+\.\d+\.\d+")
```

### Split

The split plugin function will process the value of a key (Argument 1) and split that value into fields.
Argument 2 is the character to split at.
Argument 3 is the number of the field to use as a tag value.
```
Example: dst=a:b:c
split("dst", ":", 0) will return 'a'
split("dst", ":", 1) will return 'b'

tags.destination_address=split("dst", ":", 0)
```
### D64

The d64 plugin function will decode a standard mapping which is in d64 format.
```
tags.customfield_1 = d64("userdata1")
```
### LIT

The lit (short for literal) does no processing, it simply puts the value of the tag as the same as the value
```
tags.customheader_1=lit("Hello People")
```
This will set the value for tags.customheader_0 as "Hello People"
