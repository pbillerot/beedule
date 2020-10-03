/**
 * Script.js
 */
$(document).ready(function () {

    var isUsed = false;

    // TABLESORT
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
        var $url = $(this).data('url')
            + '?sortid=' + this.id.substring(4) // col_<id>
            + '&sortdirection=' + $sortdirection
        window.location = $url;
        event.preventDefault();
    });

    // $('table').tablesort()
    // $.tablesort.DEBUG = true;
    // $('table').on('tablesort:complete', function (event, tablesort) {
    //     if ($crud_view && $crud_view.length > 0) {
    //         // console.log(tablesort.$sortCells[tablesort.index].id)
    //         Cookies.set($crud_view + '_sort_id', tablesort.$sortCells[tablesort.index].id);
    //         Cookies.set($crud_view + '_sort_direction', tablesort.direction);
    //     }
    // });

    // RECHERCHE
    $('#crud-search-active').on('click', function (event) {
        $('#crud-search').show();
        $('#crud-header').hide();
        $('#crud-search-active').hide();
        $('#crud-search-input').focus();
    });
    // Recherche plein texte dans le body de la table
    // Les lignes sans le mot sont cachées
    $("#crud-search-input").on("keyup", function () {
        var value = $(this).val().toLowerCase();
        if ($crud_view && $crud_view.length > 0) {
            Cookies.set($crud_view + '_search', value)
        }
        $("#bee-table tr").filter(function () {
            $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
        });
    });
    // Recherche dans la LIST CARD
    $('.crud-searchable').searchable({
        searchField: '#crud-search-input',
        selector: '.crud-card-searchable',
        childSelector: '.searchable',
        show: function (elem) {
            elem.fadeIn(100);
        },
        hide: function (elem) {
            elem.fadeOut(100);
        },
        onSearchActive: function (elem, term) {
            elem.show();
        },
        onSearchEmpty: function (elem) {
            elem.show();
        }
    })
    // Fermer la recherche
    $('#crud-search-close').on('click', function (event) {
        $('#crud-search').hide();
        $('#crud-header').show();
        $('#crud-search-active').show();
        $('#crud-search-input').val('')
        $("#crud-search-input").trigger('keyup')
    });

    // CLIC URL
    $('.crud-jquery-url').on('click', function (event) {
        if ( isUsed ) {
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

        var $target = $(this).data('target');
        var $url = $(this).data('url');
        if (!$target || $target == '') {
            window.location = $url;
        } else {
            window.open($url, $target);
        }
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
            }).modal('show');
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
        if (Cookies.get($crud_view + '_search')) {
            $('#crud-search-active').trigger('click');
            $('#crud-search-input').val(Cookies.get($crud_view + '_search'))
            $("#crud-search-input").trigger('keyup')
        }
        // Positionnement sur la dernière ligne sélectionnée
        if (Cookies.get($crud_view)) {
            $anchor = $('#' + Cookies.get($crud_view))
            if ($anchor.length) {
                $('html, body').animate({
                    scrollTop: $anchor.offset().top - 100
                }, 1000)
                $anchor.addClass("crud-list-selected");
            }
        }
    }

});
