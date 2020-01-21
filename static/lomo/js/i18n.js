var polyglot = new Polyglot();

const messages = {
    zh: {
        'LoginError': '登陆失败!',
        'FetchError': '服务不可用',
        'Logout': '退出'
    },
    en: {
        'LoginError': 'Login Failure!',
        'FetchError': 'Service unavailable',
        'Logout': 'Logout '
    }
}
var userLang = navigator.language || navigator.userLanguage;
if (userLang.startsWith("zh")) {
    polyglot.extend(messages.zh)
} else {
    polyglot.extend(messages.en) 
}