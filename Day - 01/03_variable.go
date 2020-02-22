/*
The types in Go can be classified as follows −
1. Boolean types - They are boolean types and consists of the two predefined constants: (a) true (b) false

2. Numeric types - They are again arithmetic types and they represents a) integer types or b) floating point values throughout the program.

# Integer Types
i. uint8 - Unsigned 8-bit integers (0 to 255)
ii. uint16 - Unsigned 16-bit integers (0 to 65535)
iii. uint32 - Unsigned 32-bit integers (0 to 4294967295)
iv. uint64 - Unsigned 64-bit integers (0 to 18446744073709551615)
v. int8 - Signed 8-bit integers (-128 to 127)
vi. int16 - Signed 16-bit integers (-32768 to 32767)
vii. int32 - Signed 32-bit integers (-2147483648 to 2147483647)
viii. int64 - Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

# Floating Types
i. float32 - IEEE-754 32-bit floating-point numbers
ii. float64 - IEEE-754 64-bit floating-point numbers
iii. complex64 - Complex numbers with float32 real and imaginary parts
iv. complex128 - Complex numbers with float64 real and imaginary parts

# Other Numeric Types
i. byte - same as uint8	
ii. rune - same as int32
iii. uint - 32 or 64 bits
iv. int - same size as uint
v. uintptr - an unsigned integer to store the uninterpreted bits of a pointer value

3. String types - A string type represents the set of string values. Its value is a sequence of bytes. Strings are immutable types that is once created, it is not possible to change the contents of a string. The predeclared string type is string.

4. Derived types - They include (a) Pointer types, (b) Array types, (c) Structure types, (d) Union types and (e) Function types f) Slice types g) Interface types h) Map types i) Channel Types

*/

package main

import (
	"fmt" /* Package fmt implements formatted I/O with functions analogous to C's printf and scanf.  */
)

func main() {

    var num1 int = 10
    var num2 int = 20
    fmt.Println("Num1 + Num2 =",num1+num2)

    var num3 float32 = 10
    var num4 float32 = 20 
    fmt.Println("Num3 - Num4 =",num3-num4)
}