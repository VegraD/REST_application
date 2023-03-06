 # Assignment 1: Unisearcher


## Endpoints:

### Uniinfo

This particular endpoint shows the user, upon a request, information about a university. General usage below:
```
Method: GET
Path: /unisearcher/v1/uniinfo/
Request: uniinfo/{:partial_or_complete_university_name}
```

#### Parameters:

{partial_or_complete_university_name} the name/value the user searches for.

#### Responses:

Content-type : application/json

- 200: Everything is OK.
- 204: The request found no results, and returned empty.
- 400: The request had wrong formatting.
- 501: Not implemented.
- 500: Undefined, internal server error.


#### Body (example):
```
[
    {
        "name": "Agriculture and Forestry University",
        "country": "Nepal",
        "isocode": "NP",
        "webpages": [
            "http://www.afu.edu.np/"
        ],
        "languages": {
            "nep": "Nepali"
        },
        "map": "https://www.openstreetmap.org/relation/184633"
    }
]
```
### Neighbourunis

This particular endpoint shows the user, upon a request, information about universities situated in the border countries. General usage below:
```
Method: GET
Path: /unisearcher/v1/neighbourunis/
Request: uniinfo/{country_name}/{:partial_or_complete_university_name}{?limit={number}}
```

#### Parameters:
{country_name} the name of the country of which to find border countries.
{partial_or_complete_university_name} the name/value the user searches for.


#### Responses:

Content-type : application/json

- 200: Everything is OK.
- 204: The request found no results, and returned empty.
- 400: The request had wrong formatting.
- 501: Not implemented.
- 500: Undefined, internal server error.

#### Body (example):

```
[
    {
        "name": "Agriculture and Forestry University",
        "country": "Nepal",
        "isocode": "NP",
        "webpages": [
            "http://www.afu.edu.np/"
        ],
        "languages": {
            "nep": "Nepali"
        },
        "map": "https://www.openstreetmap.org/relation/184633"
    }
]
```

### Diagnostics interface

This particular endpoint shows information about the upkeep of the service.
```
Method: GET
Path: diag/
```

#### Body:

```
{
   "universitiesapi": "<http status code for universities API>",
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": "<time in seconds from the last service restart>"
}
```