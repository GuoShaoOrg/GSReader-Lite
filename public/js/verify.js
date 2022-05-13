$(function () {
    let userInfo = getUserInfo()
    if (userInfo === undefined || userInfo === null) {
        window.location.href = '/view/user/login'
    }
})
