{{template "header.tpl" .HeaderData}} {{template "menu.tpl" .MenuData}}
<div id="wrapper">
    <div id="page">
        {{ .Content }} {{ range .Posts }}
        <p class="postview">
            {{ .Title }} - {{ formatDate .Date }}
        </p>
        {{ else }}
        <p>
            {{ .NoPostMessage }}
        </p>
        {{ end }}
    </div>
</div>
{{template "footer.tpl" .FooterData}}