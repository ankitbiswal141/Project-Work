
# Procedures
- **Step 1: Prepare the environment (Such as initialize go mod with your GitHub path)**
- **Step 2: Install the required packages (Mentioned below)**
- **Step 3: Create a new project (With the name you want)**
- **Step 4: Include all the necesary subdirectories and files (Mentioned below)**

# Structure
  ## Subdirectories
    - /account : Contains the account related functions
    - /catalog : Contains the product and order catalog related functions
    - /graphql : Contains the GraphQL schema, models, mutation, and resolvers
    - /order   : Contains the order related functions in context to the products

  ## Files
    -- account/
       >

    -- catalog/
       >
    
    -- graphql/
       > app.dockerfile
       > gqlgen.yml
       > graph.go
       > models.go
       > mutation_resolver.go
       > query_resolver.go
       > // models_gen.go ( Will be added after installation of the mentioned libraries and dependencies. )

    --order/
      >

    - docker-compose.yaml
    
    //These two will be added after initailizing go mod
    - go.mod
    - go.sum


# Libraries Install And Dependencies:
-> github.com/jinzhu/gorm
-> github.com/jinzhu/gorm/dialects/mysql
-> github.com/99designs/gqlgen
-> github.com/99designs/gqlgen generate