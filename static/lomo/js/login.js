'use strict';

document.getElementById('login-form').addEventListener('submit', function(e) {
    e.preventDefault();
});

document.getElementById('submit').addEventListener('click', calc(calcWasm));

function loadScript(src, onload, onerror) {
    var el = document.createElement('script');
    el.src = src;
    el.onload = onload;
    el.onerror = onerror;
    document.body.appendChild(el);
}

function login(hashedPwd) {
    var auth = btoa($('#username').val() + ":" + hashedPwd + ":web")
    $.ajaxSetup({
        headers: {
            "Authorization": "Basic " + auth
        }
    });
    $.ajax({
        url: CONFIG.getLoginUrl()
    })
    .done(function (json) {
        log( "Login succeed! save token " + json.Token);
        sessionStorage.setItem("userid", json.Userid);
        sessionStorage.setItem("token", json.Token);
        sessionStorage.setItem("username", $('#username').val());
        document.location.href = '/gallery'
    })
    .fail(function( xhr, status, errorThrown ) {
        alert( "Sorry, there was a problem!" );
        log( "Error: " + errorThrown );
        log( "Status: " + status );
        console.dir( xhr );
    });
}

function getArg() {
    return {
        pass: $('#password').val(),
        salt: $('#username').val() + '@lomorage.lomoware',
        time: 3,
        mem: 4096,
        hashLen: 32,
        parallelism: 1,
        type: 2 //Argon2id
    };
}

var logTs = 0;

function log(msg) {
    if (!msg) {
        return;
    }
    
    var elapsedMs = Math.round(performance.now() - logTs);
    var elapsedSec = (elapsedMs / 1000).toFixed(3);
    var elapsed = leftPad(elapsedSec, 6);
    console.log('[' + elapsed + '] ' + msg)
}

function leftPad(str, len) {
    str = str.toString();
    while (str.length < len) {
        str = '0' + str;
    }
    return str;
}

function clearLog() {
}
