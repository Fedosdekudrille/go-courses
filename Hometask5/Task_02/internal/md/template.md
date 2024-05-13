# Report for {{date}}
## Total Revenue: {{round .GetTotal}}
## Highest Revenue Product: {{with .GetHighestRevProduct}}{{.Name}} ({{round .Total}}){{end}}
## Products
{{$id := 0}}{{range $, $value := .}}{{$id = sum $id 1}}{{$id}}. {{$value.Name}} = {{round $value.Total}}
{{end}}