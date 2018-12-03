package config

const (
	COOKIE_FILE     = "./.renrenbk_cookies.json"
	ENCRYPT_KEY_URL = "http://login.renren.com/ajax/getEncryptKey"
	LOGIN_URL       = "http://www.renren.com/ajaxLogin/login?1=1&uniqueTimestamp={ts}"
	ICODE_URL       = "http://icode.renren.com/getcode.do?t=web_login&rnd={rnd}"
	ICODE_FILEPATH  = "./static/icode.jpg"
)
