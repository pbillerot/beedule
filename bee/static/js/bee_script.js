/**
 * Script.js
 */
$(document).ready(function () {
  var isUsed = false;
  var $crud_view = $('#crud_view').val();

  // CLIC IMAGE POPUP
  var $eddy_refresh = $('#eddy_refresh').val();
  var $eddy_rubriques = $('#eddy_rubriques').val();
  var $eddy_mode = $('#eddy_mode').val(); // yaml, sql

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
    var regex = new RegExp('' + curWord + '', 'i');
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
        mode: $eddy_mode,
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
        ch: parseInt(sessionStorage.getItem("ch"))
      });
    }
  }

  // PARTAGE
  if (navigator.share) {
    $('.crud-share').on('click', function (event) {
      if (navigator.share) {
        navigator.share({
          title: document.title,
          text: $(this).data('text'),
          url: $(this).data('url'),
        })
          .then(() => console.log('Successful share'))
          .catch((error) => console.log('Error sharing', error));
      }
      event.preventDefault();
    });
  } else {
    // On cache le menu Partage car non accessible
    //$('.crud-share').hide();
  } // end navigator.share

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

  // RECHERCHE dans une vue 
  // Ouvrir la recherche 
  $('.crud-search-active').on('click', function (event) {
    var $content = $(this).closest('.crud-search-div');
    $content.find('.crud-search').show();
    $content.find('.header').hide();
    $content.find('.meta').hide();
    $content.find('input').focus();
    event.preventDefault();
  });
  // Fermer la recherche
  $('.crud-search-close').on('click', function (event) {
    var $content = $(this).closest('.crud-search-div');
    $content.find('.crud-search-input').val('');
    $content.find('.crud-search').hide();
    $content.find('.header').show();
    $content.find('.meta').show();
    if ($content.find('.crud-search-input-1').val().length == 0) {
      // abandon d'une recherche
      event.preventDefault();
      return;
    }
    $content.find('.crud-search-go').trigger('click');
    event.preventDefault();
  });
  // Validation par touche entrée
  $('.crud-search-input').on('keypress', function (e) {
    var $content = $(this).closest('.crud-search-div');
    if (e.which == 13) {
      $content.find('.crud-search-go').trigger('click');
    }
  });
  // Envoi de la recherche au serveur
  $('.crud-search-go').on('click', function (event) {
    var $content = $(this).closest('.crud-search-div');
    var $datas = new FormData();
    var $url = $content.find('.crud-search-input').data("url");
    $datas.append("search", $content.find('.crud-search-input').val().toLowerCase());
    $datas.append("_xsrf", $("#xsrf").val());
    var request =
      $.ajax({
        type: "POST",
        url: $url,
        data: $datas,
        dataType: 'json',
        cache: false,
        contentType: false,
        processData: false,
        beforeSend: function () {
          //Code à jouer avant l'appel ajax en lui même
        }
      });
    request.done(function (response) {
      //Code à jouer en cas d'éxécution sans erreur
      console.log(response);
    });
    request.fail(function (response) {
      //Code à jouer en cas d'éxécution en erreur
      console.log(response);
    });
    request.always(function () {
      //Code à jouer après done OU fail dans tous les cas 
      window.location.reload();
    });
    event.preventDefault();
  });
  // Boucle pour faire apparaître les champs de recherche avec une valeur
  $('.crud-search-input').each(function (index) {
    var $content = $(this).closest('.crud-search-div');
    if ($(this).val().length > 0) {
      // backup search dans crud-search-input-1 afin de traiter l'abandon d'une recherche
      $content.find('.crud-search-input-1').val($(this).val());
      $content.find('.crud-search-active').trigger('click');
    }
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
    var $sortid = this.id.substring(4)

    var $datas = new FormData();
    var $node = $(this).closest('table');
    $datas.append("sortid", $sortid);
    $datas.append("sortdirection", $sortdirection);
    $datas.append("_xsrf", $("#xsrf").val());
    var request =
      $.ajax({
        type: "POST",
        url: "/bee/sort/" + $node.data("app") + "/" + $node.data("table") + "/" + $node.data("view"),
        data: $datas,
        dataType: 'json',
        cache: false,
        contentType: false,
        processData: false,
        beforeSend: function () {
          //Code à jouer avant l'appel ajax en lui même
        }
      });
    request.done(function (response) {
      //Code à jouer en cas d'éxécution sans erreur
      console.log(response);
    });
    request.fail(function (response) {
      //Code à jouer en cas d'éxécution en erreur
      console.log(response);
    });
    request.always(function () {
      //Code à jouer après done OU fail dans tous les cas 
      window.location.reload();
    });

    event.preventDefault();
  });
  // TRI retour au tri par défaut de la vue CLICK-DROIT
  $(".crud-unsort").on("press", function (event) {
    var $datas = new FormData();
    var $node = $(this).closest('table');
    $datas.append("sortid", "");
    $datas.append("sortdirection", "");
    $datas.append("_xsrf", $("#xsrf").val());
    var request =
      $.ajax({
        type: "POST",
        url: "/bee/sort/" + $node.data("app") + "/" + $node.data("table") + "/" + $node.data("view"),
        data: $datas,
        dataType: 'json',
        cache: false,
        contentType: false,
        processData: false,
        beforeSend: function () {
          //Code à jouer avant l'appel ajax en lui même
        }
      });
    request.done(function (response) {
      //Code à jouer en cas d'éxécution sans erreur
      console.log(response);
    });
    request.fail(function (response) {
      //Code à jouer en cas d'éxécution en erreur
      console.log(response);
    });
    request.always(function () {
      //Code à jouer après done OU fail dans tous les cas 
      window.location.reload();
    });

    event.preventDefault();
  });
  // positionnement sur le dernier item sélectionné dans une table
  $('table').each(function (index) {
    if ($crud_view && $crud_view.length > 0) {
      // nous sommes dans une vue simple 
      // et non pas dans une carte avec une vue
      // on passe
      return;
    }
    var $node = $(this);
    var $app = $node.data("app");
    var $table = $node.data("table");
    var $view = $node.data("view");
    var $anchorid = $app + "-" + $table + "-" + $view

    if (Cookies.get($anchorid)) {
      // est-ce que l'ancre existe ?
      var $anchor = $('tr[data-url="' + Cookies.get($anchorid) + '"]');
      if ($anchor.length) {
        // oui, c'est super
        // recherche du container scrollable 
        var $container = $anchor.closest('div');
        $container[0].scrollTo({
          top: $anchor.position().top - 100,
          left: 0,
          behavior: 'smooth'
        });
        $anchor.addClass("crud-list-selected");
      }
    }
  });
  // positionnement sur la dernière carte sélectionnée
  // sélection sur le clic sur crud-jquery-url
  if ($crud_view && $crud_view.length > 0) {
    var $anchorid = $crud_view
    if (Cookies.get($anchorid)) {
      var $anchor = $('div[data-url="' + Cookies.get($anchorid) + '"]');
      if ($anchor.length) {
        // cas des vues list_card
        $('html, body').animate({
          scrollTop: $anchor.offset().top - 200
        }, 1000);
        $anchor.addClass("crud-list-selected");
      } else {
        $anchor = $('tr[data-url="' + Cookies.get($anchorid) + '"]');
        if ($anchor.length) {
          // cas des vues list_table
          $('html, body').animate({
            scrollTop: $anchor.offset().top - 200
          }, 1000);
          $anchor.addClass("crud-list-selected");
        }
      }
    }
  };

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
            $('#beeForm').attr('action', $url);
            $('#beeForm', document).submit();
          }
        }).modal('show');
    } else {
      // Sans demande de confirmation
      $('#beeForm').attr('action', $url);
      $('#beeForm', document).submit()
    }
    event.preventDefault();
  });

  // query sql en ajax
  $('.crud-jquery-ajax').on('click', function (event) {
    var $datas = new FormData();
    var $url = $(this).data('url');
    var xsrf = $("#xsrf").val();
    $datas.append("_xsrf", xsrf);
    $.ajax({
      type: "POST",
      url: $url,
      data: $datas,
      dataType: 'json',
      cache: false,
      contentType: false,
      processData: false,
    })
      //Ce code sera exécuté en cas de succès - La réponse du serveur est passée à done()
      //On peut par exemple convertir cette réponse en chaine JSON et insérer cette chaine dans un div id="res"
      .done(function (data) {
        let mes = JSON.stringify(data);
        if (data.Response != "ok") {
          $.toast({
            message: data.Message,
            class: 'error',
            className: {
              toast: 'ui message'
            },
            position: 'bottom center',
            minDisplayTime: 1500
          });
        } else {
          $.toast({
            message: data.Message,
            class: 'success',
            className: {
              toast: 'ui message'
            },
            position: 'bottom center',
            minDisplayTime: 1500
          });
        }
        //$("div#res").append(mes);
      })
      //Ce code sera exécuté en cas d'échec - L'erreur est passée à fail()
      //On peut afficher les informations relatives à la requête et à l'erreur
      .fail(function (error) {
        $.toast({
          message: "La requête s'est terminée en échec. Infos : " + JSON.stringify(error),
          class: 'error',
          className: {
            toast: 'ui message'
          },
          position: 'bottom center',
          minDisplayTime: 1500
        });
      })
      //Ce code sera exécuté que la requête soit un succès ou un échec
      .always(function () {
        setTimeout(() => { window.location.reload(true) }, 1500);
      });
    event.preventDefault();
  });

  // Exécute en ajax une requête SQL sur le serveur et remplit les champs reçus du formulaire courant
  /**
    <a class="crud-jquery-ajax"
     data-url="/bee/ajax/{{$appid}}/{{$tableid}}/{{$viewid}}/{{$id}}/{{$key}}" title="{{$element.LabelLong}}">
    </a>
   */
  $('.crud-ajax-sql').on('click', function (event) {
    var $datas = new FormData();
    var $url = $(this).data('url');
    // ajout de données variables à la requête POST
    var $dataset = $(this).data();
    for( var d in $dataset) {
      if (d == "url") continue;
      var val = $("#" + $dataset[d]).val()
      $datas.append(d, val);
    }
    var xsrf = $("#xsrf").val();
    $datas.append("_xsrf", xsrf);
    $.ajax({
      type: "POST",
      url: $url,
      data: $datas,
      dataType: 'json',
      cache: false,
      contentType: false,
      processData: false,
    })
      //Ce code sera exécuté en cas de succès - La réponse du serveur est passée à done()
      //On peut par exemple convertir cette réponse en chaine JSON et insérer cette chaine dans un div id="res"
      .done(function (data) {
        if (data.Response != "ok") {
          $.toast({
            message: data.Message,
            class: 'error',
            className: {
              toast: 'ui message'
            },
            position: 'bottom center',
            minDisplayTime: 1500
          });
        } else {
          // mise à jour des rubriques trouvées dans la table
          for( var rub in data.Dataset) {
            if ($('#' + rub).is("select")) {
              $('#' + rub).dropdown('set selected', data.Dataset[rub]);
            } else {
              $('#' + rub).val(data.Dataset[rub])
            }

          }
          $.toast({
            message: data.Message,
            class: 'success',
            className: {
              toast: 'ui message'
            },
            position: 'bottom center',
            minDisplayTime: 1500
          });
        }
        //$("div#res").append(mes);
      })
      //Ce code sera exécuté en cas d'échec - L'erreur est passée à fail()
      //On peut afficher les informations relatives à la requête et à l'erreur
      .fail(function (error) {
        let mes = JSON.stringify(error);
        alert(mes)
        $.toast({
          message: "La requête s'est terminée en échec. Infos : " + JSON.stringify(error),
          class: 'error',
          className: {
            toast: 'ui message'
          },
          position: 'bottom center',
          minDisplayTime: 1500
        });
      })
      //Ce code sera exécuté que la requête soit un succès ou un échec
      // .always(function () {
      //   setTimeout(() => { window.location.reload(true) }, 1500);
      // });
    event.preventDefault();
  });

  // CLIC URL
  $('.crud-jquery-url').on('click', function (event) {
    if (isUsed) {
      event.preventDefault();
      return
    }
    // Mémo de la ligne sélectionnée d'une table dans un cookie
    if ($(this).prop("nodeName") == "TR") {
      var $node = $(this).closest("table");
      var $app = $node.data("app");
      var $table = $node.data("table");
      var $view = $node.data("view");
      var $anchorid = $app + "-" + $table + "-" + $view
      Cookies.set($anchorid, $(this).data("url"))
      $(this).addClass("crud-list-selected");
    } else if ($(this).hasClass("card")) {
      var $anchorid = $("#crud_view").val();
      Cookies.set($anchorid, $(this).data("url"))
      $(this).addClass("crud-list-selected");
    }
    var target = $(event.target);
    if (target.hasClass("crud-jquery-action") || target.parent().hasClass("crud-jquery-action")) {
      // pour laisser la main à crud-jquery-action
      // Cas d'un button dans une card
      event.preventDefault();
      return
    }
    if (target.hasClass("crud-jquery-ajax") || target.parent().hasClass("crud-jquery-ajax")) {
      // pour laisser la main à crud-jquery-ajax
      // Cas d'un button dans une card
      event.preventDefault();
      return
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
    $('#beeForm', document).submit();
    $(this).attr('disabled', 'disabled');
    event.preventDefault();
  });
  $('.field').on('keypress', function (event) {
    if (event.which == 13) {
      $('.crud-jquery-submit').trigger('click');
    }
  });

  // CLIC IMAGE POPUP
  $('.crud-popup-image').on('click', function (event) {
    isUsed = true;
    // Mémo du contexte dans un cookie
    if ($crud_view && $crud_view.length > 0) {
      var $anchor = $(this).closest('.card');
      Cookies.set($crud_view, $anchor.data("url"))
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
  // PDF VIEWER
  $('.crud-popup-pdf').on('click', function (event) {
    var $path = $(this).data('url') + "etc";
    var $base = $(this).data('url');
    var $height = screen.availHeight - 400;
    var $content = $('#crud-modal-pdf').find('.content');
    $content[0].innerHTML = '<object data="' + $path + '" type="application/pdf" height="' + $height + '" width="100%" typemustmatch></object>';
    $('#crud-modal-pdf')
      .modal({
        closable: true,
        onHide: function () {
          return true;
        },
        onApprove: function () {
          var link = document.createElement('a');
          link.href = $path;
          link.download = $base;
          link.click();
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
          $('#crud-form-delete-id', document).submit();
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