# modelgen
gqlgen modelgen plugin with gorm and validation

How To use
```
//add directive to schema file
directive @meta(
    gorm: String,
) on OBJECT | FIELD_DEFINITION | ENUM_VALUE | INPUT_FIELD_DEFINITION | ENUM | INPUT_OBJECT | ARGUMENT_DEFINITION

directive @ValidateMeta(
    Validate: String,
) on OBJECT | FIELD_DEFINITION | ENUM_VALUE | INPUT_FIELD_DEFINITION | ENUM | INPUT_OBJECT | ARGUMENT_DEFINITION

//How to use in type
type User{
    Id: Number! @meta(gorm: "primary_key;auto_increment;not_null") // use gorm
    Avatar: Text!
    phone_number: Text!
    is_active: Bool!
    name: Text! @ValidateMeta(Validate:"required") //use validation
    birth_date: Time!
    national_code: Text!
    Roles(page: Number,pageSize:Number):[Role] @goField(forceResolver:true) @meta(gorm: "many2many:users_roles")
    CreatedAt:CreatedAt
    UpdatedAt:UpdatedAt
    DeletedAt:DeletedAt
}
```