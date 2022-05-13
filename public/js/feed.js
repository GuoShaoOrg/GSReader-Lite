function AddFeedChannelByUserID() {
    let rssLink = $('#rss-link').val()
    let userInfo = getUserInfo()
    let userId = ""
    let token = ""
    if (userInfo !== null) {
        userId = userInfo.uid
        token = userInfo.token
    }
    let postData = {
        link: rssLink,
        userId: userId
    }
    $.ajax({
        headers:{
            Authorization: token + "@@" + userId
        },
        method: 'POST',
        url: '/v1/api/feed/link/uid',
        data: JSON.stringify(postData),
        success: function (data) {
            let jsonData = JSON.parse(data)
            if (jsonData.error === 0) {
                $('#rss-link').val("")
            }
            mdui.snackbar({
                message: jsonData.msg,
                position: 'top',
            })
        },
        error: function (data) {
            let jsonData = JSON.parse(data)
            mdui.snackbar({
                message: jsonData.msg,
                position: 'top',
            })
        }
    })
}