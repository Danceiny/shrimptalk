<table border="0">
  <tr>
    <th>NickName</th>
  </tr>
  {{range .Users}}
  <tr>
    <td><a rel="stylesheet" type="text/css" href="/login/{{ .NickNameHex}}" >{{ .NickNameHex}} </a></td>
  </tr>
  {{end}}
</table>


