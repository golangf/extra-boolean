package boolean

// NOT, EQ, NEQ, IMPLY, NIMPLY (FIXED)

// Checks if value is false.
//
//   // Not(a, b)
//   // a: a boolean
//
// Example:
//   Not(false) == true
//   Not(true)  == false
func Not(a bool) bool {
	return !a
}

// Checks if antecedent ⇔ consequent (a ⇔ b).
//
//   // Eq(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Eq(true, true)   == true
//   Eq(false, false) == true
//   Eq(true, false)  == false
//   Eq(false, true)  == false
func Eq(a bool, b bool) bool {
	return a == b
}

// Checks if antecedent ⇎ consequent (a ⇎ b).
//
//   // Neq(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Neq(true, false)  == true
//   Neq(false, true)  == true
//   Neq(true, true)   == false
//   Neq(false, false) == false
func Neq(a bool, b bool) bool {
	return a != b
}

// Checks if antecedent ⇒ consequent (a ⇒ b).
//
//   // Imply(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Imply(true, true)   == true
//   Imply(false, true)  == true
//   Imply(false, false) == true
//   Imply(true, false)  == false
func Imply(a bool, b bool) bool {
	return !a || b
}

// Checks if antecedent ⇏ consequent (a ⇏ b).
//
//   // Nimply(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Nimply(true, false)  == true
//   Nimply(true, true)   == false
//   Nimply(false, true)  == false
//   Nimply(false, false) == false
func Nimply(a bool, b bool) bool {
	return !Imply(a, b)
}

// AND (VARIABLE)

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And(a bool, b bool) bool {
	return And2(a, b)
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And0() bool {
	return true
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And1(a bool) bool {
	return a
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And2(a bool, b bool) bool {
	return a && b
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And3(a bool, b bool, c bool) bool {
	return a && b && c
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And4(a bool, b bool, c bool, d bool) bool {
	return a && b && c && d
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And5(a bool, b bool, c bool, d bool, e bool) bool {
	return a && b && c && d && e
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return a && b && c && d && e && f
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return a && b && c && d && e && f && g
}

// Checks if all values are true.
//
//   // And[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   And(true, true)               == true
//   And(true, false)              == false
//   And4(true, true, true, true)  == true
//   And4(true, false, true, true) == false
func And8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return a && b && c && d && e && f && g && h
}

// OR (VARIABLE)

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or(a bool, b bool) bool {
	return Or2(a, b)
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or0() bool {
	return false
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or1(a bool) bool {
	return a
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or2(a bool, b bool) bool {
	return a || b
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or3(a bool, b bool, c bool) bool {
	return a || b || c
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or4(a bool, b bool, c bool, d bool) bool {
	return a || b || c || d
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or5(a bool, b bool, c bool, d bool, e bool) bool {
	return a || b || c || d || e
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return a || b || c || d || e || f
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return a || b || c || d || e || f || g
}

// Checks if any value is true.
//
//   // Or[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Or(true, false)                 == true
//   Or(false, false)                == false
//   Or4(false, true, false, true)   == true
//   Or4(false, false, false, false) == false
func Or8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return a || b || c || d || e || f || g || h
}

// XOR (VARIABLE)

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor(a bool, b bool) bool {
	return Xor2(a, b)
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor0() bool {
	return false
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor1(a bool) bool {
	return a
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor2(a bool, b bool) bool {
	return a != b
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor3(a bool, b bool, c bool) bool {
	return Xor2(Xor2(a, b), Xor1(c))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor4(a bool, b bool, c bool, d bool) bool {
	return Xor2(Xor2(a, b), Xor2(c, d))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor5(a bool, b bool, c bool, d bool, e bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor1(e))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor2(e, f))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor3(e, f, g))
}

// Checks if odd no. of values are true.
//
//   // Xor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xor(true, false)              == true
//   Xor(true, true)               == false
//   Xor4(true, true, true, false) == true
//   Xor4(true, true, true, true)  == false
func Xor8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return Xor2(Xor4(a, b, c, d), Xor4(e, f, g, h))
}

// COUNT (VARIABLE)

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count(a bool, b bool) int {
	return Count2(a, b)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count0() int {
	return 0
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count1(a bool) int {
	if a {
		return 1
	}
	return 0
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count2(a bool, b bool) int {
	return Count1(a) + Count1(b)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count3(a bool, b bool, c bool) int {
	return Count2(a, b) + Count1(c)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count4(a bool, b bool, c bool, d bool) int {
	return Count2(a, b) + Count2(c, d)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count5(a bool, b bool, c bool, d bool, e bool) int {
	return Count4(a, b, c, d) + Count1(e)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count6(a bool, b bool, c bool, d bool, e bool, f bool) int {
	return Count4(a, b, c, d) + Count2(e, f)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) int {
	return Count4(a, b, c, d) + Count3(e, f, g)
}

// Counts no. of true values.
//
//   // Count[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Count(true, true)                 == 2
//   Count(true, false)                == 1
//   Count4(true, true, true, false)   == 3
//   Count4(false, true, false, false) == 1
func Count8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) int {
	return Count4(a, b, c, d) + Count4(e, f, g, h)
}

// NAND (VARIABLE)

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand(a bool, b bool) bool {
	return Nand2(a, b)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand0() bool {
	return !And0()
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand1(a bool) bool {
	return !And1(a)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand2(a bool, b bool) bool {
	return !And2(a, b)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand3(a bool, b bool, c bool) bool {
	return !And3(a, b, c)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand4(a bool, b bool, c bool, d bool) bool {
	return !And4(a, b, c, d)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand5(a bool, b bool, c bool, d bool, e bool) bool {
	return !And5(a, b, c, d, e)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return !And6(a, b, c, d, e, f)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return !And7(a, b, c, d, e, f, g)
}

// Checks if any value is false.
//
//   // Nand[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nand(true, false)              == true
//   Nand(true, true)               == false
//   Nand4(true, true, false, true) == true
//   Nand4(true, true, true, true)  == false
func Nand8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return !And8(a, b, c, d, e, f, g, h)
}

// NOR (VARIABLE)

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor(a bool, b bool) bool {
	return Nor2(a, b)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor0() bool {
	return !Or0()
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor1(a bool) bool {
	return !Or1(a)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor2(a bool, b bool) bool {
	return !Or2(a, b)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor3(a bool, b bool, c bool) bool {
	return !Or3(a, b, c)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor4(a bool, b bool, c bool, d bool) bool {
	return !Or4(a, b, c, d)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor5(a bool, b bool, c bool, d bool, e bool) bool {
	return !Or5(a, b, c, d, e)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return !Or6(a, b, c, d, e, f)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return !Or7(a, b, c, d, e, f, g)
}

// Checks if all values are false.
//
//   // Nor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Nor(false, false)                == true
//   Nor(true, false)                 == false
//   Nor4(false, false, false, false) == true
//   Nor4(false, false, true, false)  == false
func Nor8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return !Or8(a, b, c, d, e, f, g, h)
}

// XNOR (VARIABLE)

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor(a bool, b bool) bool {
	return Xnor2(a, b)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor0() bool {
	return !Xor0()
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor1(a bool) bool {
	return !Xor1(a)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor2(a bool, b bool) bool {
	return !Xor2(a, b)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor3(a bool, b bool, c bool) bool {
	return !Xor3(a, b, c)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor4(a bool, b bool, c bool, d bool) bool {
	return !Xor4(a, b, c, d)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor5(a bool, b bool, c bool, d bool, e bool) bool {
	return !Xor5(a, b, c, d, e)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor6(a bool, b bool, c bool, d bool, e bool, f bool) bool {
	return !Xor6(a, b, c, d, e, f)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor7(a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	return !Xor7(a, b, c, d, e, f, g)
}

// Checks if even no. of values are true.
//
//   // Xnor[n](a, b, ...)
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Xnor(true, true)                == true
//   Xnor(false, true)               == false
//   Xnor4(true, true, false, false) == true
//   Xnor4(true, true, true, false)  == false
func Xnor8(a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	return !Xor8(a, b, c, d, e, f, g, h)
}

// SELECT (VARIABLE)

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select(i int, a bool, b bool) bool {
	return Select2(i, a, b)
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select0(i int) bool {
	return false
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select1(i int, a bool) bool {
	switch i {
	case 0:
		return a
	default:
		return false
	}
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select2(i int, a bool, b bool) bool {
	switch i {
	case 0:
		return a
	case 1:
		return b
	default:
		return false
	}
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select3(i int, a bool, b bool, c bool) bool {
	switch i {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	default:
		return false
	}
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select4(i int, a bool, b bool, c bool, d bool) bool {
	switch i {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	case 3:
		return d
	default:
		return false
	}
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select5(i int, a bool, b bool, c bool, d bool, e bool) bool {
	switch i {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	case 3:
		return d
	case 4:
		return e
	default:
		return false
	}
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select6(i int, a bool, b bool, c bool, d bool, e bool, f bool) bool {
	switch i {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	case 3:
		return d
	case 4:
		return e
	case 5:
		return f
	default:
		return false
	}
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select7(i int, a bool, b bool, c bool, d bool, e bool, f bool, g bool) bool {
	switch i {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	case 3:
		return d
	case 4:
		return e
	case 5:
		return f
	case 6:
		return g
	default:
		return false
	}
}

// Checks if ith value is true.
//
//   // Select[n](i, a, b, ...)
//   // i: index
//   // a: 1st boolean
//   // b: 2nd boolean
//
// Example:
//   Select(0, true, false)               == true
//   Select(1, true, false)               == false
//   Select4(1, true, true, false, false) == true
//   Select4(2, true, true, false, false) == false
func Select8(i int, a bool, b bool, c bool, d bool, e bool, f bool, g bool, h bool) bool {
	switch i {
	case 0:
		return a
	case 1:
		return b
	case 2:
		return c
	case 3:
		return d
	case 4:
		return e
	case 5:
		return f
	case 6:
		return g
	case 7:
		return h
	default:
		return false
	}
}

// EQV, IMP (SHORTCUTS)

// Checks if antecedent ⇔ consequent (a ⇔ b).
//
//   // Eqv(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Eqv(true, true)   == true
//   Eqv(false, false) == true
//   Eqv(true, false)  == false
//   Eqv(false, true)  == false
func Eqv(a bool, b bool) bool {
	return Eq(a, b)
}

// Checks if antecedent ⇒ consequent (a ⇒ b).
//
//   // Imp(a, b)
//   // a: antecedent
//   // b: consequent
//
// Example:
//   Imp(true, true)   == true
//   Imp(false, true)  == true
//   Imp(false, false) == true
//   Imp(true, false)  == false
func Imp(a bool, b bool) bool {
	return Imply(a, b)
}