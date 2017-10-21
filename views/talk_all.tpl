<table border="1">
  <tr>
    <th>TalkName</th>
    <th>User</th>
  </tr>
  {{range .Talk}}
  <tr>
    <td><a rel="stylesheet" type="text/css" href="/talk/{{ .TalkNameHex}}" >{{ .TalkNameHex}} </a></td>
    <td>{{ .Now }}</td>
  </tr>
  {{end}}
</table>


