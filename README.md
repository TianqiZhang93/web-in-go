# web-in-go
A golang web server.

## file organization
- bin: output
- common: const values ...
- config: config models and methods
- controller: controller aster router
- coverage: UT output
- dao: dao level
- dist: build output
- env_configs: configs for different enviroments(test dev online...)
- env_databases: databases for different enviroments(test dev online...)
- middleware: middlewares
- models: vo do po models
- router: web api router
- service: service logic
- tools: small tools


## model
- vo: view objects in controller
- bo: business objects in service
- po: persistant objects in dao

## work flow
router -(middlewares)-> controller -> service -> dao -> database
