# Design
Workout planner design.

# Table Of Contents
- [Collections](#collections)
	- [Plan](#plan)
	- [Exercise Record](#exercise-record)
	- [Exercise](#exercise)
	- [Muscle Group](#muscle-group)
- [Sub Schemas](#sub-schemas)
	- [Exercise Count](#exercise-count)
	- [Day Of Week](#day-of-week)
- [API](#api)
	- [Muscle Group Endpoints](#muscle-group-endpoints)
	- [Exercise Endpoints](#exercise-endpoints)
	- [Day Plan Endpoints](#day-plan-endpoints)

# Collections
Database models for document store.

## Plan
Plan to complete exercises during the week.

Immutable once [Exercise Records](#exercise-record) which reference this 
plan exist.

Schema:

- `_id` (ID)
- `name` (String)
- `day_plans` ([Day Plan](#day-plan)[7])
	- *Holds plans of exercise to complete in a day*
	- *If an index is empty the day is considered a rest day*
	- `exercise_id` (ID)
	- `min_exercise_count` ([Exercise Count](#exercise-count))
		- Optional
		- If not included then the plan does not include a range. Instead the
			`max_exercise_count` will be treated as the prescribed amount
	- `max_exercise_count` ([Exercise Count](#exercise-count))
- `archived` (Boolean)
	- Indicates that the plan is no longer relevant and only exists so past
		[Exercise Records](#exercise-record) remain valid

## Exercise Record
Recording of an exercise which took place.

Schema:

- `_id` (ID)
- `plan_id` (ID)
- `exercise_id` (ID)
- `exercise_count` ([Exercise Count](#exercise-count))
- `date_time` (Date Time)
- `day_of_week` ([Day Of Week](#day-of-week))

## Exercise
Identifies a movement which targets a muscle group.

Values are seeded.

Schema:

- `name` (String)
- `includes_weight` (Boolean)
	- Indicates if an Exercise Count must have a weight or not
- `muscle_group_id`

## Muscle Group
Identifies a group of muscles on the body.  

Values are seeded.

Schema:

- `name` (String)

# Sub Schemas
Schemas used inside collection schemas.

## Exercise Count
Count of sets and repetitions of an exercise.

Schema:

- `exercise_id` (ID)
- `sets` (Integer)
- `repetitions` (Integer)
- `weight_pounds` (Integer, Nullable)
	- Optional
	- Required if the associated exercise's includes_weight field is true

## Day Of Week
Enumeration which indicates the day of the week. Maps day of the week to a 
number.

Values:

- `monday=0`
- `tuesday=1`
- `wednesday=2`
- `thursday=3`
- `friday=4`
- `saturday=5`
- `sunday=6`

# API
Basic HTTP API.

Authentication method TBD.

Request and response bodies are JSON formatted.

Models returned by endpoints will resolve any foreign keys.

## Muscle Group Endpoints
### Get All Muscle Groups
#### Request
GET `/api/v0/muscle_groups`

#### Response
Body:

- `muscle_groups` (Muscle Group[])

## Exercise Endpoints
### Get All Exercises
#### Request
GET `/api/v0/exercises`

#### Response
Body:

- `exercises` (Exercise[])

## Day Plan Endpoints
### Get Day Plan
#### Request
GET `/api/v0/day_plans/:id`

Query parameters:

- `:id` (Integer)
	- ID of day plan to return

#### Response
Body:

- `day_plan` (Day Plan)
