var polyglot = new Polyglot();

const messages = {
    zh: {
        'LoginError': '登陆失败!',
        'FetchError': '服务不可用',
        'Logout': '退出',
        'Conflict': '文件已存在'
    },
    en: {
        'LoginError': 'Login Failure!',
        'FetchError': 'Service unavailable',
        'Logout': 'Logout ',
        'Conflict': 'File already exists'
    }
}
var userLang = navigator.language || navigator.userLanguage;
if (userLang.startsWith("zh")) {
    polyglot.extend(messages.zh)
} else {
    polyglot.extend(messages.en) 
}