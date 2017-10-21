<div class=form>
User:{{.NickName}}
  <form action="{{.NickName}}/new" method='POST'>
  发送内容:<br>
  <textarea name="detail" rows="12" cols="32">
  这里输入你要发送的内容！
  </textarea>
  </br>
  <button type="submit">提交</button>
</form>