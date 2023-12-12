# Day 8: Haunted Wasteland

## Part 1

First part was easy to solve using brute force, I created a struct to hold the network elements:

```go
type Element struct {
    left string
    right string
}
```

Used this struct to navigate, starting at the "AAA" element, making sure to loop through the instructions.

## Part 2

Part 2 was a little bit trickier

My first idea was to use the brute force approach (didn't learn with day 5). Instead of storing a single element, I created an array of elements with all `..A` elements and (tried to) navigated until all elements reached the `..Z` landmark. Took to long, so I decided for a new approach.

> It would be perfect if it was that easy to pivot an idea, in reality I tried to tweak a few things befora abandoning the brute force approach.

Next tried to understand the scenario, my first idea was to find how many steps from `..A` to `..Z`, let's call this `a`, and then from `..Z` to `..Z`, let's call it `b`. A linear equation in the form of `y = a + b*x`, it would be easy to find the point of intersection of any two lines, and luckly it would be the same for all lines.

To my surprise I found that `a == b` for all cases, would I be lucky enough to find that `a`'s are prime numbers and simply multiplying then would be enough? I'm not that lucky, it was a huge number, even changed my var to uint64 to hold it. Was not the correct number.

Next idea need to find the less common multiplier, gotta be that number. I was too lazy to create a function for that, simply used the one at ([Golang] Calculate Least Common Multiple (LCM) by GCD)[https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/]. And it simply worked.

## What did I learned about Go?

I'm getting familiar with Go concepts
* I'm still not sure if I'm handling slices the more efficient way.
* Structs are great and easy to use.
* For loops are easy and flexible.

## What did I learned about myself?

Really enjoyed how the ideas came naturally, from brute force, to linear equations, to prime numbers, to LCM.

## TODO

As with the last ones, there is the possibility to join both parts in a single function, reducing the number of lines. Too lazy to do it today.

