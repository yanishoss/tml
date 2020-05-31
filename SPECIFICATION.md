# TML Specification

## Introduction

*TML* stands for *Training Markup Language*. *TML* is a markup language aiming at providing a simple and clean way to describe a training/workout with the exercises and the performances.

## Language Specification

```tml
[<comments>]
#<exercise name>
<weight>[<unit>] x <reps> [x <sets>] [@ <RPE>]
[<comments>]
```

\<comments\>: string, ends when an exercise is encountered, when an empty line is encountered, or at the EOF

\<exercise name\>: string, ends when a newline character is encountered

\<weight\>: float

\<unit\>: "kg", "lbs", "count, "min" or "s", "count" by default

\<reps\>: integer in [0, +inf]

\<sets\>: integer in [0, +inf], "1" by default

\<RPE\>: float in [0; 10], none by default
