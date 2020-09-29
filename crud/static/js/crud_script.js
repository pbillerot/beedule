/**
 * Script.js
 */
$(document).ready(function () {

    // POSITIONNEMENT DERNIERE LIGNE SELECTIONNEE
    // Initialisation du contexte
    var $crud_view = $('#crud_view').val()
    if ($crud_view && $crud_view.length > 0) {
        // Nous sommes dans une vue 
        if (Cookies.get($crud_view)) {
            // Positionnement sur la dernière ligne sélectionnée
            $anchor = $('#' + Cookies.get($crud_view))
            $('html, body').animate({
                scrollTop: $anchor.offset().top - 100
            }, 1000)
            $anchor.css("background-color", "seashell");
        }
    }

    // TABLESORT
    $('table').tablesort()

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
    // Si recherche dans Cookie : aff du input et sélection
    if ($crud_view && $crud_view.length > 0) {
        if (Cookies.get($crud_view + '_search')) {
            $('#crud-search-active').trigger('click');
            $('#crud-search-input').val(Cookies.get($crud_view + '_search'))
            $("#crud-search-input").trigger('keyup')
        }
    }
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
        if (event.target.nodeName == "BUTTON") {
            // pour laisser la main à crud-jquery-button
            // Cas d'un button dans une card
            event.preventDefault();
            return
        }
        // Mémo du contexte dans un cookie
        if ($crud_view && $crud_view.length > 0) {
            if (this.id) {
                Cookies.set($crud_view, this.id)
            } else {
                // on remonte sur <a pour trover l'id
                var ele = this.closest('a');
                Cookies.set($crud_view, ele.id)
            }
        }

        var $target = $(this).data('target');
        if (!$target || $target == '') {
            window.location = $(this).data('url');
        } else {
            window.open($(this).data('url'), $target);
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
        $('.ui.modal')
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
});
