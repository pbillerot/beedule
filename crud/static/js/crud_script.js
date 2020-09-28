/**
 * Script.js
 */
$(document).ready(function () {
    // Positionnement sur la ligne dernièrement sélectionnée
    // var $anchor_url = $('#anchor').val()
    // if ($anchor_url.length > 0) {
    //     $anchor = $('#' + $anchor_url)
    //     $('html, body').animate({
    //         scrollTop: $anchor.offset().top - 100
    //     }, 1000)
    //     $anchor.css("background-color", "seashell");
    // }

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

    // tablesort
    $('table').tablesort()

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

    // Affichage du champ de recherche
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
        $("#bee-table tr").filter(function () {
            $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
        });
    });

    // Appel URL dans la même fenêtre
    $('.crud-jquery-url').on('click', function (event) {
        if (event.target.nodeName == "BUTTON") {
            // pour laisser la main à crud-jquery-button
            // Cas d'un button dans une card
            event.preventDefault();
            return
        }
        // Mémo du contexte dans un cookie
        if ($crud_view.length > 0) {
            Cookies.set($crud_view, this.id)
        }

        var $target = $(this).data('target');
        if (!$target || $target == '') {
            window.location = $(this).data('url');
        } else {
            window.open($(this).data('url'), $target);
        }
        event.preventDefault();
    });

    // Appel URL dans la même fenêtre
    $('.crud-jquery-button').on('click', function (event) {
        var $target = $(this).data('target');
        if (!$target || $target == '') {
            window.location = $(this).data('url');
        } else {
            window.open($(this).data('url'), $target);
        }
        event.preventDefault();
    });

    // Validation du formulaire
    $('.crud-jquery-submit').on('click', function (event) {
        $('form', document).submit();
        event.preventDefault();
    });

    // Action
    // Demande confirmation
    $('.crud-jquery-action').on('click', function (event) {
        var $url = $(this).data('url');
        if ($(this).data('confirm') == true) {
            $('#crud-action').html($(this).html());
            $('.ui.modal')
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

    // Suppression d'un enregistrement
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

    // Recherche dans la LIST
    $('.crud-searchable').searchable({
        searchField: '#crud-search-input',
        selector: '.crud-item-searchable',
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

});
