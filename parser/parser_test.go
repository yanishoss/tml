package parser

import (
	"github.com/yanishoss/tml/lexer"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	l := lexer.New(`
some comments about the workout
some comments about the workout
some comments about the workout
some comments about the workout


# Squat



60kg x 3 x 5 @ 6.5
60kg x 3 x 5 @ 6
60kg x 3 x 5
60kg x 3 @ 6
60 x 3



good squat session
good squat session
good squat session
good squat session



# Deadlift



60kg x 3 x 5 @ 6
60kg x 3 x 5 @ 6.4


good deadlift session
good deadlift session
good deadlift session
good deadlift session



`)

	getFloat := func(f float64) *float64 {
		return &f
	}

	getString := func(s string) *string {
		return &s
	}

	checkString := func(s1, s2 *string) bool {
		if (s1 != nil && s2 == nil) || (s2 != nil && s1 == nil) {
			return false
		}

		if s1 == s2 {
			return true
		}

		return *s1 == *s2
	}

	checkFloat := func(s1, s2 *float64) bool {
		if (s1 != nil && s2 == nil) || (s2 != nil && s1 == nil) {
			return false
		}

		if s1 == s2 {
			return true
		}

		return *s1 == *s2
	}

	correctWorkout := Workout{
		Comment: getString("some comments about the workout\nsome comments about the workout\nsome comments about the workout\nsome comments about the workout\n"),
		Exercises: []Exercise{
			{
				Name: "Squat",
				Rows: []Row{
					{5, 3, getFloat(6.5), 60, "kg"},
					{5, 3, getFloat(6), 60, "kg"},
					{5, 3, nil, 60, "kg"},
					{1, 3, getFloat(6), 60, "kg"},
					{1, 3, nil, 60, "kg"},
				},
				Comment: getString("good squat session\ngood squat session\ngood squat session\ngood squat session\n"),
			},

			{
				Name: "Deadlift",
				Rows: []Row{
					{5, 3, getFloat(6), 60, "kg"},
					{5, 3, getFloat(6.4), 60, "kg"},
				},
				Comment: getString("good deadlift session\ngood deadlift session\ngood deadlift session\ngood deadlift session\n"),
			},
		},
	}

	p := New(l, Config{
		DefaultUnit: "kg",
		ValidUnits:  []string{"kg"},
		RPERange:    [2]float64{0, 11},
	})

	w, err := p.Parse()

	if err != nil {
		t.Error(err)
	}

	if !checkString(w.Comment, correctWorkout.Comment) {
		t.Error("field \"Workout.Comment\" not correct")
	}

	for i, e := range w.Exercises {
		if !checkString(e.Comment, correctWorkout.Exercises[i].Comment) {
			t.Errorf("field \"Workout.Exercises[%d].Comment\" not correct\n", i)
		}

		if e.Name != correctWorkout.Exercises[i].Name {
			t.Errorf("field \"Workout.Exercises[%d].Name\" not correct\n", i)
		}

		for j, row := range e.Rows {
			if !checkFloat(row.RPE, w.Exercises[i].Rows[j].RPE) {
				t.Errorf("field \"Workout.Exercises[%d].Rows[%d].RPE\" not correct\n", i, j)
			}

			if row.Reps != w.Exercises[i].Rows[j].Reps {
				t.Errorf("field \"Workout.Exercises[%d].Rows[%d].Reps\" not correct\n", i, j)
			}

			if row.Unit != w.Exercises[i].Rows[j].Unit {
				t.Errorf("field \"Workout.Exercises[%d].Rows[%d].Unit\" not correct\n", i, j)
			}

			if row.Weight != w.Exercises[i].Rows[j].Weight {
				t.Errorf("field \"Workout.Exercises[%d].Rows[%d].Weight\" not correct\n", i, j)
			}

			if row.Sets != w.Exercises[i].Rows[j].Sets {
				t.Errorf("field \"Workout.Exercises[%d].Rows[%d].Sets\" not correct\n", i, j)
			}
		}
	}
}
