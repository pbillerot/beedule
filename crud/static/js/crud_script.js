/**
 * Script.js
 */
$(document).ready(function () {

    var isUsed = false;

    // Coloriage syntaxique
    if ($("#codemirror-markdown").length != 0) {
        var myCodeMirror = CodeMirror.fromTextArea(
            document.getElementById('codemirror-markdown')
            , {
                lineNumbers: false,
                lineWrapping: true,
                mode: 'yaml-frontmatter',
                readOnly: false,
                theme: 'eclipse',
                viewportMargin: 20
            }
        );
    }

    // Collapse
    $('.crud-collapse').on('click', function (event) {
        var portlet = $(this).closest('div');
        if ($(this).hasClass('open')) {
            portlet.find('.icon').removeClass("open");
        } else {
            portlet.find('.icon').addClass("open");
        }
        portlet.find('.list').toggle();
        portlet.find('.message').toggle();
        // portlet.find('.content').toggle();
        event.preventDefault();
    });

    // RECHERCHE
    // Ouvrir la recherche
    $('#crud-search-active').on('click', function (event) {
        $('#crud-search').show();
        // $('#crud-header').hide();
        $('#crud-search-active').hide();
        $('#search').focus();
        event.preventDefault();
    });
    // Fermer la recherche
    $('#crud-search-close').on('click', function (event) {
        $('#crud-search').hide();
        // $('#crud-header').show();
        $('#crud-search-active').show();
        $('#search').val('');
        $('#crud-form-search').val('')
        $('#crud-form-searchstop').val('true')
        $('form').submit();
        event.preventDefault();
    });
    // Validation par touche entrée
    $('#search').on('keypress', function (e) {
        if (e.which == 13) {
            $('#crud-search-go').trigger('click');
        }
    });
    // Lancement de la recherche
    $('#crud-search-go').on('click', function (event) {
        var value = $('#search').val().toLowerCase();
        if (value.length > 0) {
            $('#crud-form-search').val(value)
        } else {
            $('#crud-form-search').val('')
            $('#crud-form-searchstop').val('true')
        }
        $('form').submit();
        event.preventDefault();
    });

    // TRI COLONNE DE LA TABLE
    $(document).on('click', '.crud-ajax-sort', function (event) {
        var $sortdirection = "ascending"
        if (!$(this).hasClass('sorted')) {
            $(this).closest('tr').find('.sorted').removeClass('sorted');
            $(this).closest('tr').find('.ascending').removeClass('ascending');
            $(this).closest('tr').find('.descending').removeClass('descending');
            $(this).addClass("sorted")
            $(this).addClass($sortdirection)
        } else {
            // on inverse le tri
            if ($(this).hasClass('ascending')) {
                $(this).closest('tr').find('.ascending').removeClass('ascending');
                $sortdirection = "descending"
                $(this).addClass($sortdirection)
            }
        }
        $('#crud-form-sortid').val(this.id.substring(4))
        $('#crud-form-sortdirection').val($sortdirection)
        $('form').submit();
        event.preventDefault();
    });

    // CLIC URL
    $('.crud-jquery-url').on('click', function (event) {
        if (isUsed) {
            event.preventDefault();
            return
        }
        if (event.target.nodeName == "BUTTON") {
            // pour laisser la main à crud-jquery-button
            // Cas d'un button dans une card
            event.preventDefault();
            return
        }
        // Mémo du contexte dans un cookie
        if ($crud_view && $crud_view.length > 0) {
            Cookies.set($crud_view, this.id)
            $(this).addClass("crud-list-selected");
        }

        var $url = $(this).data('url');
        window.location = $url;
        event.preventDefault();
    });

    // CLIC BUTTON URL
    $('.crud-jquery-button').on('click', function (event) {
        var $target = $(this).data('target');
        if (!$target || $target == '') {
            window.location = $(this).data('url');
        } else {
            window.open($(this).data('url'), $target);
        }
        event.preventDefault();
    });

    // VALIDATION FORMULAIRE
    $('.crud-jquery-submit').on('click', function (event) {
        $('form', document).submit();
        event.preventDefault();
    });

    // ACTION DEMANDE CONFIRMATION
    $('.crud-jquery-action').on('click', function (event) {
        var $url = $(this).data('url');
        if ($(this).data('confirm') == true) {
            $('#crud-action').html($(this).html());
            $('#crud-modal-confirm')
                .modal({
                    closable: false,
                    onDeny: function () {
                        return true;
                    },
                    onApprove: function () {
                        $('form').attr('action', $url);
                        $('form', document).submit();
                    }
                }).modal('show');
        } else {
            // Sans demande de confirmation
            $('form').attr('action', $url);
            $('form', document).submit()
        }
        event.preventDefault();
    });

    // CLIC IMAGE POPUP
    $('.crud-popup-image').on('click', function (event) {
        isUsed = true;
        // Mémo du contexte dans un cookie
        if ($crud_view && $crud_view.length > 0) {
            var $anchor = $(this).closest('.card');
            Cookies.set($crud_view, $anchor.attr('id'))
            $(this).closest('.cards').find('.crud-list-selected').removeClass('crud-list-selected');
            $anchor.addClass("crud-list-selected");
        }

        var $url = $(this).data('url');
        $('#crud-image').attr('src', $url)
        $('#crud-modal-image')
            .modal({
                closable: true,
                onHide: function () {
                    isUsed = false;
                    return true;
                }
            }).modal('show');
        event.preventDefault();
    });

    // CLIC IMAGE EDITOR POPUP
    $('.crud-popup-image-editor').on('click', function (event) {
        isUsed = true;

        var $url = $(this).data('url');
        var $key = $(this).data('key');
        const config = {
            language: 'fr',
            tools: ['adjust', 'effects', 'filters', 'rotate', 'crop', 'resize', 'text'],
            translations : {
                fr: {
                    'toolbar.download': 'Valider'
                },
            }
        };
        var mime = $url.endsWith('.png') ? 'image/png' : 'image/jpeg';
        // https://github.com/scaleflex/filerobot-image-editor
        const ImageEditor = new FilerobotImageEditor(config, {
            onBeforeComplete: (props) => {
                console.log("onBeforeComplete", props);
                console.log("canvas-id", props.canvas.id);
                var canvas = document.getElementById(props.canvas.id);
                var dataurl = canvas.toDataURL(mime, 1);
                $("#" + $key).val(dataurl);
                $("#" + $key + "_img").attr('src', dataurl);
                return false;
            },
            onComplete: (props) => {
                console.log("onComplete", props);
                return true;
            }
        });
        ImageEditor.open($url);
        event.preventDefault();
    });

    // SUPPRESSION D'UN ENREGISTREMENT
    $('.crud-jquery-delete').on('click', function (event) {
        $('#crud-modal-confirm')
            .modal({
                closable: false,
                onDeny: function () {
                    return true;
                },
                onApprove: function () {
                    $('form', document).submit();
                }
            }).modal('show');
        event.preventDefault();
    });

    // IHM SEMANTIC
    $('.ui.checkbox').checkbox();
    $('.ui.radio.checkbox').checkbox();
    $('.ui.dropdown').dropdown();
    $('select.dropdown').dropdown();
    $('.message .close')
        .on('click', function () {
            $(this)
                .closest('.message')
                .transition('fade')
                ;
        }
        );
    $('.hide')
        .on('click', function () {
            $(this)
                .closest('.message')
                .transition('fade')
                ;
        }
        );

    // Toaster
    $('#toaster')
        .toast({
            class: $('#toaster').data('color'),
            position: 'bottom right',
            message: $('#toaster').val()
        });
    // Calendar
    $('#standard_calendar')
        .calendar({
            ampm: false,
            text: {
                days: ['D', 'L', 'M', 'M', 'J', 'V', 'S'],
                months: ['Janvier', 'Février', 'Mars', 'Avril', 'Mai', 'Juin', 'Juillet', 'Août', 'Septembre', 'Octobre', 'Novembre', 'Decembre'],
                monthsShort: ['Jan', 'Fev', 'Mar', 'Avr', 'Mai', 'Juin', 'Juil', 'Aou', 'Sep', 'Oct', 'Nov', 'Dec'],
                today: 'Aujourd\'hui',
                now: 'Maintenant',
                am: 'AM',
                pm: 'PM'
            },
            // formatter: {
            //     date: function (date, settings) {
            //         if (!date) return '';
            //         var day = date.getDate();
            //         var month = date.getMonth() + 1;
            //         var year = date.getFullYear();
            //         return year + '-' + month + '-' + day;
            //     }
            // }
        })
        ;

    // APRES CHARGEMENT HTML ET JAVASCRIPT
    // CONTEXTE DE LA VUE
    var $crud_view = $('#crud_view').val()
    if ($crud_view && $crud_view.length > 0) {
        // Si recherche dans Cookie : aff du input et sélection
        var $search = $('#search').val();
        if ($search != "") {
            $('#crud-search-active').trigger('click');
        }
        // Positionnement sur la dernière ligne sélectionnée
        // voir ligne avec CrudIndexAnchor
        if (Cookies.get($crud_view)) {
            $anchor = $('#' + Cookies.get($crud_view))
            if ($anchor.length) {
                $('html, body').animate({
                    scrollTop: $anchor.offset().top - 100
                }, 1000)
                $anchor.addClass("crud-list-selected");
                // Collpase du folder
                if ($anchor.hasClass("message")) {
                    $anchorCollapse = $('.' + Cookies.get($crud_view))
                    $anchorCollapse.trigger("click");
                }
            }
        }
    }
});
