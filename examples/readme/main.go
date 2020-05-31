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

