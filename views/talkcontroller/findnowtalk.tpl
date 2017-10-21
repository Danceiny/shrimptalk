<table border="0">
  <tr>
    <th>The topics I received</th>
  </tr>
  {{range .Talks}}
  <tr>
    <td><a rel="stylesheet" type="text/css" href="/talk/mytalk" >{{ .TalkNameHex}} </td></a>
  </tr>
  {{end}}
</table>


