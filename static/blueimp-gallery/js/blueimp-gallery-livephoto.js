/*
 * blueimp Gallery Video Factory JS
 * https://github.com/blueimp/Gallery
 *
 * Copyright 2013, Sebastian Tschan
 * https://blueimp.net
 *
 * Licensed under the MIT license:
 * https://opensource.org/licenses/MIT
 */

/* global define */

;(function(factory) {
    'use strict'
    if (typeof define === 'function' && define.amd) {
      // Register as an anonymous AMD module:
      define(['./blueimp-helper', './blueimp-gallery'], factory)
    } else {
      // Browser globals:
      factory(window.blueimp.helper || window.jQuery, window.blueimp.Gallery)
    }
  })(function($, Gallery) {
    'use strict'

    var model = (function() {
		var URL = window.webkitURL || window.mozURL || window.URL;

		return {
			getEntries : function(url, onend) {
				zip.createReader(new zip.HttpReader(url, true), function(zipReader) {
					zipReader.getEntries(onend);
				}, onerror);
			},
			getEntryFile : function(entry, onend, onprogress) {
				var writer, zipFileEntry

				function getData() {
					entry.getData(writer, function(blob) {
						var blobURL = URL.createObjectURL(blob);
						onend(blobURL);
					}, onprogress);
				}
			    writer = new zip.BlobWriter();
			    getData();
			}
		};
	})();

    $.extend(Gallery.prototype.options, {
        // The class for video content elements:
        videoContentClass: 'video-content',
    })

    var handleSlide = Gallery.prototype.handleSlide

    $.extend(Gallery.prototype, {
      handleSlide: function(index) {
        handleSlide.call(this, index)
      },

      livephotoFactory: function(obj, callback) {
        var that = this
        var options = this.options
        var livephoto = document.createElement('img')
        var errorArgs = [
          {
            type: 'error',
            target: livephoto
          }
        ]

        var url = this.getItemProperty(obj, options.urlProperty)
        var title = this.getItemProperty(obj, options.titleProperty)
        if (title) {
            livephoto.title = title
        }

		model.getEntries(url, function(entries) {
            var arrayOfPromises = [];
			entries.forEach(function(entry) {
                var entryPromise = new Promise((resolve, reject) => {
                    var ext = entry.filename.split('.').pop().toLowerCase();
                    if (ext == 'jpg') {
                        model.getEntryFile(entry, function(blobURL) {
                            livephoto.setAttribute("src", blobURL);
                            console.log('unzip ' + entry.filename)
                            resolve();
                        });
                    } else if (ext == 'mov') {
                        model.getEntryFile(entry, function(blobURL) {
                            livephoto.setAttribute("data-live-photo", blobURL);
                            console.log('unzip ' + entry.filename)
                            resolve();
                        });
                    } else if (ext == 'heic') {
                        model.getEntryFile(entry, function(blobURL) {
                            fetch(blobURL)
                            .then((res) => res.blob())
                            .then((blob) => heic2any({
                                blob,
                                toType: "image/jpeg",
                                quality: 0.3,
                                multiple: true
                            }))
                            .then((conversionResult) => {
                                conversionResult.forEach(image => {
                                    livephoto.setAttribute("src", URL.createObjectURL(image));
                                    console.log('unzip ' + entry.filename)
                                    resolve();
                                });
                            })
                            .catch((e) => {
                                console.log(e);
                                reject();
                            });
                        });
                    } else {
                        console.log('Not supported extension: ' + ext);
                        reject();
                    }
                });

                arrayOfPromises.push(entryPromise);
            });

            Promise.all(arrayOfPromises).then(function() {
                console.log("live photo loading done! " + title);

                var elems = LivePhotos.initialize(livephoto)
                elems.forEach(function(ele) {
                    ele.container.classList.add(that.options.slideContentClass);
                    ele.container.classList.add(that.options.videoContentClass);
                    //ele.play();
                })

                that.setTimeout(callback, [
                    {
                      type: 'load',
                      target: elems[0].container
                    }
                ])
            });
        });

        return livephoto
      }
    })

    return Gallery
  })
