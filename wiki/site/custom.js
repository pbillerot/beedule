/**
 * bee-cript.js
 */
$(document).ready(function () {

    // Bouton édition conditionné si victor dans l'url
    $('#top-github-link').each(function () {
        // ajout de la class bee-edit-open
        var $link = $(this).find("a");
        if (window.location.pathname.indexOf("hugo/") > -1) {
            $link.addClass("bee-edit-open");
        } else{
            $(this).attr("style", "display: none;");
        }
    });
    // appel de l'éditeur dans une fenêtre détachée
    $(".bee-edit-open").on('click', function (event) {
        var $path = $(this).attr('href');
        window.open($path, '_blank', 'left=' + (screen.availWidth - 1024 - 5) + ',top=5,width=1024,height=1014,scrolling=yes,scrollbars=yes,resizeable=yes');
        event.preventDefault();
    });
    // les liens externes seront ouverts dans un nouvel onglet
    $('a[href]').each(function () {
        // ajout de la class bee-edit-open
        var $href = $(this).attr("href");
        if ($href.indexOf("http") > -1) {
            $(this).attr("target", "_blank");
        }
    });

});