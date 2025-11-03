# cli-habit-tracker
CLI habit tracker written in Go with Sqlite as db.

### Functionality
- add habit to habits table in SQlite
- update count everytime you do the habit
- delete old habits
- list habits

### Functions
Add function:
- Name: AddHabit
- Inputs: name of habit
- Returns:
- Error cases:

Do function:
- Name: AddCount
- Inputs: name of habit
- Returns: String with new count
- Error cases: Try to update habit that is not in db.

List function:
- Name: ListHabits
- Inputs: db with database pointer?
- Returns: String with all habits and their count
- Error cases: Db fails to respond.

Delete function:
- Name: DeleteHabit
- Inputs: name of habit
- Returns: string confirmind that habit x was deleted
- Error case: Entry was not deleted
