const $ = mdui.$

function setUserInfo(userInfo) {
    localStorage.setItem("userInfo", JSON.stringify(userInfo))
}

function getUserInfo() {
    let userInfo = JSON.parse(localStorage.getItem("userInfo"))
    let uid = userInfo['uid']
    console.log(uid)
    console.log(userInfo)
    if (uid !== null && uid !== undefined && uid !== '') {
        return userInfo
    }
    return null
}

function cleanUserInfo() {
    localStorage.removeItem("userInfo")
}

function login(account, password) {
    let email = ""
    let mobile = 0
    if (isEmail(account)) {
        email = account
    } else if (isMobile(account)) {
        mobile = account
    }
    let postData = {
        password: password,
        email: email,
        mobile: parseInt(mobile),
    }
    $.ajax({
        method: 'POST',
        url: '/v1/api/user/login',
        data: JSON.stringify(postData),
        success: function (data) {
            console.log(data)
            let jsonData = JSON.parse(data)
            if (jsonData.error !== 0) {
                mdui.snackbar({
                    message: jsonData.msg,
                    position: 'top',
                })
            } else {
                mdui.snackbar({
                    message: '登录成功',
                    position: 'top',
                })
                setUserInfo(jsonData.data[0])
                window.location.href = '/'
            }
        },
        error: function (data) {

        }
    })
}

function isEmail(email) {
    return String(email)
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,:\s@"]+(\.[^<>()[\]\\.,:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        )
}

function isMobile(mobile) {
    let reg=/^1[3-9]\d{9}$/
    return reg.test(mobile)
}