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
            parseDescriptionStringToHtml()
            itemStart = itemStart + 10
        }
    });
}

function parseDescriptionStringToHtml() {
    $('#feed-item-list').find('.feed-item-description-tag').each(function (index, element) {
        $(this).html($(this).text())
    })
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