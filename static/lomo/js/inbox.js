
function fetchInbox() {
    $.ajax({
        url: CONFIG.getInboxUrl(),
        dataType: 'json'
    }).then(
        function (resp) {
            console.log( "fetchInbox user/group succeed!");
            console.log(resp);

            var arrayOfPromises = [];

            fetchSharedAssetsUsers(resp, arrayOfPromises);
            fetchSharedAssetsGroups(resp, arrayOfPromises);

            $.when.apply($, arrayOfPromises).then(function() {
                console.log("all done!");
            });
        },

        function( xhr, status, errorThrown ) {
            alert( polyglot.t("FetchError") );
            console.log( "Error: " + errorThrown );
            console.log( "Status: " + status );
            console.dir( xhr );
        }
    )

    function fetchSharedAssetsGroups(resp, arrayOfPromises) {
        var groupCnt = resp.Groups.length;
        for (var i = 0; i < groupCnt; ++i) {
            console.log('fetching assets share from group' + resp.Groups[i]);
            arrayOfPromises.push(
                $.ajax({
                    url: CONFIG.getGroupInboxUrl(resp.Groups[i]),
                    dataType: 'json'
                }).then(
                    function (resp) {
                        console.log("getGroupInboxUrl succeed!");
                        console.log(resp);
                        var recordCnt = resp.Records.length;
                        for (var k = 0; k < recordCnt; ++k) {
                            var assetRec = resp.Records[k];
                            elem = '<a href="' + CONFIG.getInboxAssetUrl(assetRec.ID)
                                + '" title="' + assetRec.AssetID
                                + '" data-type="' + getMimeType(assetRec.AssetID) + '" data-gallery>'
                                + '<img class="lazy" data-src="' + CONFIG.getInboxPreviewUrl(assetRec.ID)
                                + '"/></a>';
                            $("#links").append(elem);
                        }
                        $('.lazy').lazy();
                    },

                    function (xhr, status, errorThrown) {
                        alert(polyglot.t("FetchError"));
                        console.log("Error: " + errorThrown);
                        console.log("Status: " + status);
                        console.dir(xhr);
                    }
                )
            );
        }
    }

    function fetchSharedAssetsUsers(resp, arrayOfPromises) {
        var userCnt = resp.Users.length;
        for (var i = 0; i < userCnt; ++i) {
            console.log('fetching assets share from user' + resp.Users[i]);
            arrayOfPromises.push(
                $.ajax({
                    url: CONFIG.getUserInboxUrl(resp.Users[i]),
                    dataType: 'json'
                }).then(
                    function (resp) {
                        console.log("getUserInboxUrl succeed!");
                        console.log(resp);
                        var recordCnt = resp.Records.length;
                        for (var k = 0; k < recordCnt; ++k) {
                            var assetRec = resp.Records[k];
                            elem = '<a href="' + CONFIG.getInboxAssetUrl(assetRec.ID)
                                + '" title="' + assetRec.AssetID
                                + '" data-type="' + getMimeType(assetRec.AssetID) + '" data-gallery>'
                                + '<img class="lazy" data-src="' + CONFIG.getInboxPreviewUrl(assetRec.ID)
                                + '"/></a>';
                            $("#links").append(elem);
                        }
                        $('.lazy').lazy();
                    },

                    function (xhr, status, errorThrown) {
                        alert(polyglot.t("FetchError"));
                        console.log("Error: " + errorThrown);
                        console.log("Status: " + status);
                        console.dir(xhr);
                    }
                )
            );
        }
    }
}

function getMimeType(filename) {
    var ext = filename.split('.').pop().toLowerCase();

    var extToMimes = {
        'jpg': 'image/jpeg',
        'jpeg': 'image/jpeg',
        'png': 'image/png',
        'heif': 'image/heif',
        'heic': 'image/heic',
        'zip': 'livephoto/zip',
        '3gp': 'video/3gpp',
        '3g2': 'video/3gpp2',
        'mov': 'video/mp4',
        'mp4': 'video/mp4',
        'avi': 'video/x-msvideo',
        'mpg': 'video/mpeg',
        'mpeg': 'video/mpeg',
        'webm': 'video/webm',
    }

    if (extToMimes.hasOwnProperty(ext)) {
        return extToMimes[ext];
    }
    return '';
}

$(function() {
    $.ajaxSetup({
        headers: {
            "Authorization": "token=" + sessionStorage.getItem("token")
        }
    });

    if (sessionStorage.getItem("token") === null) {
        document.location.href = '/';
    }

    $('a#logout').text(polyglot.t("Logout") + sessionStorage.getItem("username"))
    $('a#logout').click(function() {
    sessionStorage.removeItem("token");
    sessionStorage.removeItem("userid");
    sessionStorage.removeItem("username");
    document.location.href = '/';
    });

    fetchInbox();

    $('#blueimp-gallery').data('fullScreen', 'true');

    $('#blueimp-gallery').on('slide', function(event, index, slide) {
        // Gallery slide event handler
        $('video').trigger('pause');
        // console.log($("div.slide")[index]);
        // console.log($("div.slide").eq(index).find('video').length);
        $("div.slide").eq(index).find('video').trigger('play');
    })
});
