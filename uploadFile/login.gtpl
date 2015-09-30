<html>
<head>
	<title>login page</title>
</head>
<body>
	<form action="http://127.0.0.1:9999/login" method="post">
		<input type="text" name="username">
		<input type="password" name="password">
		<input type="hidden" name="token" value="{{.}}" />
		<input type="submit" value="login" />
	</form>
</body>
</html>