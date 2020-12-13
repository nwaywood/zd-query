# zd-query

zd-query is a command line application for demonstrate efficient in memory queries of data

## Installation

Install the binary to the root of the project with:

`make build`

## Usage

`./zd-query --help` for usage

Example usage commands:

-   `./zd-query users id 1`
-   `./zd-query tickets org_id 117`
-   `./zd-query organizations name Comtext`

## Assumptions and Design Decisions

### Design Decisions

#### Interactive CLI vs non interactive CLI

#### How to acheive better than O(n) query time

### Assumptions

-   `organization`s contain a value called `shared_tickets`. It is assumed that this value has no impact on query results
-   When displaying `ticket`s for a `user`, only the `ticket`s they submitted need to be displayed, not `ticket`s that are assigned to them
-   A `user` should always have an associated `org`, if not, an error will be raised
-   A `ticket` should always have an associated `org` and a submitter, if not, an error will be raised
