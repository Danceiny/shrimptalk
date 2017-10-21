<div class=form>
用户 {{.NickName}} 正在输入信息!
</br>
</br>
</br>
  <form action="{{.NickName}}/new" method='POST'>
  <textarea name="detail" rows="16" cols="72">
这里输入你要发送的内容！
  </textarea>
  </br>
  <button type="submit">提交</button>
</form>