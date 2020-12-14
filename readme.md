# zd-query

zd-query is a command line application for demonstrating efficient in memory queries of data. It loads the source data from the `/data` directory and then lets you query that data by any field.

## Installation

Install the binary to the root of the project with:

`make build`

## Usage

`./zd-query --help` for usage

Example usage commands:

-   `./zd-query users id 1`
-   `./zd-query tickets org_id 117`
-   `./zd-query organizations name Comtext`

To achieve the desired performance, this application caches data to make subsequent query fast. If you need to regenerate this cache (e.g. because the source data has changed), this can be done with `make clear-cache`

## Assumptions and Design Decisions

### Design Decisions

For this application there were two main design decisions that I made. Below I have summarised the rational behind my decisions

#### Interactive CLI vs Non-Interactive CLI

The first design I had to make was what kind of UX I wanted for the application, an interactive or non-interactive CLI. An interactive CLI would have made it easier to achieve the desired performance, but I decided to go with a non-interactive cli for the following reasons:

-   I made the assumption that when the app is being used manually (not via a script), a user would not be making a large number of queries at a time, thus minimising the benefit of an interactive CLI
-   The application usage would be more familar to users, since non-interactive CLI's are the standard way to build CLI applications
-   Being the standard way, it would also enable me to leverage CLI frameworks to ensure the application is used correctly without having to code as much custom error handling
-   To allow for the application to be used in scripts, making it more versatile (for generating reports etc.)

#### How to achieve better than O(n) query time

With the primary user-based decision decided (building a non-interactive CLI), the next main decision was how to achieve the desired performance of better than O(n) for all queries. I decided to achieve the required performance by caching secondary indexes of all the fields. The following steps lead me to this design decision:

1. The source data is in the form of 3 arrays, and the only way to get better than O(n) query performance on arrays is if they are sorted. However sorting an array is O(nlogn) and since every value is querable an array would need to be resorted every time it's queried by a different value. Therefore having the data as arrays for querying it isn't satisfactory for performance unless a copy of each array is made, each one sorted by a different field, which is quite an impracticle solution
1. This lead me to achieving the desired performance via secondary indexing instead. If all the fields are indexed then any query would be O(1)
1. The issue with secondary indexing is that to build the indexes I need to iteratate through the arrays, which is O(n). So even though the query performance is O(1), an overall query would be O(n) since the indexes need to be built before each query. Note that this wouldn't have been as much of an issue for an interactive CLI so it would only need to build the indexes once when the application is first started up, but its still a performance hit that you don't want to put on the user if you don't have to.
1. Therefore I am caching the secondary indexes that I build so that the performance hit of building the indexes is only incurred once, the first time the application is run. After that, the indexes will be loaded from the cache and will only be regenerated if the cache is cleared

### Assumptions

-   `organization`s contain a value called `shared_tickets`. It is assumed that this value has no impact on query results
-   When displaying `ticket`s for a `user`, only the `ticket`s they submitted need to be displayed, not `ticket`s that are assigned to them
-   A `user` should always have an associated `org`, if not, an error will be raised
-   A `ticket` should always have an associated `org` and a submitter, if not, an error will be raised
