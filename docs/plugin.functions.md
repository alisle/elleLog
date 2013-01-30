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

tags.destination_address=regex("dst", "\d+\.\d+\.\d+\.\d+")

### Split

The split plugin function will process the value of a key (Argument 1) and split that value into fields.
Argument 2 is the character to split at.
Argument 3 is the number of the field to use as a tag value.

Example: dst=a:b:c
split("dst", ":", 0) will return 'a'
split("dst", ":", 1) will return 'b'

tags.destination_address=split("dst", ":", 0)
