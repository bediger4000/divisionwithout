# Divide without division

The basic problem involves doing integer arithmetical division
without using built-in division operator(s).
I've seen it show up as the Daily Coding Problem in 2 similar,
but not identical ways.

---

### Daily Coding Problem: Problem #380 [Medium] 

This problem was asked by Nextdoor.

Implement integer division without using the division operator.
Your function should return a tuple of
(dividend,remainder)
and it should take two numbers,
the product and divisor.

For example,
calling divide(10, 3) should return (3, 1)
since the divisor is 3 and the remainder is 1.

Bonus: Can you do it in O(log n) time?

---

### Daily Coding Problem: Problem #896 [Medium] 

This question was asked by ContextLogic.

Implement division of two positive integers without using the division,
multiplication,
or modulus operators.
Return the quotient as an integer,
ignoring the remainder.

---

## Algorithm Analysis

Neither of the algorithms deny the use of the subtraction operator,
or any bitwise operators.
It's certainly possible to divide by 2 without subtraction,
but I wonder if division by any other divisor is possible
without subtraction?

The "[Medium]" rating would seem to direct away from
counting subtractions of the divisor until the remainder
is numerically less than the divisor.
Nevertheless, I [implemented](div1.go) that algorithm.
It works great, and it's in line with how one might implement
division with [Recursive Functions](https://plato.stanford.edu/entries/recursive-functions/)
It doesn't work in O(log n) time.

Speaking of O(log n) time, what's n?
Size of divisor in binary, or something else?
That would seem to necessitate a question by the job candidate.
