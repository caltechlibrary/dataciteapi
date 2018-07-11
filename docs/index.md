
# USAGE

	dataciteapi [OPTIONS] works DOI

## SYNOPSIS


dataciteapi is a command line utility to retrieve "types" and "works" objects
from the DataCite API. It follows the etiquette suggested at
	
	https://support.datacite.org/docs/api

EXAMPLES

Return the works for the doi "10.1000/xyz123"

	dataciteapi -mailto="jane.doe@example.edu" works "10.1000/xyz123"



## OPTIONS

```
    -generate-markdown-docs   output documentation in Markdown
    -h, -help                 display help
    -l, -license              display license
    -m, -mailto               set the mailto value for API access
    -v, -version              display app version
```


dataciteapi v0.0.1
