# Server on GO

## Migration algorithm for windows

[Using scoop](https://scoop.sh/)

1. Run `scoop install migrate` for installation migrate cli.

2. Run `migrate -version` for check migrate cli version.

3. Run `migrate create -ext sql -dir ./migrations -seq init` to create initial migrations to the selected directory.

4. Run `migrate -path ./migrations -database 'postgres://postgres:secret@localhost:5432/postgres?sslmode=disable' up`
migration up method to the selected directory.

## GO CLI
$ `go mod tidy` 
>The command adds all the missing modules needed to build the packages and dependencies of the current module;
>>removes unused modules that do not provide any relevant packages;
>>>adds all missing entries to go.sum and removes unnecessary ones.
>
>`-v`
>>The flag causes tidy to output information about remote modules to the standard error output (std err).
>
>`-e`
>>The flag causes go mod tidy to try to continue despite errors that occur when downloading packages.