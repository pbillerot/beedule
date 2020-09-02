/**
 * Script.js
 */
$(document).ready(function () {
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
        $('#crud-search').toggle();
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
        })
        ;
});
