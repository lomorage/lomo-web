var polyglot = new Polyglot();

const messages = {
    zh: {
        'LoginError': '登陆失败!',
        'FetchError': '服务不可用',
        'Logout': '退出',
        'Conflict': '文件已存在',
        'DeleteFail': '文件删除失败',
        'DeleteConfirm': '确定删除文件?',
        'OK': '确定',
        'Cancel': '取消'
    },
    en: {
        'LoginError': 'Login Failure!',
        'FetchError': 'Service unavailable',
        'Logout': 'Logout ',
        'Conflict': 'File already exists',
        'DeleteFail': 'Delete file failure',
        'DeleteConfirm': 'Confirm deletion?',
        'OK': 'OK',
        'Cancel': 'Cancel'
    }
}
var userLang = navigator.language || navigator.userLanguage;
if (userLang.startsWith("zh")) {
    polyglot.extend(messages.zh)
} else {
    polyglot.extend(messages.en) 
}