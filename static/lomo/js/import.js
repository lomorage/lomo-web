/*
 * jQuery File Upload Demo
 * https://github.com/blueimp/jQuery-File-Upload
 *
 * Copyright 2010, Sebastian Tschan
 * https://blueimp.net
 *
 * Licensed under the MIT license:
 * https://opensource.org/licenses/MIT
 */

/* global $ */

$(function() {
  'use strict';

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

  $.ajaxSetup({
    headers: {
        "Authorization": "token=" + sessionStorage.getItem("token")
    }
  });

  // Initialize the jQuery File Upload widget:
  $('#fileupload').fileupload({
    // Uncomment the following to send cross-domain cookies:
    //xhrFields: {withCredentials: true},
    url: CONFIG.getUploadUrl()
  });

  // Enable iframe cross-domain access via redirect option:
  $('#fileupload').fileupload(
    'option',
    'redirect',
    window.location.href.replace(/\/[^/]*$/, '/cors/result.html?%s')
  );

  $('#fileupload').fileupload('option', {
    multipart: false,
    // Enable image resizing, except for Android and Opera,
    // which actually support image resizing, but fail to
    // send Blob objects via XHR requests:
    disableImageResize: /Android(?!.*Chrome)|Opera/.test(
      window.navigator.userAgent
    ),
    autoUpload: true,
    maxFileSize: 100*1024*1024*1024,
    //acceptFileTypes: /(\.|\/)(gif|jpe?g|png)$/i,

    getFilesFromResponse: function(data) {
      if (data.result) {
        var media_type = data.files[0].type
        if (media_type === 'video/quicktime') {
          media_type = 'video/mp4'
        }
        return [{
          name: data.result.Name,
          size: data.files[0].size,
          type: media_type,
          url: CONFIG.getAssetUrl(data.result.Name),
          thumbnailUrl: CONFIG.getPreviewUrl(data.result.Name),
          deleteUrl: CONFIG.getAssetUrl(data.result.Name),
          deleteType: 'DELETE'
        }];
      }
      return [];
    },
  });
  // Upload server status check for browsers with CORS support:
  if ($.support.cors) {
    $.ajax({
      type: 'HEAD'
    }).fail(function() {
      $('<div class="alert alert-danger"/>')
        .text('Upload server currently unavailable - ' + new Date())
        .appendTo('#fileupload');
    });
  }
});
