<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <title>{{.PageTitle}}</title>
    <link rel="stylesheet" href="/assets/markdown.css" />
    <link rel="stylesheet" href="/assets/style.css" />

    <link rel="stylesheet" href="/assets/highlightjs.css">
    <script src="/assets/highlight.js"></script>
    <script>
        hljs.configure({
            languages: []
        });
        hljs.initHighlightingOnLoad();
    </script>
</head>

<body>
    <header class="disable-default-link">
        {{ if .BlogLogoURL }}
        <div id="header-logo-box">
            <a href="/"><img src="{{ .BlogLogoURL }}" alt="logo" id="header-logo" title="{{ .BlogName }}" /></a>
        </div>
        {{ end }}
        <div><a href="/">{{ .BlogName }}</a></div>
    </header>