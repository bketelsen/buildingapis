# Group pet

## Pet [/pet]

## Pet By PetId [/pet/{petId}]

+ Parameters
    + petId (, required)

        ID of pet to return{LONG}


### getPetById [GET]

Find pet by ID

+ Response 200 (application/json)

successful operation

    + Attributes (Pet)

+ Response 400 

Invalid ID supplied

+ Response 404 

Pet not found

