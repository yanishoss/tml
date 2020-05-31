# Training Markup Language

> *TML* stands for *Training Markup Language*. *TML* is a markup language aiming at providing a simple and clean way to describe a training/workout with the exercises and the performances.

## Example
```tml
I really enjoyed my workout because of my TML logs!

# Squat
150kg x 5 @ 7
160kg x 5 @ 8
160kg x 5 @ 8
160kg x 5 @ 9
my squat session felt really great! I hit my RPEs seamlessly!

# Deadlift
150kg x 5 @ 7
160kg x 5 @ 8
160kg x 5 @ 8
160kg x 5 @ 9
my deadlift session felt really great! I hit my RPEs seamlessly!

# Push ups
1 count x 25 x 4
it felt really hard, trust me!
```

## Installation
```shell script
go get -u github.com/yanishoss/tml
```

## Parsing in Go
```go
package main

import (
    "fmt"
    "github.com/yanishoss/tml"
)

func main() {
    input := `
I really enjoyed my workout because of my TML logs!

# Squat
150kg x 5 @ 7
160kg x 5 @ 8
160kg x 5 @ 8
160kg x 5 @ 9
my squat session felt really great! I hit my RPEs seamlessly!

# Deadlift
150kg x 5 @ 7
160kg x 5 @ 8
160kg x 5 @ 8
160kg x 5 @ 9
my deadlift session felt really great! I hit my RPEs seamlessly!
`

    workout, err := tml.Parse(input, tml.WithDefaultConfig())

    if err != nil {
        panic(err)
    }
    
    // Play with the workout object
    
    fmt.Println("Workout's comment: ", *workout.Comment)
    fmt.Println("Workout's exercises number: ", len(workout.Exercises))

    return 
}
```

## Types
```go
package parser

type Row struct {
	Sets   int      `json:"sets"`
	Reps   int      `json:"reps"`
	RPE    *float64 `json:"rpe,omitempty"`
	Weight float64  `json:"weight"`
	Unit   string   `json:"unit"`
}

type Exercise struct {
	Name    string  `json:"name"`
	Rows    []Row   `json:"rows"`
	Comment *string `json:"comment,omitempty"`
}

type Workout struct {
	Comment   *string    `json:"comment,omitempty"`
	Exercises []Exercise `json:"exercises"`
}

type Config struct {
	DefaultUnit string
	ValidUnits  []string
	RPERange    [2]float64
}
```

## [Language Specification](https://github.com/yanishoss/tml/blob/master/SPECIFICATION.md "TML Specification")