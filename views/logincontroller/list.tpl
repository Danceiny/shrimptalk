<table border="0">
  <tr>
    <th>NickName</th>
  </tr>
  {{range .Users}}

  <tr>
    <td><a rel="stylesheet" type="text/css" href="/login/{{ .NickNameHex}}" >{{ .NickNameHex}} </a></td>
  </tr>
  {{end}}
<div>
<tr>
    <td><a rel="stylesheet" type="text/css" href="/register" >Register </a></d>
  </tr>
</div>
</table>


