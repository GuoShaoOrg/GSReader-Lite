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

var itemStart = 0
var itemSize = 10
function loadFeedItemByUserID() {
    $.ajax({
        method: 'GET',
        url: '/view/all/feed/item/list',
        data: {
            userId: userInfo.uid,
            start: itemStart,
            size: itemSize
        },
        success: function (data) {
            if (data.includes('id="no-more-items"')) {
                $('#load-more-items').hide()
            }
            $('#feed-item-list').append(data)
            itemStart = itemStart + 10
        }
    });
}

function refreshFeedItems() {
    $('#feed-item-list').empty()
    itemStart = 0
    $('#load-more-items').show()
    loadFeedItemByUserID()
}

function AddChannelByLinkWithUserID() {
    $.ajax({
        method: 'GET',
        url: '/view/feed/add',
        success: function (data) {
            $('#add-container').append(data)
        }
    });
}