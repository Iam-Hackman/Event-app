package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a_byte rune = '2'
	fmt.Println(a_byte)

	var a_float float32 = 0.11
	fmt.Println(a_float)

	var complex_num complex64 = complex(a_float, a_float)
	fmt.Println(complex_num)

	var real_comp float32 = real(complex_num)
	fmt.Println(real_comp)

	var imag_comp float32 = imag(complex_num)
	fmt.Println(imag_comp)

	var my_name string
	my_name = "Agyemang Hackman"
	fmt.Println(my_name)

	my_name = "Ntim Agyemang"
	fmt.Println(my_name)

	fmt.Printf("%s %s\n", my_name, my_name)

	var test_variable int = 20
	fmt.Println(test_variable)

	var test_var2 = "Hackman"
	fmt.Println(test_var2)

	test_var3 := 18
	fmt.Println(test_var3)

	fmt.Printf(" %s", test_var2)

	test_var4, test_var5 := "First", "name"
	fmt.Println(test_var4, test_var5)

	const myfaith int = 1
	const mysuccess = myfaith
	fmt.Println(mysuccess)

	// - minus arithmetic operator
	// + plus arithmetic operator
	// * multiplication arithmetic operator
	// / division arithmetic operator
	// & And logical operator
	// &^ And not logical operator
	// | Or logical operator
	// == equal logical operator
	// != Not equal logical operator
	// > Greater than logical operator
	// >= Greater than or equal to logical operator
	// < Less than logical operator
	// <= Less than or equal to logical operator
	// % quotient operator
	// x++ is a unary operator for increment
	// x-- is a unary operator for decrement
	// &x is a unaary operator to check the storage location for a value
	// && conditional And operator which returns a boolean
	// || conditional Or operator which returns a boolean
	// *x unary porrator for checking the value at storage location p
	// <-x receives from channel x
	// ^x unary operator called INVERT that inverts bit values.
	// ^ XOR operator sets bit to 1 if both comparing bit values are 1(set)

	/* ++ --
	* / % << >> & &^
	+ - | ^

	*/

	a := 10
	b := 5

	a += 6
	fmt.Println(a)
	a -= 6
	fmt.Println(a)
	a *= 2
	fmt.Println(a)
	a /= 2
	fmt.Println(a)
	a %= 2
	fmt.Println(a)
	a += 10
	fmt.Println(a)

	sum := a + b
	fmt.Println(sum)

	firstname := "Hackman"
	lastname := " Agyemang"
	fullname := firstname + lastname
	fmt.Println(fullname)

	var x uint8 = 207
	var y uint8 = 150

	c := x | y
	fmt.Println(strconv.FormatUint(uint64(x), 2))
	fmt.Println(strconv.FormatUint(uint64(y), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

	num := 19
	bitnum := strconv.FormatUint(uint64(num), 2)
	fmt.Println(20 << 2)
	fmt.Println(strconv.FormatUint(uint64(num), 2))
	fmt.Println(bitnum)

	//a | (1<<i) ith bit of a to 1
	// a &^ (1<<i) with set the ith bit of a to 0 (bit clear)
	// a ^ (1<<i) will invert the ith bit of a.
}
