### Named Graphs

When we use keyword GRAPH - reference data from a specific named graph.

------

```rq
SELECT DISTINCT ?g
WHERE
{
    GRAPH ?g { ?s ?p ?o }
}
```
