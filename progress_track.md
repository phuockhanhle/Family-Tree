# Model

## Todo
- Write test
- Restructure functions and variables

## Doing
- [x] Split database migration to dedicate branch
- [ ] Implement with Neo4j
  - [ ] Connect go driver with Neo4j
  - [ ] Implement operations: MATCH, CREATE, SET
  - [ ] Implement tests

## Later
- Migrate to SQL
  - PeopleManager 
  - Create relations table 
    - for spouse relation: husband first, wife second column
    - for parent relation: parent first, child second column
- Modify People
  - Call ra person return id
  - Keep List children, remove parents
- Add TreeGroups{Direct, Indirect, InheritFather, InheritMother} + mechanism for IsSameRoot [T]

# View
## Features nice-to-have
- View with timeline