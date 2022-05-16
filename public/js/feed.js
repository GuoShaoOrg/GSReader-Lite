

function parseDescriptionStringToHtml() {
    $('#feed-item-list').find('.feed-item-description-tag').each(function (index, element) {
        $(this).html($(this).text())
    })
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