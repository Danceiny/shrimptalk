<table border="0">
  <tr>
    <th>NickName</th>
    <th>Comment</th>
  </tr>
  {{range .Detail }}
  <tr>
    <td>{{ .NickName }}</td>
    <td>{{ .Comment }}</td>
  </tr>
  {{end}}
</table>