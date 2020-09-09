{{template "header.tpl" .HeaderData}}
{{template "menu.tpl" .MenuData}}
<div id="page">
	{{ .Content }}
</div>
{{template "footer.tpl" .FooterData}}
