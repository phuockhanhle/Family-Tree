# Model

## Todo
- Write test
- Restructure functions and variables

## Doing
- Split database migration to dedicate branch [K]
- Modify People [K]
  - Call ra person return id
  - Keep List children, remove parents
- Add TreeGroups{Direct, Indirect, InheritFather, InheritMother} + mechanism for IsSameRoot [T]
- Migrate to database
  - PeopleManager 
  - Create relations table 
    - for spouse relation: husband first, wife second column
    - for parent relation: parent first, child second column

# View