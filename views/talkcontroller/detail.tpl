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
<td><a rel="stylesheet" type="text/css" href="/success" >返回上一级 </td></a>
</table>