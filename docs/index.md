
# USAGE

	dataciteapi [OPTIONS] works DOI

## SYNOPSIS


_dataciteapi_ is a command line utility to retrieve "works" objects
from the DataCite API. It follows the protocols described at
```	
  https://support.datacite.org/docs/api
```	


## DESCRIPTION


_dataciteapi_ is a command line utility to retrieve "works" objects
from the DataCite API. It follows the protocols described at
```	
  https://support.datacite.org/docs/api
```	


## OPTIONS

Below are a set of options available.

```
    -generate-manpage    generate man page
    -generate-markdown   output documentation in Markdown
    -h, -help            display help
    -l, -license         display license
    -m, -mailto          set the mailto value for API access
    -v, -version         display app version
```


## EXAMPLES


Return the works for the doi "10.1037/0003-066x.59.1.29"

```	
    dataciteapi -mailto="jdoe@example.edu" \
        works "10.1037/0003-066x.59.1.29"
```	



dataciteapi v0.0.5
