<nav>
	{{ range .Links }}
	<a href="{{ .Dest }}" title="{{ .Title }}">{{ .Text }}</a>
	{{ end }}
</nav>