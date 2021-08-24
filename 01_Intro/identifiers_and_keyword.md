# Identifiers in Golang
- An identifier is a token that must be composed of the Unicode letters, Unicode digits and _ (underscore), and start with either a Unicode letter or _.
- Rules for Defining Identifiers
    1. The first position of an identifier must be a letter or an underscore. And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as a character ‘_’.
    2. Idiomatic Go favors mixed caps (camel case) naming.
    3. Package-level identifiers must be unique across the given package.
    4. Identifiers must be unique within a code block (functions, control statements).
    5. The name of an identifier is case sensitive.
    6. Keywords are not allowed to use as an identifier name.
    7. identifiers name should be meaningful and advised to be of 4-15 characters.

- In Golang, there are some predeclared identifiers available for constants, types, and functions. These names are not reserved; you are allowed to use them in your declaration.  
`For Functions:`
make, len, cap, new, append, copy, close, 
delete, complex, real, imag, panic, recover  
`For Constants:`
true, false, iota, nil  
`For Types:`
int, int8, int16, int32, int64, uint,
uint8, uint16, uint32, uint64, uintptr,
float32, float64, complex128, complex64,
bool, byte, rune, string, error  

-  `Blank Identifier[_]`: It is used as the anonymous placeholder instead of the regular identifier, and it has a unique meaning in declarations, as an operand, and in assignments.

- `Exported Identifier`: The identifier which is allowed to access it from another package. The first character should be in the Unicode upper case letter and should be unique from the other set of identifiers available in your program or the package.

# Keywords in Golang

-  Unique words that help compilers understand and parse user code.
-  Go has only 25 keywords and can be categorized into four groups:
    1. `Go program`: const, func, import, package, type, var
    2. `Composite type denotations`: chan, interface, map, struct
    3. `Control Flow`: break, case, continue, default, else, fallthrough, for, goto, if, range, return, select, switch
    4. `Misc`[Control Flow]: defer, go