{{template "header.tpl" .HeaderData}} {{template "menu.tpl" .MenuData}}
<div id="wrapper">
    <div id="page disable-default-link">
        {{ if .InvalidPageRequestMessage }}
        <div class="error-block">{{ .InvalidPageRequestMessage }}</div>
        {{ else }} {{ range .Posts }}
        <div class="postview disable-default-link">
            <h2><a href="{{ .Link }}">{{ .Title }}</a></h2>
            <p>{{ formatDate .Date }}</p>
        </div>
        {{ else }}
        <div id="nopost">
            {{ .NoPostMessage }}
        </div>
        {{ end }}
        <div class="page-number-component">
            {{ if gt .NumPosts 0 }} {{ if gt .PageNumber 1 }}
            <a href="?p={{ sub .PageNumber 1 }}" class="pagination-button-link"><button class="pagination-button">&lt;</button></a> {{ end }}
            <div class="page-number">{{ .PageNumber }} / {{ .MaxPage }}</div>
            {{ if gt .MaxPage .PageNumber }}
            <a href="?p={{ add .PageNumber 1 }}" class="pagination-button-link"><button class="pagination-button">&gt;</button>
            </a>{{ end }} {{ end }}
        </div>
        {{ end }}
    </div>
</div>
{{template "footer.tpl" .FooterData}}