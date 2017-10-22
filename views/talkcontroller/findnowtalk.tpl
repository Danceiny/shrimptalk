<table border="0">
  <tr>
    <th>The topics I received</th>
  </tr>
  {{range .Talks}}
  <tr>
    <td><a rel="stylesheet" type="text/css" href="/talk/{{.TalkNameHex}}/answer" >{{ .TalkNameHex}} </td></a>
  </tr>
  {{end}}
</table>


