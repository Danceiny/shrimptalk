<table border="0">
  <tr>
    <th>TalkName</th>
    <th>User</th>
    <th>Hot</th>
  </tr>
  {{range .Talk}}
  <tr>
    <td><a rel="stylesheet" type="text/css" href="/talk/{{ .TalkNameHex}}" >{{ .TalkNameHex}} </a></td>
    <td>{{ .Now }}</td>
    <td>{{ .Max }}</td>
  </tr>
  {{end}}
</table>