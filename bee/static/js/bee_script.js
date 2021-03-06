/**
 * Script.js
 */
$(document).ready(function () {
  var isUsed = false;

  // CLIC IMAGE POPUP
  var $eddy_refresh = $('#eddy_refresh').val();
  var $eddy_rubriques = $('#eddy_rubriques').val();

  // CODEMIRROR : coloration syntaxique et auto-complétion
  // https://codemirror.net/
  // ajout de la liste des rubriques dans $eddy_dico
  // et dans $keyword pour le coloriage syntaxique
  // https://github.com/erosman/CodeMirror-plus
  var $keyword = {}
  if ($eddy_rubriques) {
    rubriques = $eddy_rubriques.split(',');
    for (id in rubriques) {
      $eddy_dico.push({ text: rubriques[id], displayText: 'r_' + rubriques[id] });
      $keyword[rubriques[id]] = 'keyword';
    }
  }
  // déclaration de la foncion autocomplete
  CodeMirror.commands.autocomplete = function (cm) {
    CodeMirror.showHint(cm, CodeMirror.hint.eddy, { list: $eddy_dico });
  }
  CodeMirror.hint.eddy = function (cm, options) {
    var list = options.list || [];
    var cursor = cm.getCursor();
    var currentLine = cm.getLine(cursor.line);
    var start = cursor.ch;
    var end = start;
    while (end < currentLine.length && /[\w$]+/.test(currentLine.charAt(end))) ++end;
    while (start && /[\w$]+/.test(currentLine.charAt(start - 1))) --start;
    var curWord = start != end && currentLine.slice(start, end);
    var regex = new RegExp('' + curWord, 'i');
    var result = {
      list: (!curWord ? list : list.filter(function (item) {
        if (typeof item == "object") {
          return item.displayText.match(regex);
        } else {
          return item.match(regex);
        }
      })).sort(),
      from: CodeMirror.Pos(cursor.line, start),
      to: CodeMirror.Pos(cursor.line, end)
    };
    return result;
  };
  // Activation de CODEMIRROR
  if ($("#codemirror-markdown").length != 0) {
    var myCodeMirror = CodeMirror.fromTextArea(
      document.getElementById('codemirror-markdown')
      , {
        lineNumbers: false,
        lineWrapping: true,
        mode: 'yaml',
        readOnly: false,
        autoRefresh: true,
        theme: 'eclipse',
        tabSize: 2,
        autofocus: true,
        viewportMargin: 20,
        keyword: $keyword,
      }
    );
    myCodeMirror.setOption("extraKeys", {
      Tab: function (cm) {
        var spaces = Array(cm.getOption("indentUnit") + 1).join(" ");
        cm.replaceSelection(spaces);
      },
      "Ctrl-S": function (cm) {
        var cursor = cm.getCursor();
        sessionStorage.setItem("ch", cursor.ch);
        sessionStorage.setItem("line", cursor.line);
        $("#button_validate").trigger('click');
      },
      "Ctrl-Space": "autocomplete",
      "Ctrl-/": function (cm) {
        cm.toggleComment();
      }
    });
    myCodeMirror.on("change", function (cm) {
      var cursor = cm.getCursor();
      sessionStorage.setItem("ch", cursor.ch);
      sessionStorage.setItem("line", cursor.line);
      $('#button_validate').removeAttr('disabled');
    })
    myCodeMirror.focus();
    if (sessionStorage.getItem("line")) {
      myCodeMirror.setCursor({ 
        line: parseInt(sessionStorage.getItem("line")), 
        ch: parseInt(sessionStorage.getItem("ch")) });
    }
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
    $(this).attr('disabled', 'disabled');
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
  // CLIC IMAGE POPUP
  $('.crud-popup-chart').on('click', function (event) {
    isUsed = true;
    // Mémo du contexte dans un cookie
    if ($crud_view && $crud_view.length > 0) {
      var $anchor = $(this).closest('.card');
      Cookies.set($crud_view, $anchor.attr('id'))
      $(this).closest('.cards').find('.crud-list-selected').removeClass('crud-list-selected');
      $anchor.addClass("crud-list-selected");
    }

    var $html = $(this).html();
    var canvasParent = $('#crud-chart')
    canvasParent.html($html)
    $('#crud-modal-chart')
      .modal({
        closable: true,
        onHide: function () {
          isUsed = false;
          return true;
        },
        onVisible: function () {
          drawChart(canvasParent.children("canvas"));
        }
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

  // Actualisation de la fenetre parent
  if ($eddy_refresh && $eddy_refresh.length > 0) {
    var val = Cookies.get($eddy_refresh)
    if (val == "true") {
      window.opener.location.reload();
      Cookies.remove($eddy_refresh);
    }
  } // end eddy_refresh

  /**
   * Ouverture d'une fenêtre en popup
   * TODO voir si accepter par tout les browsers
   */
  $(document).on('click', '.eddy-window-open', function (event) {
    // Préparation window.open
    var $height = $(this).data("height") ? $(this).data("height") : 'max';
    var $width = $(this).data("width") ? $(this).data("width") : 'large';
    var $posx = $(this).data("posx") ? $(this).data("posx") : 'left';
    var $posy = $(this).data("posy") ? $(this).data("posy") : '3';
    var $target = $(this).attr("target") ? $(this).attr("target") : 'eddy-win';
    var $url = $(this).data("url");
    if (window.opener == null) {
      window.open($url
        , $target
        , computeWindow($posx, $posy, $width, $height, false));
    } else {
      window.opener.open($url
        , $target
        , computeWindow($posx, $posy, $width, $height, false));
    }
    event.preventDefault();
  });

  /**
   * Fermeture de la fenêtre popup
   */
  $(document).on('click', '.crud-jquery-close', function (event) {
    if ($('#button_validate').length > 0
      && $('#button_validate').attr('disabled') != "disabled") {
      $('#crud-action').html("Abandonner les modifications ?");
      $('#crud-modal-confirm')
        .modal({
          closable: false,
          onDeny: function () {
            return true;
          },
          onApprove: function () {
            window.close();
          }
        }).modal('show');
    } else {
      window.close();
    }
    event.preventDefault();
  });

});

/**
 * Calcul du positionnement et de la taille de la fenêtre sur l'écran
 * @param {string} posx left center right ou px
 * @param {int} posy px
 * @param {string} pwidth max large xlarge ou px
 * @param {string} pheight max ou px
 * @param {srtien} full_screen yes no 0 1
 */
function computeWindow(posx, posy, pwidth, pheight, full_screen) {
  if (full_screen) {
    pheight = screen.availHeight - 70;
    pwidth = screen.availWidth - 6;
  }
  var height = pheight != null ? (/^max$/gi.test(pheight) ? screen.availHeight - 120 : pheight) : 830;
  var width = 900;
  if (pwidth != null) {
    width = pwidth;
    if (/^max$/gi.test(pwidth)) width = screen.availWidth - 6;
    if (/^large$|^l$/gi.test(pwidth)) width = 1024;
    if (/^xlarge$|^xl$/gi.test(pwidth)) width = 1248;
  } // end largeur
  var left = 3;
  if (posx != null) {
    left = posx;
    if (/^left$/gi.test(posx)) left = 3;
    if (/^right$/gi.test(posx)) left = screen.availWidth - width - 18;
    if (/^center$/gi.test(posx)) left = (screen.availWidth - width) / 2;
  } // end posx
  var top = 6
  if (posy != null) {
    height = screen.availHeight - posy - 10;
    top = posy;
  }

  return 'left=' + left + ',top=' + top + ',height=' + height + ',width=' + width + ',scrolling=yes,scrollbars=yes,resizeable=yes';
}

/**
 * Blocage du carriage return dans un champ input par exemple
 * @param {object event} event 
 */
function blockCR(event) {
  if (event.keyCode == 13) {
    event.preventDefault();
  }
}