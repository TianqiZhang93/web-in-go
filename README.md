# web-in-go
A golang web server.

## file organization
- common: const values ...
- config: config models and methods
- controller: controller aster router
- dao: dao level
- env_configs: configs for different enviroments(test dev online...)
- env_databases: databases for different enviroments(test dev online...)
- middleware: middlewares
- models: vo do po models
- route: web api router
- service: service logic
- tools: small tools


## model
- vo: view objects in controller
- bo: business objects in service
- po: persistant objects in dao

## work flow
route -(middlewares)-> controller -> service -> dao -> database
