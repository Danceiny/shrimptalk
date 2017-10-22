<div class=form>
用户 {{.NickName}} 正在回答收到的Topic!
</br>
</br>
</br>
历史消息：<br/>
  {{range .talk}}
{{.Comment}}
<br/>
{{end}}
<br/>
<form action="{{.NickName}}/postanswer" method='POST'>
  <textarea name="detail" rows="16" cols="72">

这里输入你要发送的内容！
  </textarea>
  </br>
  <button type="submit">提交</button>
</form>