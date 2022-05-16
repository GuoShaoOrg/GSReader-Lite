const $ = mdui.$

function setUserInfo(userInfo) {
    localStorage.setItem("userInfo", JSON.stringify(userInfo))
}

function getUserInfo() {
    let userInfo = JSON.parse(localStorage.getItem("userInfo"))
    if (userInfo === undefined || userInfo === null) {
        return null
    }
    let uid = userInfo['uid']
    if (uid !== null && uid !== undefined && uid !== '') {
        return userInfo
    }
    return null
}

function cleanUserInfo() {
    localStorage.removeItem("userInfo")
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

function parseDescriptionStringToHtml() {
    $('#feed-item-list').find('.feed-item-description-tag').each(function (index, element) {
        $(this).html($(this).text())
    })
}

function getSubChannelListTmpl() {
    let userInfo = getUserInfo()
    let userId = ""
    if (userInfo !== null) {
        userId = userInfo.uid
    }
    $.ajax({
        method: 'GET',
        url: '/view/feed/sub_list',
        data: {
            userId: userId,
            start: 0,
            size: 10,
        },
        success: function (data) {
            $('#sub-channel-drawer-list').append(data)
        }
    });
}