---
title: "dataciteapi(1) user manual | version 1.0.3 6bb4c5f"
author: "R. S. Doiel"
pubDate: 2024-04-04
---

# NAME

dataciteapi

# SYNOPSIS

{appName} [OPTIONS] works DOI

# DESCRIPTION

dataciteapi retrieves "works" from the DataCite API.

dataciteapi is a command line utility to retrieve "works" objects
from the DataCite API. It follows the protocols described at

  https://support.datacite.org/docs/api

NOTE: As of release v1.0.3 you can pass an arXiv id and it will be
converted to the DOI version of arXiv registered with DataCite
under the 10.48550 prefix.

# OPTIONS

-help
: display help

-license
: display license

-mailto string
: set the mailto value for API access

-version
: display app version

# EXAMPLES

Return the works for the doi "10.1037/0003-066x.59.1.29"

~~~
    {appName} -mailto="jdoe@example.edu" \
        works "10.1037/0003-066x.59.1.29"
~~~

Get the data cite record for "arXiv:2202.01037"

~~~
    {appName} -mailto="jdoe@example.edi" \
	    works "arXiv:2202.01037"
~~~
