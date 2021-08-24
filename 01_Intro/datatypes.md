# Datatypes in GO  
    1. Boolean Type 	: true, false
	2. Numeric Type 	: integer, float
	3. String Type 		: String
	4. Aggregate Type 	: Arrays and Struct
	5. Reference Type	: Pointers, Slices, Functions, Maps and Channels
    6. Interface Type   : Interface

## Integers


    1.	uint8   :   Unsigned 8-bit integers (0 to 255)
    2.	uint16  :   Unsigned 16-bit integers (0 to 65535)
    3.	uint32  :   Unsigned 32-bit integers (0 to 4294967295)
    4.	uint64  :   Unsigned 64-bit integers (0 to 18446744073709551615)
    5.	int8    :   Signed 8-bit integers (-128 to 127)
    6.	int16   :   Signed 16-bit integers (-32768 to 32767)
    7.	int32   :   Signed 32-bit integers (-2147483648 to 2147483647)
    8.	int64   :   Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

## Floating Types

    1.	float32     :   IEEE-754 32-bit floating-point numbers
    2.  float64     :   IEEE-754 64-bit floating-point numbers
    3.  complex64   :   Complex numbers with float32 real and imaginary parts
    4.  complex128  :   Complex numbers with float64 real and imaginary parts

    The value of an n-bit integer is n bits and is represented using two's complement arithmetic operations.

##  Other Numeric Types 
    1.  byte    :   same as uint8
    2.  rune    :   same as int32
    3.  uint    :   32 or 64 bits
    4.  int     :   same size as uint
    5.  uintptr :   an unsigned integer to store the uninterpreted bits of a pointer value