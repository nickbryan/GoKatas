# Go Katas
I always struggle to think of meaningful project ideas to flex my skills. Usually I end up implementing
some sort of game/game engine but these projects are usually too large to find the time to complete. Recently
I have been looking for something to help me explore Go. I figured that doing a bunch of kata style projects 
would be an interesting way to hone my skills and explore Go a bit more. With each kata being fairly small, 
it is easier to find the time to do them around work and life commitments.

This repository is made up of multiple sub-projects that each implement a specific challenge. Each will be
complemented by a suite of tests indicating the desired behavior of each challenge.

Below you can find links to each kata and a brief description of what it entails. As the list grows, I may 
end up breaking them down into categories and restructure the repository.

## Implemented

### Karate Chop
[CodeKata Kata02: Karate Chop](http://codekata.com/kata/kata02-karate-chop/) - Dave Thomas

Implement five different versions of a binary search algorithm that takes an integer search target and a sorted
array of integers. It should return the integer index of the target in the array, or -1 if the target is not
in the array.

* An iterative binary search satisfying the given test cases + benchmarks.
* A recursive binary search satisfying the given test cases + benchmarks.
* A tail recursive binary search satisfying the given test cases + benchmarks.
* A parallel iterative binary search satisfying the given test cases + benchmarks.
* A parallel recursive binary search satisfying the given test cases + benchmarks (hard to come up with a fresh idea).

### Data Munging
[CodeKata Kata04: Data Munging](http://codekata.com/kata/kata04-data-munging/) - Dave Thomas

* __Part One:__ Parse weather.dat and output the day number (column one) with the smallest temperature spread.
* __Part Two:__ Parse football.dat and output the name of the team with the smallest goal difference.
* __Part Three:__ Take the two programs written previously and factor out as much common code as possible,
leaving you with two smaller programs and some kind of shared functionality.

Writing this using TDD allowed me to create some form of abstractions up front (to more easily test), this
made implementing Part Two fairly simple with some refactoring. The difficulty came when looking to write 
e2e tests and covering some of the built in Go functions such as `runtime.Caller`. As has been mentioned
many times before "Prefer duplication over the wrong abstraction"; this is a hard one to get right and there 
will always be trade offs but I feel that this kata is a perfect demonstration of that sentiment. Next time 
I attempt this I will focus on trying to create the "right" abstractions over trying to factor out as much
code as possible.

### Bloom Filters (Spell Checker)
[CodeKata Kata05: Bloom Filters](http://codekata.com/kata/kata05-bloom-filters/) - Dave Thomas

Implement a Bloom filter based spell checker.

