{{template "header.tpl" .HeaderData}}
{{template "menu.tpl" .MenuData}}
<div id="page">
	{{ .Content }}
	{{ range .Posts }}
		<div class="postview">
			{{ .Title }} - {{ formatDate .Date }}
		</div>
	{{ else }}
		<div>
			No post here
		</div>
	{{ end }}
</div>
{{template "footer.tpl" .FooterData}}
