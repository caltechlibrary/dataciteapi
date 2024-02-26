---
title: "dataciteapi(1) user manual | version 1.0.2 9702707"
author: "R. S. Doiel"
pubDate: 2024-02-26
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

    {appName} -mailto="jdoe@example.edu" \
        works "10.1037/0003-066x.59.1.29"

