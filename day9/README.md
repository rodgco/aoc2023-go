# Day 9: Mirage Maintenance

## Part 1

First I tought that it would be simple as finding the difference between any two numbers in the input and add the difference to last value. I was wrong, really needed to drill down the differences to extrapolate the additional number.

A recursive function to the rescue: every layer would be processed by the same function until finding the last line (all zeros). Struggled a little bit with the return of this last layer, and remembered that on a recursive function it's almost always the best choice to break the recursion at the beggining. And voi lá, it worked!

## Part 2

Easy peasy! Just used the option to return multiple values on the extrapolate function. Just added the extra return value for the before element. Got it right at my first attempt. Hole-in-one!

## What I learned about Go?

Returning multiple values on a func can be useful!

## What I learned about myself?

I still can recurse, and I'm proud to solve this with 58 lines of code.

## What else can I do?

I'm pretty sure there are ways to improve the code to turn the input into an array of arrays of ints.
