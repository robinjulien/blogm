{{template "header.tpl" .HeaderData}}
{{template "menu.tpl" .MenuData}}
<div id="wrapper">
<div id="page" class="markdowned">
	{{ .Content }}
</div>
</div>
{{template "footer.tpl" .FooterData}}
