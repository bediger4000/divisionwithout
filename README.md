# Divide without division

The basic problem involves doing integer arithmetical division
without using built-in division operator(s).
I've seen it show up as the Daily Coding Problem in two similar,
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
Size of divisor in binary,
size of numerator in decimal digits
or something else?
The job candidate needs to ask about that.

I could not find an O(log n) solution by myself.
I had to google for [an answer](https://plato.stanford.edu/entries/recursive-functions/).
I [reimplemented in Go](div2.go).
I would have failed the "bonus" part of one of the questions.

The subtract-and-shift method makes as many trips through
the subtract-and-shift loop as the position of the highest
1-bit in the denominator.
Since base 2 is a positional representation,
that's roughly log-base-2 of the denominator,
so this method is O(log n) in the magnitude of the denominator.
There's also a varying number of bitwise left shifts
depending on the relative sizes of the numerator and denominator.
There doesn't appear to be some obscure high-constant-factor
algorithm for dividing extremely large numbers by small numbers,
like there is for multiplying two very large numbers.

## Interview Analysis

This is an odd question.
There's an obvious, very easy solution,
and an inobvious, tricky-to-program solution.
The rating hints at the inobvious solution,
as does the "bonus" given by one of the problem statements.

It seems to me that an interviewer could expect just about any candidate
to come up with the repeated subtraction algorithm.
The shifting back and forth of the inobvious algorithm
does not lend itself to getting a correct whiteboard answer.
Fiddling and testing are probably required to get it correct.

Neither algorithm does a lot to check a candidate's
comp sci prowess.
No data structure necessary, not even an array, no pointers.
There's really only 1 loop.

As far as test cases go,
divisors of 1, 2 and several primes seem necessary.
Tests where numerator and denominator are multiples
are important,
particularly multiples of 2 since the inobvious algorithm
does a lot of bit-shifting.
Probably also need to include division-by-zero,
and negative numerator and denominator cases.

I feel like this isn't a great question.
Either a candidate knows the inobvious algorithm
and repeats it from memory, or the candidates goes down in flames.
Not a lot of programming can be elicited from a candidate.
