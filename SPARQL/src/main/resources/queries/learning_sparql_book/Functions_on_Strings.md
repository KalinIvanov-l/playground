### Functions on strings

-----

`STRSTARTS`

Purpose: STRSTARTS(string, compare_string) checks if a string starts with the compare string.
Example: STRSTARTS("United Kingdom","Uni") returns true, meaning the string starts with Uni.

`STRENDS`

Purpose: STRENDS(string, compare_string) checks if the string ends with the compare string.
Example: STRENDS("India", "ia") returns true, meaning the string ends with ia.

`STRBEFORE`

Purpose: STRBEFORE(string, compare_string) returns the part of the string before the compare string.
Example: STRBEFORE("South Africa", "th") returns Sou. If the compare string is not found, it returns empty.

`STRAFTER`

Purpose: STRAFTER(string, compare_string) returns the part of the string after the compare string.
Example: STRAFTER("Thailand", "ai") returns land. If the compare string is not found, it returns empty.

`STRLEN`

Purpose: STRLEN(string) returns the length of a string.
Example: STRLEN(?country) where ?country is China will return 5.

`CONTAINS`

Purpose: CONTAINS (string, compare_string) checks if the string contains the compare string.
Example: CONTAINS("Madagascar", "dag") returns true, meaning the string contains dag.

`SUBSTR`

Purpose: SUBSTR(string, start_position, string_length) returns a substring of a string starting at the position marked by start_position, and will have the given length supplied through string_length.
Example: SUBSTR(?country, 7, 4) where ?country is Costa Rica will return Rica. Omitting the length as in SUBSTR(?country, 7) returns Rica, i.e., the remaining string from the starting position.

`UCASE`

Purpose: UCASE(string) returns the string in upper case.
Example: UCASE("Japan") will be returned as JAPAN.

`LCASE`

Purpose: LCASE(string) returns the string in lower case.
Example: LCASE("UK") will be returned as uk.

`CONCAT`

Purpose: CONCAT(string1, string2, ...) returns the concatenation of two or more strings.
Example: CONCAT("Republic", " of ", "Ireland") returns Republic of Ireland.

`ENCODE_FOR_URI`

Purpose: ENCODE_FOR_URI(string) converts any special characters in the string for use in a URL for the web.
Example: ENCODE_FOR_URI("Washington DC") returns Washington%20DC.

`LANGMATCHES`

Purpose: LANGMATCHES(language_tag, language_range) checks if the language tag matches the language range. See langMatches for more info.
Example: FILTER LANGMATCHES(LANG(?label), "en") filters language tags in the English language. The language tags may contain a region as well, e.g. en-IE. A language range of * matches any non-empty language-tag string.

`REGEX`

Purpose: REGEX (string, pattern, flag) checks if the string matches the pattern.
The pattern may contain different special characters.
See regex for more info.

`REPLACE`

Purpose: REPLACE (string, pattern, replacement, flag) returns the string after replacing all occurrences of pattern in string with replacement. pattern is interpreted the same as in REGEX. flag is optional.
Example: REPLACE("United_Kingdom", "[_]", " ") removes all the underscores from the original string and replaces them with spaces.
