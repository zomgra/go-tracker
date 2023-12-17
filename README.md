# go-tracker Mykyta Serdiuk


For start program use:

1. `make db`

2. `make migrate-up`

3. `make run`


For testing use:

1. `make db` _if you already dont use it_

2. `make test_db` 
 
3. `make migrate_up_test`

4. `make test`


Listen app uri:

`localhost:8000/api/shipment/{barcode}` -- GET
`localhost:8000/api/shipment?quantity=X` -- POST


Dependencies for use:

* Requared
_Docker_
_go_

* Optional
_migrate_ 
