<table border="0">
  <tr>
    <th>TalkName</th>
  </tr>
  {{range .Talks}}
  <tr>
 <td><a rel="stylesheet" type="text/css" href="/talk/{{.TalkNameHex}}" >{{ .TalkNameHex}} </td></a>
  </tr>
  {{end}}
<td><a rel="stylesheet" type="text/css" href="/success" >返回上一级 </td></a>
  
</table>


