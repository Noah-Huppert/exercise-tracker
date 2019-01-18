# Design
Workout planner design.

# Table Of Contents
- [Models](#models)
	- [Muscle Group](#muscle-group)
	- [Exercise](#exercise)
	- [Exercise Count](#exercise-count)
	- [Day Plan](#day-plan)
	- [Day Plan Exercise](#day-plan-exercise)
	- [Plan](#plan)
	- [Day Of Week](#day-of-week)
	- [Exercise Record](#exercise-record)

# Models
Database models for row store.  

All models have a few properties in common:

- Have a primary key field name id
- All primary and foreign keys are of type integer and serial
- All fields can not be null unless specified otherwise

## Muscle Group
Identifies a group of muscles on the body.  

Rows are seeded.

Columns:

- `name` (String)

## Exercise
Identifies a movement which targets a muscle group.

Rows are seeded.

Columns:

- `name` (String)
- `includes_weight` (Boolean)
	- Indicates if an Exercise Count must have a weight or not
- `muscle_group_id`

## Exercise Count
Count of sets and repetitions of an exercise.

Columns:

- `exercise_id`
- `sets` (Integer)
- `repetitions` (Integer)
- `weight_pounds` (Integer, Nullable)
	- Optional, required if the associated Exercise's includes_weight field
		is true

## Day Plan
A plan to complete certain exercises on a day.

Columns:

- `name` (String)

## Day Plan Exercise
Exercise in a day plan.

Columns:

- `day_plan_id`
- `exercise_id`
- `min_exercise_count_id` (Nullable)
	- Optional, if not included then the plan exercise does not include
		a range
- `max_exercise_count_id`

## Plan
Plan for a week.

All `_day_plan_id` fields a nullable. If they are null it means that day is a 
rest day.

Plans and related resources are immutable once an Exercise Record exists which 
points towards the Plan.

Columns:

- `name` (String)
- `monday_day_plan_id`
- `tuesday_day_plan_id`
- `wednesday_day_plan_id`
- `thursday_day_plan_id`
- `friday_day_plan_id`
- `saturday_day_plan_id`
- `sunday_day_plan_id`

## Day Of Week
Enumerations which indicates the day of the week.

Values:

- `monday`
- `tuesday`
- `wednesday`
- `thursday`
- `friday`
- `saturday`
- `sunday`

## Exercise Record
Recording of an exercise which tool place.

Columns:

- `plan_id`
- `day_plan_exercise_id`
- `date_time` (Date Time)
- `day_of_week` (Day Of Week)
- `exercise_count_id`
